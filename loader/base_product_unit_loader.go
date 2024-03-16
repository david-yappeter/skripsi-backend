package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type BaseProductUnitLoader struct {
	loaderByProductId dataloader.Loader
}

func (l *BaseProductUnitLoader) loadByProductId(id string) (*model.ProductUnit, error) {
	thunk := l.loaderByProductId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductUnit), nil
}

func (l *BaseProductUnitLoader) ProductFnNotStrict(product *model.Product) func() error {
	return func() error {
		baseProductUnit, err := l.loadByProductId(product.Id)
		if err != nil && err != constant.ErrNoData {
			return err
		}

		product.BaseProductUnit = baseProductUnit

		return nil
	}
}

func NewBaseProductUnitLoader(baseProductUnitRepository repository.ProductUnitRepository) *BaseProductUnitLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		productIds := make([]string, len(keys))
		for idx, k := range keys {
			productIds[idx] = k.String()
		}

		baseProductUnits, err := baseProductUnitRepository.FetchBaseByProductIds(ctx, productIds)
		if err != nil {
			panic(err)
		}

		baseProductUnitByProductId := map[string]model.ProductUnit{}
		for _, baseProductUnit := range baseProductUnits {
			baseProductUnitByProductId[baseProductUnit.ProductId] = baseProductUnit
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var baseProductUnit *model.ProductUnit
			if v, ok := baseProductUnitByProductId[k.String()]; ok {
				baseProductUnit = &v
			}

			result := &dataloader.Result{Data: baseProductUnit, Error: nil}
			if baseProductUnit == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &BaseProductUnitLoader{
		loaderByProductId: NewDataloader(batchFn),
	}
}
