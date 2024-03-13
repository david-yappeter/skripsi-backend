package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ShopOrderItemsLoader struct {
	loaderByShopOrderId dataloader.Loader
}

func (l *ShopOrderItemsLoader) loadByShopOrderId(id string) ([]model.ShopOrderItem, error) {
	thunk := l.loaderByShopOrderId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ShopOrderItem), nil
}

func (l *ShopOrderItemsLoader) ShopOrderFn(shopOrder *model.ShopOrder) func() error {
	return func() error {
		if shopOrder != nil {
			shopOrderItems, err := l.loadByShopOrderId(shopOrder.Id)
			if err != nil {
				return err
			}

			shopOrder.ShopOrderItems = shopOrderItems
		}

		return nil
	}
}

func NewShopOrderItemsLoader(shopOrderItemRepository repository.ShopOrderItemRepository) *ShopOrderItemsLoader {
	batchByShopOrderIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		shopOrderItems, err := shopOrderItemRepository.FetchByShopOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		shopOrderItemByShopOrderId := map[string][]model.ShopOrderItem{}
		for _, shopOrderItem := range shopOrderItems {
			shopOrderItemByShopOrderId[shopOrderItem.ShopOrderId] = append(shopOrderItemByShopOrderId[shopOrderItem.ShopOrderId], shopOrderItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var shopOrderItems []model.ShopOrderItem
			if v, ok := shopOrderItemByShopOrderId[k.String()]; ok {
				shopOrderItems = v
			}

			result := &dataloader.Result{Data: shopOrderItems, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ShopOrderItemsLoader{
		loaderByShopOrderId: NewDataloader(batchByShopOrderIdFn),
	}
}
