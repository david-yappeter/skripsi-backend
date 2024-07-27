package use_case

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/internal/printer_template"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
	"golang.org/x/text/language"
)

type transactionLoaderParams struct {
	items   bool
	payment bool
}

type TransactionUseCase interface {
	//  create
	CheckoutCart(ctx context.Context, request dto_request.TransactionCheckoutCartRequest) (model.Transaction, []int16)

	// read
	Get(ctx context.Context, request dto_request.TransactionGetRequest) model.Transaction
	Reprint(ctx context.Context, request dto_request.TransactionReprintRequest) []int16
}

type transactionUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewTransactionUseCase(
	repositoryManager repository.RepositoryManager,
) TransactionUseCase {
	return &transactionUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *transactionUseCase) mustGetCurrentUserActiveCashierSession(ctx context.Context) model.CashierSession {
	authUser := model.MustGetUserCtx(ctx)

	cashierSession, err := u.repositoryManager.CashierSessionRepository().GetByUserIdAndStatusActive(ctx, authUser.Id)
	panicIfErr(err, constant.ErrNoData)

	if cashierSession == nil {
		panic(dto_response.NewBadRequestErrorResponse("TRANSACTION.USER_MUST_HAVE_ACTIVE_CASHIER_SESSION"))
	}

	return *cashierSession
}

func (u *transactionUseCase) shouldGetActiveCartByCashierSessionId(ctx context.Context, cashierSessionId string) *model.Cart {
	cart, err := u.repositoryManager.CartRepository().GetByCashierSessionIdAndIsActive(ctx, cashierSessionId, true)
	panicIfErr(err, constant.ErrNoData)

	return cart
}

func (u *transactionUseCase) mustLoadTransactionDatas(ctx context.Context, transactions []*model.Transaction, option transactionLoaderParams) {
	transactionItemsLoader := loader.NewTransactionItemsLoader(u.repositoryManager.TransactionItemRepository())
	transactionPaymentsLoader := loader.NewTransactionPaymentsLoader(u.repositoryManager.TransactionPaymentRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			if option.items {
				group.Go(transactionItemsLoader.TransactionFn(transactions[i]))
			}

			if option.payment {
				group.Go(transactionPaymentsLoader.TransactionFn(transactions[i]))
			}
		}
	}))

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			if option.items {
				for j := range transactions[i].TransactionItems {
					group.Go(productUnitLoader.TransactionItemFn(&transactions[i].TransactionItems[j]))
				}
			}
		}
	}))

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			if option.items {
				for j := range transactions[i].TransactionItems {
					if transactions[i].TransactionItems[j].ProductUnit != nil {
						group.Go(productLoader.ProductUnitFn(transactions[i].TransactionItems[j].ProductUnit))
						group.Go(unitLoader.ProductUnitFn(transactions[i].TransactionItems[j].ProductUnit))
					}
				}
			}
		}
	}))

}

