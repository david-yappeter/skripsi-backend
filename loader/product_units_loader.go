package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductUnitsLoader struct {
	loaderByProductId dataloader.Loader
}

func (l *ProductUnitsLoader) loadByProductId(id string) ([]model.ProductUnit, error) {
	thunk := l.loaderByProductId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductUnit), nil
}

func (l *ProductUnitsLoader) ProductFn(product *model.Product) func() error {
	return func() error {
		if product != nil {
			productUnits, err := l.loadByProductId(product.Id)
			if err != nil {
				return err
			}

			product.ProductUnits = productUnits
		}

		return nil
	}
}

func NewProductUnitsLoader(productUnitRepository repository.ProductUnitRepository) *ProductUnitsLoader {
	batchByUserIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productUnits, err := productUnitRepository.FetchByProductIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productUnitsByProductId := map[string][]model.ProductUnit{}
		for _, productUnit := range productUnits {
			productUnitsByProductId[productUnit.ProductId] = append(productUnitsByProductId[productUnit.ProductId], productUnit)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productUnits []model.ProductUnit
			if v, ok := productUnitsByProductId[k.String()]; ok {
				productUnits = v
			}

			result := &dataloader.Result{Data: productUnits, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ProductUnitsLoader{
		loaderByProductId: NewDataloader(batchByUserIdFn),
	}
}
