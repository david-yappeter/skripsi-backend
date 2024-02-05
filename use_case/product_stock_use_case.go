package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type productStockLoaderParams struct {
	product bool
}

type ProductStockUseCase interface {
	//  read
	Fetch(ctx context.Context, request dto_request.ProductStockFetchRequest) ([]model.ProductStock, int)
	Get(ctx context.Context, request dto_request.ProductStockGetRequest) model.ProductStock

	//  update
	Adjustment(ctx context.Context, request dto_request.ProductStockAdjustmentRequest) model.ProductStock
}

type productStockUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewProductStockUseCase(
	repositoryManager repository.RepositoryManager,
) ProductStockUseCase {
	return &productStockUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *productStockUseCase) mustLoadProductStockDatas(ctx context.Context, productStocks []*model.ProductStock, option productStockLoaderParams) {
	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productStocks {
				group.Go(productLoader.ProductStockFn(productStocks[i]))
			}
		}),
	)
}

func (u *productStockUseCase) Fetch(ctx context.Context, request dto_request.ProductStockFetchRequest) ([]model.ProductStock, int) {
	queryOption := model.ProductStockQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	productStocks, err := u.repositoryManager.ProductStockRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductStockRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductStockDatas(ctx, util.SliceValueToSlicePointer(productStocks), productStockLoaderParams{
		product: true,
	})

	return productStocks, total
}

func (u *productStockUseCase) Get(ctx context.Context, request dto_request.ProductStockGetRequest) model.ProductStock {
	productStock := mustGetProductStock(ctx, u.repositoryManager, request.ProductStockId, true)

	u.mustLoadProductStockDatas(ctx, []*model.ProductStock{&productStock}, productStockLoaderParams{
		product: true,
	})

	return productStock
}

func (u *productStockUseCase) Adjustment(ctx context.Context, request dto_request.ProductStockAdjustmentRequest) model.ProductStock {
	productStock := mustGetProductStock(ctx, u.repositoryManager, request.ProductStockId, true)

	productStock.Qty = request.Qty

	panicIfErr(
		u.repositoryManager.ProductStockRepository().Update(ctx, &productStock),
	)

	return productStock
}
