package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type BalanceUseCase interface {
	// admin create
	AdminCreate(ctx context.Context, request dto_request.AdminBalanceCreateRequest) model.Balance

	// admin read
	AdminFetch(ctx context.Context, request dto_request.AdminBalanceFetchRequest) ([]model.Balance, int)
	AdminGet(ctx context.Context, request dto_request.AdminBalanceGetRequest) model.Balance

	// admin update
	AdminUpdate(ctx context.Context, request dto_request.AdminBalanceUpdateRequest) model.Balance

	// admin delete
	AdminDelete(ctx context.Context, request dto_request.AdminBalanceDeleteRequest)
}

type balanceUseCase struct {
	repositoryManager repository.RepositoryManager
}

func (u *balanceUseCase) mustValidateAllowDeleteBalance(ctx context.Context, balanceId string) {

}

func (u *balanceUseCase) AdminCreate(ctx context.Context, request dto_request.AdminBalanceCreateRequest) model.Balance {
	balance := model.Balance{
		Id:            util.NewUuid(),
		AccountNumber: request.AccountNumber,
		AccountName:   request.AccountName,
		BankName:      request.BankName,
		Name:          request.Name,
		Amount:        0,
	}

	panicIfErr(
		u.repositoryManager.BalanceRepository().Insert(ctx, &balance),
	)

	return balance
}

func (u *balanceUseCase) AdminFetch(ctx context.Context, request dto_request.AdminBalanceFetchRequest) ([]model.Balance, int) {
	queryOption := model.BalanceQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	balances, err := u.repositoryManager.BalanceRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.BalanceRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return balances, total
}

func (u *balanceUseCase) AdminGet(ctx context.Context, request dto_request.AdminBalanceGetRequest) model.Balance {
	balance := mustGetBalance(ctx, u.repositoryManager, request.BalanceId, true)

	return balance
}

func (u *balanceUseCase) AdminUpdate(ctx context.Context, request dto_request.AdminBalanceUpdateRequest) model.Balance {
	balance := mustGetBalance(ctx, u.repositoryManager, request.BalanceId, true)

	balance.Name = request.Name
	balance.AccountName = request.AccountName
	balance.AccountNumber = request.AccountNumber
	balance.BankName = request.BankName

	panicIfErr(
		u.repositoryManager.BalanceRepository().Update(ctx, &balance),
	)

	return balance
}

func (u *balanceUseCase) AdminDelete(ctx context.Context, request dto_request.AdminBalanceDeleteRequest) {
	balance := mustGetBalance(ctx, u.repositoryManager, request.BalanceId, true)

	u.mustValidateAllowDeleteBalance(ctx, request.BalanceId)

	panicIfErr(
		u.repositoryManager.BalanceRepository().Delete(ctx, &balance),
	)
}

func NewBalanceUseCase(
	repositoryManager repository.RepositoryManager,
) BalanceUseCase {
	return &balanceUseCase{
		repositoryManager: repositoryManager,
	}
}