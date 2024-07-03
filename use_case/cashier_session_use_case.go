package use_case

import (
	"context"
	"io"
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

type cashierSessionLoaderParams struct {
	user bool
}

type CashierSessionUseCase interface {
	//  create
	Start(ctx context.Context, request dto_request.CashierSessionStartRequest) model.CashierSession

	//  read
	Fetch(ctx context.Context, request dto_request.CashierSessionFetchRequest) ([]model.CashierSession, int)
	FetchForCurrentUser(ctx context.Context, request dto_request.CashierSessionFetchForCurrentUserRequest) ([]model.CashierSession, int)
	FetchTransaction(ctx context.Context, request dto_request.CashierSessionFetchTransactionRequest) ([]model.Transaction, int)
	Get(ctx context.Context, request dto_request.CashierSessionGetRequest) model.CashierSession
	DownloadReport(ctx context.Context, requets dto_request.CashierSessionDownloadReportRequest) (io.ReadCloser, int64, string, string)
	GetByCurrentUser(ctx context.Context) *model.CashierSession

	//  update
	End(ctx context.Context) model.CashierSession
}

type cashierSessionUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewCashierSessionUseCase(
	repositoryManager repository.RepositoryManager,
) CashierSessionUseCase {
	return &cashierSessionUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *cashierSessionUseCase) mustValidateCashierSessionUserNotDuplicate(ctx context.Context, userId string) {
	isExist, err := u.repositoryManager.CashierSessionRepository().IsExistByUserIdAndStatusActive(ctx, userId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("CASHIER_SESSION.USER_ALREADY_HAVE_CASHIER_SESSION"))
	}
}

func (u *cashierSessionUseCase) mustValidateCashierSessionNoCart(ctx context.Context, cashierSessionId string) {
	isExist, err := u.repositoryManager.CartRepository().IsExistByCashierSessionId(ctx, cashierSessionId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("CASHIER_SESSION.STILL_HAVE_CART"))
	}
}

func (u *cashierSessionUseCase) mustLoadCashierSessionsData(ctx context.Context, cashierSessions []*model.CashierSession, option cashierSessionLoaderParams) {
	userLoader := loader.NewUserLoader(u.repositoryManager.UserRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range cashierSessions {
			if option.user {
				group.Go(userLoader.CashierSessionFn(cashierSessions[i]))
			}
		}
	}))
}

func (u *cashierSessionUseCase) Start(ctx context.Context, request dto_request.CashierSessionStartRequest) model.CashierSession {
	currentDateTime := util.CurrentDateTime()
	authUser := model.MustGetUserCtx(ctx)
	u.mustValidateCashierSessionUserNotDuplicate(ctx, authUser.Id)

	cashierSession := model.CashierSession{
		Id:           util.NewUuid(),
		UserId:       authUser.Id,
		Status:       data_type.CashierSessionStatusActive,
		StartingCash: request.StartingCash,
		EndingCash:   nil,
		StartedAt:    currentDateTime,
	}

	panicIfErr(
		u.repositoryManager.CashierSessionRepository().Insert(ctx, &cashierSession),
	)

	u.mustLoadCashierSessionsData(ctx, []*model.CashierSession{&cashierSession}, cashierSessionLoaderParams{
		user: true,
	})

	return cashierSession
}

func (u *cashierSessionUseCase) Fetch(ctx context.Context, request dto_request.CashierSessionFetchRequest) ([]model.CashierSession, int) {
	queryOption := model.CashierSessionQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		StartedAtGte: request.StartedAt,
		EndedAtLte:   request.EndedAt,
		UserId:       request.UserId,
		Status:       request.Status,
		Phrase:       request.Phrase,
	}

	cashierSessions, err := u.repositoryManager.CashierSessionRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CashierSessionRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadCashierSessionsData(ctx, util.SliceValueToSlicePointer(cashierSessions), cashierSessionLoaderParams{
		user: true,
	})

	return cashierSessions, total
}

