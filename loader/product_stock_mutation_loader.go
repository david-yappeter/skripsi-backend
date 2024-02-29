package loader

import (
	"context"
	"myapp/constant"
	"myapp/data_type"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductStockMutationLoader struct {
	loaderByProductReceiveItemIdType dataloader.Loader
}

func (l *ProductStockMutationLoader) loadByProductReceiveItemIdType(id string) (*model.ProductStockMutation, error) {
	thunk := l.loaderByProductReceiveItemIdType.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductStockMutation), nil
}

func (l *ProductStockMutationLoader) ProductReceiveItemFn(productReceiveItem *model.ProductReceiveItem) func() error {
	return func() error {
		productStockMutation, err := l.loadByProductReceiveItemIdType(productReceiveItem.Id)
		if err != nil {
			return err
		}

		productReceiveItem.ProductStockMutation = productStockMutation

		return nil
	}
}

func (l *ProductStockMutationLoader) ProductReceiveItemNotStrictFn(productReceiveItem *model.ProductReceiveItem) func() error {
	return func() error {
		productStockMutation, err := l.loadByProductReceiveItemIdType(productReceiveItem.Id)
		if err != nil && err != constant.ErrNoData {
			return err
		}

		productReceiveItem.ProductStockMutation = productStockMutation

		return nil
	}
}

func NewProductStockMutationLoader(productStockMutationRepository repository.ProductStockMutationRepository) *ProductStockMutationLoader {
	batchByProductReceiveItemTypeFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productStockMutations, err := productStockMutationRepository.FetchByTypeAndIdentifierIds(ctx, data_type.ProductStockMutationTypeProductReceiveItem, ids)
		if err != nil {
			panic(err)
		}

		productStockMutationByIdentifierId := map[string]model.ProductStockMutation{}
		for _, productStockMutation := range productStockMutations {
			productStockMutationByIdentifierId[productStockMutation.IdentifierId] = productStockMutation
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productStockMutation *model.ProductStockMutation
			if v, ok := productStockMutationByIdentifierId[k.String()]; ok {
				productStockMutation = &v
			}

			result := &dataloader.Result{Data: productStockMutation, Error: nil}
			if productStockMutation == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductStockMutationLoader{
		loaderByProductReceiveItemIdType: NewDataloader(batchByProductReceiveItemTypeFn),
	}
}
