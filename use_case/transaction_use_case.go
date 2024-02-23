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

	var (
		transactionTotal = 0.0
		transactionItems []model.TransactionItem
	)

	for _, cartItem := range cartItems {
		transactionTotal += cartItem.Qty * cartItem.ProductUnit.ScaleToBase * *cartItem.ProductUnit.Product.Price
	}

	transaction := model.Transaction{
		Id:               util.NewUuid(),
		CashierSessionId: cashierSession.Id,
		Status:           data_type.TransactionStatusProductPaid,
		Total:            transactionTotal,
		PaymentAt:        currentTime.NullDateTime(),
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			transactionRepository := u.repositoryManager.TransactionRepository()
			transactionItemRepository := u.repositoryManager.TransactionItemRepository()

			if err := transactionRepository.Insert(ctx, &transaction); err != nil {
				return err
			}

			if err := transactionItemRepository.InsertMany(ctx, transactionItems); err != nil {
				return err
			}

			return nil
		}),
	)

	return transaction
}