func (u *cashierSessionUseCase) FetchForCurrentUser(ctx context.Context, request dto_request.CashierSessionFetchForCurrentUserRequest) ([]model.CashierSession, int) {
	currentUser := model.MustGetUserCtx(ctx)

	queryOption := model.CashierSessionQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		UserId:       &currentUser.Id,
		StartedAtGte: request.StartedAt,
		EndedAtLte:   request.EndedAt,
		Status:       request.Status,
		Phrase:       request.Phrase,
	}

	cashierSessions, err := u.repositoryManager.CashierSessionRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CashierSessionRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadCashierSessionsData(ctx, util.SliceValueToSlicePointer(cashierSessions), cashierSessionLoaderParams{
		user: true,
	})

	return cashierSessions, total
}

func (u *cashierSessionUseCase) FetchTransaction(ctx context.Context, request dto_request.CashierSessionFetchTransactionRequest) ([]model.Transaction, int) {
	mustGetCashierSession(ctx, u.repositoryManager, request.CashierSessionId, true)
	queryOption := model.TransactionQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		CashierSessionId: &request.CashierSessionId,
		Status:           request.Status,
	}

	transactions, err := u.repositoryManager.TransactionRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.TransactionRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return transactions, total
}

func (u *cashierSessionUseCase) Get(ctx context.Context, request dto_request.CashierSessionGetRequest) model.CashierSession {
	cashierSession := mustGetCashierSession(ctx, u.repositoryManager, request.CashierSessionId, false)

	u.mustLoadCashierSessionsData(ctx, []*model.CashierSession{&cashierSession}, cashierSessionLoaderParams{
		user: true,
	})

	return cashierSession
}

func (u *cashierSessionUseCase) DownloadReport(ctx context.Context, request dto_request.CashierSessionDownloadReportRequest) (io.ReadCloser, int64, string, string) {
	cashierSession := mustGetCashierSession(ctx, u.repositoryManager, request.CashierSessionId, false)

	if cashierSession.Status != data_type.CashierSessionStatusCompleted {
		panic(dto_response.NewBadRequestErrorResponse("CASHIER_SESSION.STATUS_MUST_BE_COMPLETED"))
	}

	userLoader := loader.NewUserLoader(u.repositoryManager.UserRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		group.Go(userLoader.CashierSessionFn(&cashierSession))
	}))

	transactions, err := u.repositoryManager.TransactionRepository().Fetch(ctx, model.TransactionQueryOption{
		CashierSessionId: &request.CashierSessionId,
	})
	panicIfErr(err)

	transactionItemsLoader := loader.NewTransactionItemsLoader(u.repositoryManager.TransactionItemRepository())
	transactionPaymentsLoader := loader.NewTransactionPaymentsLoader(u.repositoryManager.TransactionPaymentRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			group.Go(transactionItemsLoader.TransactionFn(&transactions[i]))
			group.Go(transactionPaymentsLoader.TransactionFn(&transactions[i]))
		}
	}))

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	transactionItemCostsLoader := loader.NewTransactionItemCostsLoader(u.repositoryManager.TransactionItemCostRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			for j := range transactions[i].TransactionItems {
				group.Go(productUnitLoader.TransactionItemFn(&transactions[i].TransactionItems[j]))
				group.Go(transactionItemCostsLoader.TransactionItemFn(&transactions[i].TransactionItems[j]))
			}
		}
	}))

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range transactions {
			for j := range transactions[i].TransactionItems {
				group.Go(productLoader.ProductUnitFn(transactions[i].TransactionItems[j].ProductUnit))
				group.Go(unitLoader.ProductUnitFn(transactions[i].TransactionItems[j].ProductUnit))
			}
		}
	}))

	reportExcel, err := NewReportTransactionExcel(
		util.CurrentDateTime(),
		cashierSession,
	)
	panicIfErr(err)

	for _, transaction := range transactions {
		for _, transactionItem := range transaction.TransactionItems {
			for _, cost := range transactionItem.TransactionItemCosts {
				// calculate revenue
				revenue := transactionItem.PricePerUnit
				if transactionItem.DiscountPerUnit != nil {
					revenue -= *transactionItem.DiscountPerUnit
				}
				revenue -= cost.BaseCostPrice
				revenue *= transactionItem.Qty

				// discount
				discountPerUnit := 0.0
				if transactionItem.DiscountPerUnit != nil {
					discountPerUnit = *transactionItem.DiscountPerUnit
				}

				// payment method
				paymentMethod := ""
				if len(transaction.TransactionPayments) > 0 {
					paymentMethod = transaction.TransactionPayments[0].PaymentType.String()
				}

				// calculate total (follow cost, not transaction item)
				total := cost.Qty * (transactionItem.PricePerUnit - discountPerUnit)

				// add data
				reportExcel.AddSheet1Data(ReportTransactionExcelSheet1Data{
					Id:              transaction.Id,
					Status:          transaction.Status.String(),
					PaymentMethod:   paymentMethod,
					ProductId:       transactionItem.ProductUnit.ProductId,
					UnitId:          transactionItem.ProductUnit.UnitId,
					ProductName:     transactionItem.ProductUnit.Product.Name,
					UnitName:        transactionItem.ProductUnit.Unit.Name,
					Qty:             cost.Qty,
					PricePerUnit:    transactionItem.PricePerUnit,
					DiscountPerUnit: discountPerUnit,
					Total:           total,
					CostPerUnit:     cost.BaseCostPrice,
					Revenue:         revenue,
					PaymentAt:       transaction.PaymentAt.DateTime().Time(),
				})
			}
		}
	}

	readCloser, contentLength, err := reportExcel.ToReadSeekCloserWithContentLength()
	panicIfErr(err)

	return readCloser, contentLength, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "cashier_session_transactions.xlsx"
}

