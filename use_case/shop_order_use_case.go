package use_case

import (
	"context"
	"myapp/model"
	"myapp/repository"
)

type ShopOrderItemUseCase interface {
	Fetch(ctx context.Context) ([]model.ShopOrder, int)
}

type shopOrderItemUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewShopOrderItemUseCase(
	repositoryManager repository.RepositoryManager,
) ShopOrderItemUseCase {
	return &shopOrderItemUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *shopOrderItemUseCase) Fetch(ctx context.Context) ([]model.ShopOrder, int) {

	return []model.ShopOrder{}, 0
}