func (u *transactionUseCase) CheckoutCart(ctx context.Context, request dto_request.TransactionCheckoutCartRequest) (model.Transaction, []int16) {
	switch request.PaymentType {
	case data_type.TransactionPaymentTypeCash:
		if request.CashPaid == nil {
			panic(dto_response.NewBadRequestErrorResponse("TRANSACTION.PAYMENT_TYPE_CASH_REQUIRED_CASH_PAID"))
		}
	}

	currentTime := util.CurrentDateTime()
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)
	cashierUser := mustGetUser(ctx, u.repositoryManager, cashierSession.UserId, true)

	cart := u.shouldGetActiveCartByCashierSessionId(ctx, cashierSession.Id)
	if cart == nil {
		panic(dto_response.NewBadRequestErrorResponse("TRANSACTION.NO_ACTIVE_CART"))
	}

	cartItems, err := u.repositoryManager.CartItemRepository().FetchByCartIds(ctx, []string{cart.Id})
	panicIfErr(err)

	if len(cartItems) == 0 {
		panic(dto_response.NewBadRequestErrorResponse("TRANSACTION.CART_ITEMS_EMPTY"))
	}

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range cartItems {
				group.Go(productUnitLoader.CartItemFn(&cartItems[i]))
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range cartItems {
				group.Go(productLoader.ProductUnitFn(cartItems[i].ProductUnit))
				group.Go(unitLoader.ProductUnitFn(cartItems[i].ProductUnit))
			}
		}),
	)

	productDiscountLoader := loader.NewProductDiscountLoader(u.repositoryManager.ProductDiscountRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range cartItems {
				group.Go(productDiscountLoader.ProductFnNotStrict(cartItems[i].ProductUnit.Product))
			}
		}),
	)

	var (
		transaction                    = model.Transaction{}
		transactionPayment             = model.TransactionPayment{}
		transactionTotal               = 0.0
		transactionItems               = []model.TransactionItem{}
		transactionItemCosts           = []model.TransactionItemCost{}
		productUnitsMapById            = map[string]model.ProductUnit{}
		productStockMapByProductUnitId = map[string]model.ProductStock{}
	)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			cartRepository := u.repositoryManager.CartRepository()
			cartItemRepository := u.repositoryManager.CartItemRepository()
			transactionRepository := u.repositoryManager.TransactionRepository()
			transactionPaymentRepository := u.repositoryManager.TransactionPaymentRepository()
			transactionItemRepository := u.repositoryManager.TransactionItemRepository()
			transactionItemCostRepository := u.repositoryManager.TransactionItemCostRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			// create transaction
			transaction = model.Transaction{
				Id:               util.NewUuid(),
				CashierSessionId: cashierSession.Id,
				Status:           data_type.TransactionStatusPaid,
				Total:            0,
				PaymentAt:        currentTime.NullDateTime(),
			}

			// fetch product stock
			for _, cartItem := range cartItems {
				productUnit := mustGetProductUnit(ctx, u.repositoryManager, cartItem.ProductUnitId, true)
				productUnitsMapById[productUnit.Id] = productUnit

				productStock, err := u.repositoryManager.ProductStockRepository().GetByProductId(ctx, productUnit.ProductId)
				if err != nil {
					return err
				}

				productStockMapByProductUnitId[productUnit.Id] = *productStock
			}

			for _, cartItem := range cartItems {
				// add transaction total
				transactionTotal += cartItem.Qty * cartItem.ProductUnit.ScaleToBase * *cartItem.ProductUnit.Product.Price

				// determine if transaction applied discount (must check minimum quantity)
				var discountPerUnit *float64 = nil
				productDiscount := cartItem.ProductUnit.Product.ProductDiscount

				if productDiscount != nil && cartItem.Qty >= productDiscount.MinimumQty {
					if productDiscount.DiscountAmount != nil {
						discountPerUnit = productDiscount.DiscountAmount
					} else {
						discountPerUnit = util.Float64P(*productDiscount.DiscountPercentage * *cartItem.ProductUnit.Product.Price / 100.0)
					}
				}

				// deduct transaction total if discount exist
				if discountPerUnit != nil {
					transactionTotal -= *discountPerUnit
				}

				// create transaction items
				transactionItem := model.TransactionItem{
					Id:              util.NewUuid(),
					TransactionId:   transaction.Id,
					ProductUnitId:   cartItem.ProductUnitId,
					Qty:             cartItem.Qty,
					PricePerUnit:    *cartItem.ProductUnit.Product.Price,
					DiscountPerUnit: discountPerUnit,
					ProductUnit:     cartItem.ProductUnit,
				}
				transactionItems = append(transactionItems, transactionItem)

				// deduct product stock
				currentProductStock := productStockMapByProductUnitId[cartItem.ProductUnitId]
				currentProductStock.Qty -= cartItem.Qty

				productStockMapByProductUnitId[cartItem.ProductUnitId] = currentProductStock

				// deduct product stock mutation
				deductQtyLeft := cartItem.Qty
				for deductQtyLeft > 0 {
					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx, cartItem.ProductUnit.ProductId)
					if err != nil {
						return err
					}

					if deductQtyLeft > productStockMutation.BaseQtyLeft {
						transactionItemCosts = append(transactionItemCosts, model.TransactionItemCost{
							Id:                util.NewUuid(),
							TransactionItemId: transactionItem.Id,
							Qty:               productStockMutation.BaseQtyLeft,
							BaseCostPrice:     currentProductStock.BaseCostPrice,
							TotalCostPrice:    currentProductStock.BaseCostPrice * productStockMutation.BaseQtyLeft * productStockMutation.ScaleToBase,
						})

						deductQtyLeft -= productStockMutation.BaseQtyLeft
						productStockMutation.BaseQtyLeft = 0
					} else {
						transactionItemCosts = append(transactionItemCosts, model.TransactionItemCost{
							Id:                util.NewUuid(),
							TransactionItemId: transactionItem.Id,
							Qty:               deductQtyLeft,
							BaseCostPrice:     currentProductStock.BaseCostPrice,
							TotalCostPrice:    currentProductStock.BaseCostPrice * deductQtyLeft * productStockMutation.ScaleToBase,
						})

						productStockMutation.BaseQtyLeft -= deductQtyLeft
						deductQtyLeft = 0
					}

					if err := productStockMutationRepository.Update(ctx, productStockMutation); err != nil {
						return err
					}
				}
			}

			// assign transaction total
			transaction.Total = transactionTotal

			// transaction payment
			transactionPayment = model.TransactionPayment{
				Id:              util.NewUuid(),
				TransactionId:   transaction.Id,
				PaymentType:     request.PaymentType,
				ReferenceNumber: request.ReferenceNumber,
				Total:           transactionTotal,
				TotalPaid:       transactionTotal,
			}
			switch request.PaymentType {
			case data_type.TransactionPaymentTypeCash:
				if *request.CashPaid < transactionTotal {
					panic(dto_response.NewBadRequestErrorResponse("TRANSACTION.INVALID_CASH_PAID_AMOUNT"))
				}
				transactionPayment.TotalPaid = *request.CashPaid
			case data_type.TransactionPaymentTypeBcaTransfer:
				// skip
			default:
				panic("unhandled")
			}

			if err := transactionRepository.Insert(ctx, &transaction); err != nil {
				return err
			}

			if err := transactionPaymentRepository.Insert(ctx, &transactionPayment); err != nil {
				return err
			}

			if err := transactionItemRepository.InsertMany(ctx, transactionItems); err != nil {
				return err
			}

			if err := transactionItemCostRepository.InsertMany(ctx, transactionItemCosts); err != nil {
				return err
			}

			if err := cartItemRepository.DeleteManyByCartId(ctx, cart.Id); err != nil {
				return err
			}

			if err := cartRepository.Delete(ctx, cart); err != nil {
				return err
			}

			for _, productStock := range productStockMapByProductUnitId {
				// update tiktok product stock
				tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, productStock.ProductId)

				if tiktokProduct != nil {
					mustUpdateTiktokProductInventory(ctx, u.repositoryManager, tiktokProduct.TiktokProductId, int(productStock.Qty))
				}

				if err := productStockRepository.Update(ctx, &productStock); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	transaction.TransactionItems = transactionItems
	transaction.TransactionPayments = append(transaction.TransactionPayments, transactionPayment)

	isCash := request.PaymentType == data_type.TransactionPaymentTypeCash

	templateItems := []printer_template.EscposReceiptItem{}
	for _, transactionItem := range transaction.TransactionItems {
		templateItems = append(templateItems, printer_template.EscposReceiptItem{
			Name:        transactionItem.ProductUnit.Product.Name,
			PricePerQty: util.CurrencyFormat(int(transactionItem.PricePerUnitAfterDiscount()), language.Indonesian),
			Qty:         fmt.Sprintf("%.2f", transactionItem.Qty),
			TotalPrice:  util.CurrencyFormat(int(transactionItem.GrossTotal()), language.Indonesian),
		})
	}

	templateAttribute := printer_template.EscposReceiptTemplateAttribute{
		StoreName:      "Toko Setia Abadi",
		Address:        "Jl. Marelan Raya No.88 A-B, Tanah Enam Ratus, Kec. Medan Marelan",
		PhoneNumber:    "081362337116",
		Date:           transaction.PaymentAt.DateTime().Format("02/01/2006 15:04:05"),
		Cashier:        cashierUser.Name,
		Items:          templateItems,
		SubTotal:       nil,
		DiscountAmount: nil,
		GrandTotal:     util.CurrencyFormat(int(transaction.Total), language.Indonesian),
		Paid:           util.CurrencyFormat(int(transaction.TotalPayment()), language.Indonesian),
		Change:         util.CurrencyFormat(int(transaction.TotalPayment()-transaction.Total), language.Indonesian),
		OpenDrawer:     true,
		IsCash:         isCash,
	}

	printerDataContent := printer_template.EscposReceiptTemplate(templateAttribute)

	return transaction, util.ArrayUint8ToArrayInt16(printerDataContent)
}

