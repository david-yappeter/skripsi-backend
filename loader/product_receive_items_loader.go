package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveItemsLoader struct {
	loader dataloader.Loader
}

func (l *ProductReceiveItemsLoader) load(id string) ([]model.ProductReceiveItem, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReceiveItem), nil
}

func (l *ProductReceiveItemsLoader) ProductReceiveFn(productReceive *model.ProductReceive) func() error {
	return func() error {
		if productReceive != nil {
			productReceiveItems, err := l.load(productReceive.Id)
			if err != nil {
				return err
			}

			productReceive.ProductReceiveItems = productReceiveItems
		}

		return nil
	}
}

func NewProductReceiveItemsLoader(productReceiveItemRepository repository.ProductReceiveItemRepository) *ProductReceiveItemsLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceives, err := productReceiveItemRepository.FetchByProductReceiveIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveItemsByProductReceiveId := map[string][]model.ProductReceiveItem{}
		for _, productReceive := range productReceives {
			productReceiveItemsByProductReceiveId[productReceive.Id] = append(productReceiveItemsByProductReceiveId[productReceive.Id], productReceive)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceiveItems []model.ProductReceiveItem
			if v, ok := productReceiveItemsByProductReceiveId[k.String()]; ok {
				productReceiveItems = v
			}

			result := &dataloader.Result{Data: productReceiveItems, Error: nil}
			if productReceiveItems == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveItemsLoader{
		loader: NewDataloader(batchFn),
	}
}
