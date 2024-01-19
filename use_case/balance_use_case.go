package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type BalanceUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.BalanceCreateRequest) model.Balance

	//  read
	Fetch(ctx context.Context, request dto_request.BalanceFetchRequest) ([]model.Balance, int)
	Get(ctx context.Context, request dto_request.BalanceGetRequest) model.Balance

	//  update
	Update(ctx context.Context, request dto_request.BalanceUpdateRequest) model.Balance

	//  delete
	Delete(ctx context.Context, request dto_request.BalanceDeleteRequest)
}

type balanceUseCase struct {
	repositoryManager repository.RepositoryManager
}

func (u *balanceUseCase) mustValidateAllowDeleteBalance(ctx context.Context, balanceId string) {

}

func (u *balanceUseCase) Create(ctx context.Context, request dto_request.BalanceCreateRequest) model.Balance {
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

func (u *balanceUseCase) Fetch(ctx context.Context, request dto_request.BalanceFetchRequest) ([]model.Balance, int) {
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

func (u *balanceUseCase) Get(ctx context.Context, request dto_request.BalanceGetRequest) model.Balance {
	balance := mustGetBalance(ctx, u.repositoryManager, request.BalanceId, true)

	return balance
}

func (u *balanceUseCase) Update(ctx context.Context, request dto_request.BalanceUpdateRequest) model.Balance {
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

func (u *balanceUseCase) Delete(ctx context.Context, request dto_request.BalanceDeleteRequest) {
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