func (u *transactionUseCase) Get(ctx context.Context, request dto_request.TransactionGetRequest) model.Transaction {
	transaction := mustGetTransaction(ctx, u.repositoryManager, request.TransactionId, true)

	u.mustLoadTransactionDatas(ctx, []*model.Transaction{&transaction}, transactionLoaderParams{
		items:   true,
		payment: true,
	})

	return transaction
}

func (u *transactionUseCase) Reprint(ctx context.Context, request dto_request.TransactionReprintRequest) []int16 {
	transaction := mustGetTransaction(ctx, u.repositoryManager, request.TransactionId, true)
	cashierSession := mustGetCashierSession(ctx, u.repositoryManager, transaction.CashierSessionId, true)
	cashierUser := mustGetUser(ctx, u.repositoryManager, cashierSession.UserId, true)

	u.mustLoadTransactionDatas(ctx, []*model.Transaction{&transaction}, transactionLoaderParams{
		items:   true,
		payment: true,
	})

	isCash := transaction.TransactionPayments[0].PaymentType == data_type.TransactionPaymentTypeCash

	templateItems := []printer_template.EscposReceiptItem{}
	for _, transactionItem := range transaction.TransactionItems {
		templateItems = append(templateItems, printer_template.EscposReceiptItem{
			Name:        transactionItem.ProductUnit.Product.Name,
			PricePerQty: util.CurrencyFormat(int(transactionItem.PricePerUnitAfterDiscount()), language.Indonesian),
			Qty:         fmt.Sprintf("%.2f", transactionItem.Qty),
			TotalPrice:  util.CurrencyFormat(int(transactionItem.GrossTotal()), language.Indonesian),
		})
	}

	templateAttribute := printer_template.EscposReceiptTemplateAttribute{
		StoreName:      "Toko Setia Abadi",
		Address:        "Jl. Marelan Raya No.88 A-B, Tanah Enam Ratus, Kec. Medan Marelan",
		PhoneNumber:    "081362337116",
		Date:           transaction.PaymentAt.DateTime().Format("02/01/2006 15:04:05"),
		Cashier:        cashierUser.Name,
		Items:          templateItems,
		SubTotal:       nil,
		DiscountAmount: nil,
		GrandTotal:     util.CurrencyFormat(int(transaction.Total), language.Indonesian),
		Paid:           util.CurrencyFormat(int(transaction.TotalPayment()), language.Indonesian),
		Change:         util.CurrencyFormat(int(transaction.TotalPayment()-transaction.Total), language.Indonesian),
		OpenDrawer:     true,
		IsCash:         isCash,
	}

	printerDataContent := printer_template.EscposReceiptTemplate(templateAttribute)

	return util.ArrayUint8ToArrayInt16(printerDataContent)
}
