package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductUnitLoader struct {
	loader dataloader.Loader
}

func (l *ProductUnitLoader) load(id string) (*model.ProductUnit, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductUnit), nil
}

func (l *ProductUnitLoader) ProductReceiveItemFn(productReceiveItem *model.ProductReceiveItem) func() error {
	return func() error {
		productUnit, err := l.load(productReceiveItem.ProductUnitId)
		if err != nil {
			return err
		}

		productReceiveItem.ProductUnit = productUnit

		return nil
	}
}

func (l *ProductUnitLoader) DeliveryOrderItemFn(deliveryOrderItem *model.DeliveryOrderItem) func() error {
	return func() error {
		productUnit, err := l.load(deliveryOrderItem.ProductUnitId)
		if err != nil {
			return err
		}

		deliveryOrderItem.ProductUnit = productUnit

		return nil
	}
}

func NewProductUnitLoader(productUnitRepository repository.ProductUnitRepository) *ProductUnitLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productUnits, err := productUnitRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productUnitById := map[string]model.ProductUnit{}
		for _, productUnit := range productUnits {
			productUnitById[productUnit.Id] = productUnit
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productUnit *model.ProductUnit
			if v, ok := productUnitById[k.String()]; ok {
				productUnit = &v
			}

			result := &dataloader.Result{Data: productUnit, Error: nil}
			if productUnit == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductUnitLoader{
		loader: NewDataloader(batchFn),
	}
}
