package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveItemsLoader struct {
	loaderByProductReceiveId dataloader.Loader
}

func (l *ProductReceiveItemsLoader) loadByProductReceiveId(id string) ([]model.ProductReceiveItem, error) {
	thunk := l.loaderByProductReceiveId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReceiveItem), nil
}

func (l *ProductReceiveItemsLoader) ProductReceiveFn(productReceive *model.ProductReceive) func() error {
	return func() error {
		if productReceive != nil {
			productReceiveItems, err := l.loadByProductReceiveId(productReceive.Id)
			if err != nil {
				return err
			}

			productReceive.ProductReceiveItems = productReceiveItems
		}

		return nil
	}
}

func NewProductReceiveItemsLoader(productReceiveItemRepository repository.ProductReceiveItemRepository) *ProductReceiveItemsLoader {
	batchByProductReceiveIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceiveItems, err := productReceiveItemRepository.FetchByProductReceiveIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveItemsByProductReceiveId := map[string][]model.ProductReceiveItem{}
		for _, productReceiveItem := range productReceiveItems {
			productReceiveItemsByProductReceiveId[productReceiveItem.ProductReceiveId] = append(productReceiveItemsByProductReceiveId[productReceiveItem.ProductReceiveId], productReceiveItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceiveItems []model.ProductReceiveItem
			if v, ok := productReceiveItemsByProductReceiveId[k.String()]; ok {
				productReceiveItems = v
			}

			result := &dataloader.Result{Data: productReceiveItems, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveItemsLoader{
		loaderByProductReceiveId: NewDataloader(batchByProductReceiveIdFn),
	}
}