func (u *cashierSessionUseCase) GetByCurrentUser(ctx context.Context) *model.CashierSession {
	authUser := model.MustGetUserCtx(ctx)
	cashierSession, err := u.repositoryManager.CashierSessionRepository().GetByUserIdAndStatusActive(ctx, authUser.Id)
	panicIfErr(err, constant.ErrNoData)

	u.mustLoadCashierSessionsData(ctx, []*model.CashierSession{cashierSession}, cashierSessionLoaderParams{
		user: true,
	})

	return cashierSession
}

func (u *cashierSessionUseCase) End(ctx context.Context) model.CashierSession {
	currentDateTime := util.CurrentDateTime()
	authUser := model.MustGetUserCtx(ctx)
	cashierSession, err := u.repositoryManager.CashierSessionRepository().GetByUserIdAndStatusActive(ctx, authUser.Id)
	panicIfErr(err, constant.ErrNoData)

	if cashierSession == nil {
		panic("CASHIER_SESSION.NO_ACTIVE_SESSION")
	}

	if cashierSession.Status != data_type.CashierSessionStatusActive {
		panic(dto_response.NewBadRequestErrorResponse("CASHIER_SESSION.STATUS_MUST_BE_ACTIVE"))
	}

	u.mustValidateCashierSessionNoCart(ctx, cashierSession.Id)

	totalCashPayment, err := u.repositoryManager.TransactionPaymentRepository().GetTotalPaymentByCashierSessionIdAndPaymentType(ctx, cashierSession.Id, data_type.TransactionPaymentTypeCash)
	panicIfErr(err)

	cashierSession.EndingCash = util.Pointer(cashierSession.StartingCash + *totalCashPayment)
	cashierSession.Status = data_type.CashierSessionStatusCompleted
	cashierSession.EndedAt = currentDateTime.NullDateTime()

	panicIfErr(
		u.repositoryManager.CashierSessionRepository().Update(ctx, cashierSession),
	)

	u.mustLoadCashierSessionsData(ctx, []*model.CashierSession{cashierSession}, cashierSessionLoaderParams{
		user: true,
	})

	return *cashierSession
}
