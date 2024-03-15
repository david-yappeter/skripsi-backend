package use_case

import (
	"context"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type TransactionUseCase interface {
	//  create
	CheckoutCart(ctx context.Context, request dto_request.TransactionCheckoutCartRequest) model.Transaction
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

func (u *transactionUseCase) CheckoutCart(ctx context.Context, request dto_request.TransactionCheckoutCartRequest) model.Transaction {
	currentTime := util.CurrentDateTime()
	cashierSession := u.mustGetCurrentUserActiveCashierSession(ctx)

	cart := u.shouldGetActiveCartByCashierSessionId(ctx, cashierSession.Id)
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

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range cartItems {
				group.Go(productLoader.ProductUnitFn(cartItems[i].ProductUnit))
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
						discountPerUnit = util.Float64P(*productDiscount.DiscountPercentage * *cartItem.ProductUnit.Product.Price)
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
				}
				transactionItems = append(transactionItems, transactionItem)

				// deduct product stock
				currentProductStock := productStockMapByProductUnitId[cartItem.ProductUnitId]
				currentProductStock.Qty -= cartItem.Qty

				productStockMapByProductUnitId[cartItem.ProductUnitId] = currentProductStock

				// deduct product stock mutation
				deductQtyLeft := cartItem.Qty
				for deductQtyLeft > 0 {
					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductUnitIdAndBaseQtyLeftNotZero(ctx, cartItem.ProductUnitId)
					if err != nil {
						return err
					}

					if deductQtyLeft > productStockMutation.BaseQtyLeft {
						transactionItemCosts = append(transactionItemCosts, model.TransactionItemCost{
							Id:                util.NewUuid(),
							TransactionItemId: transactionItem.Id,
							Qty:               productStockMutation.BaseQtyLeft,
							BaseCostPrice:     productStockMutation.BaseCostPrice,
							TotalCostPrice:    productStockMutation.BaseCostPrice * productStockMutation.BaseQtyLeft * productStockMutation.ScaleToBase,
						})

						deductQtyLeft -= productStockMutation.BaseQtyLeft
						productStockMutation.BaseQtyLeft = 0
					} else {
						transactionItemCosts = append(transactionItemCosts, model.TransactionItemCost{
							Id:                util.NewUuid(),
							TransactionItemId: transactionItem.Id,
							Qty:               deductQtyLeft,
							BaseCostPrice:     productStockMutation.BaseCostPrice,
							TotalCostPrice:    productStockMutation.BaseCostPrice * deductQtyLeft * productStockMutation.ScaleToBase,
						})

						productStockMutation.BaseQtyLeft -= deductQtyLeft
					}

					if err := productStockMutationRepository.Update(ctx, productStockMutation); err != nil {
						return err
					}
				}
			}

			// assign transaction total
			transaction.Total = transactionTotal

			if err := transactionRepository.Insert(ctx, &transaction); err != nil {
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
				if err := productStockRepository.Update(ctx, &productStock); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	return transaction
}
