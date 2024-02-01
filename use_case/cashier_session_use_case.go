package use_case

import (
	"context"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type CashierSessionUseCase interface {
	//  create
	Start(ctx context.Context, request dto_request.CashierSessionStartRequest) model.CashierSession

	//  read
	Fetch(ctx context.Context, request dto_request.CashierSessionFetchRequest) ([]model.CashierSession, int)
	Get(ctx context.Context, request dto_request.CashierSessionGetRequest) model.CashierSession
	GetByCurrentUser(ctx context.Context) *model.CashierSession

	//  update
	MarkComplete(ctx context.Context, request dto_request.CashierSessionMarkCompleteRequest) model.CashierSession
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

func (u *cashierSessionUseCase) mustValidateAllowDeleteCashierSession(ctx context.Context, cashierSessionId string) {

}

func (u *cashierSessionUseCase) Start(ctx context.Context, request dto_request.CashierSessionStartRequest) model.CashierSession {
	authUser := model.MustGetUserCtx(ctx)
	u.mustValidateCashierSessionUserNotDuplicate(ctx, authUser.Id)

	cashierSession := model.CashierSession{
		Id:           util.NewUuid(),
		UserId:       authUser.Id,
		Status:       data_type.CashierSessionStatusActive,
		StartingCash: request.StartingCash,
		EndingCash:   nil,
	}

	panicIfErr(
		u.repositoryManager.CashierSessionRepository().Insert(ctx, &cashierSession),
	)

	return cashierSession
}

func (u *cashierSessionUseCase) Fetch(ctx context.Context, request dto_request.CashierSessionFetchRequest) ([]model.CashierSession, int) {
	queryOption := model.CashierSessionQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	cashierSessions, err := u.repositoryManager.CashierSessionRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CashierSessionRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return cashierSessions, total
}

func (u *cashierSessionUseCase) Get(ctx context.Context, request dto_request.CashierSessionGetRequest) model.CashierSession {
	cashierSession := mustGetCashierSession(ctx, u.repositoryManager, request.CashierSessionId, false)

	return cashierSession
}

func (u *cashierSessionUseCase) GetByCurrentUser(ctx context.Context) *model.CashierSession {
	authUser := model.MustGetUserCtx(ctx)
	cashierSession, err := u.repositoryManager.CashierSessionRepository().GetByUserIdAndStatusActive(ctx, authUser.Id)
	panicIfErr(err)

	return cashierSession
}

func (u *cashierSessionUseCase) MarkComplete(ctx context.Context, request dto_request.CashierSessionMarkCompleteRequest) model.CashierSession {
	cashierSession := mustGetCashierSession(ctx, u.repositoryManager, request.CashierSessionId, false)

	cashierSession.EndingCash = &request.EndingCash
	cashierSession.Status = data_type.CashierSessionStatusCompleted

	panicIfErr(
		u.repositoryManager.CashierSessionRepository().Update(ctx, &cashierSession),
	)

	return cashierSession
}
