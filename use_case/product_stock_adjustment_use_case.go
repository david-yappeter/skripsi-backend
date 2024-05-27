package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
)

type ProductStockAdjustmentUseCase interface {
	// read
	Fetch(ctx context.Context, request dto_request.ProductStockAdjustmentFetchRequest) ([]model.ProductStockAdjustment, int)
}

type productStockAdjustmentUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewProductStockAdjustmentUseCase(
	repositoryManager repository.RepositoryManager,
) ProductStockAdjustmentUseCase {
	return &productStockAdjustmentUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *productStockAdjustmentUseCase) Fetch(ctx context.Context, request dto_request.ProductStockAdjustmentFetchRequest) ([]model.ProductStockAdjustment, int) {
	queryOption := model.ProductStockAdjustmentQueryOption{
		QueryOption:    model.NewQueryOptionWithPagination(request.Page, request.Limit, model.Sorts(request.Sorts)),
		UserId:         request.UserId,
		ProductStockId: request.ProductStockId,
	}

	productStockAdjustments, err := u.repositoryManager.ProductStockAdjustmentRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductStockAdjustmentRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return productStockAdjustments, total
}
