package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReturnItemsLoader struct {
	loaderByProductReturnId dataloader.Loader
}

func (l *ProductReturnItemsLoader) loadByProductReturnId(id string) ([]model.ProductReturnItem, error) {
	thunk := l.loaderByProductReturnId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReturnItem), nil
}

func (l *ProductReturnItemsLoader) ProductReturnFn(productReturn *model.ProductReturn) func() error {
	return func() error {
		if productReturn != nil {
			productReturnItems, err := l.loadByProductReturnId(productReturn.Id)
			if err != nil {
				return err
			}

			productReturn.ProductReturnItems = productReturnItems
		}

		return nil
	}
}

func NewProductReturnItemsLoader(productReturnItemRepository repository.ProductReturnItemRepository) *ProductReturnItemsLoader {
	batchByProductReturnIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReturnItems, err := productReturnItemRepository.FetchByProductReturnIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReturnItemsByProductReturnId := map[string][]model.ProductReturnItem{}
		for _, productReturnItem := range productReturnItems {
			productReturnItemsByProductReturnId[productReturnItem.ProductReturnId] = append(productReturnItemsByProductReturnId[productReturnItem.ProductReturnId], productReturnItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReturnItems []model.ProductReturnItem
			if v, ok := productReturnItemsByProductReturnId[k.String()]; ok {
				productReturnItems = v
			}

			result := &dataloader.Result{Data: productReturnItems, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ProductReturnItemsLoader{
		loaderByProductReturnId: NewDataloader(batchByProductReturnIdFn),
	}
}
