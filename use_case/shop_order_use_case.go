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

type shopOrderLoaderParams struct {
	shopOrderItems bool
}

type ShopOrderUseCase interface {
	Fetch(ctx context.Context, request dto_request.ShopOrderFetchRequest) ([]model.ShopOrder, int)
	Get(ctx context.Context, request dto_request.ShopOrderGetRequest) model.ShopOrder
}

type shopOrderUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewShopOrderUseCase(
	repositoryManager repository.RepositoryManager,
) ShopOrderUseCase {
	return &shopOrderUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *shopOrderUseCase) mustLoadShopOrderData(ctx context.Context, shopOrders []*model.ShopOrder, option shopOrderLoaderParams) {
	shopOrderItemsLoader := loader.NewShopOrderItemsLoader(u.repositoryManager.ShopOrderItemRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			if option.shopOrderItems {
				for i := range shopOrders {
					group.Go(shopOrderItemsLoader.ShopOrderFn(shopOrders[i]))
				}
			}
		}),
	)

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			if option.shopOrderItems {
				for i := range shopOrders {
					for j := range shopOrders[i].ShopOrderItems {
						group.Go(productUnitLoader.ShopOrderItemFn(&shopOrders[i].ShopOrderItems[j]))
					}
				}
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			if option.shopOrderItems {
				for i := range shopOrders {
					for j := range shopOrders[i].ShopOrderItems {
						group.Go(productLoader.ProductUnitFn(shopOrders[i].ShopOrderItems[j].ProductUnit))
						group.Go(unitLoader.ProductUnitFn(shopOrders[i].ShopOrderItems[j].ProductUnit))
					}
				}
			}
		}),
	)
}

func (u *shopOrderUseCase) Fetch(ctx context.Context, request dto_request.ShopOrderFetchRequest) ([]model.ShopOrder, int) {
	queryOption := model.ShopOrderQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{{Field: "updated_at", Direction: "desc"}},
		),
		Phrase:         request.Phrase,
		TrackingStatus: request.TrackingStatus,
		PlatformType:   request.PlatformType,
	}

	shopOrders, err := u.repositoryManager.ShopOrderRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ShopOrderRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadShopOrderData(ctx, util.SliceValueToSlicePointer(shopOrders), shopOrderLoaderParams{
		shopOrderItems: request.WithItems,
	})

	return shopOrders, total
}

func (u *shopOrderUseCase) Get(ctx context.Context, request dto_request.ShopOrderGetRequest) model.ShopOrder {
	shopOrder := mustGetShopOrder(ctx, u.repositoryManager, request.ShopOrderId, false)

	u.mustLoadShopOrderData(ctx, []*model.ShopOrder{&shopOrder}, shopOrderLoaderParams{
		shopOrderItems: true,
	})

	return shopOrder
}
