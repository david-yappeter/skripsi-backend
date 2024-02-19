package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductStockLoader struct {
	loaderByProductId dataloader.Loader
}

func (l *ProductStockLoader) loadByProductId(id string) (*model.ProductStock, error) {
	thunk := l.loaderByProductId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductStock), nil
}

func (l *ProductStockLoader) ProductFn(product *model.Product) func() error {
	return func() error {
		productStock, err := l.loadByProductId(product.Id)
		if err != nil {
			return err
		}

		product.ProductStock = productStock

		return nil
	}
}

func (l *ProductStockLoader) ProductUnitFn(productUnit *model.ProductUnit) func() error {
	return func() error {
		productStock, err := l.loadByProductId(productUnit.Id)
		if err != nil {
			return err
		}

		productUnit.ProductStock = productStock

		return nil
	}
}

func NewProductStockLoader(productStockRepository repository.ProductStockRepository) *ProductStockLoader {
	batchByProductIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productStocks, err := productStockRepository.FetchByProductIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productStockByProductId := map[string]model.ProductStock{}
		for _, productStock := range productStocks {
			productStockByProductId[productStock.ProductId] = productStock
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productStock *model.ProductStock
			if v, ok := productStockByProductId[k.String()]; ok {
				productStock = &v
			}

			result := &dataloader.Result{Data: productStock, Error: nil}
			if productStock == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductStockLoader{
		loaderByProductId: NewDataloader(batchByProductIdFn),
	}
}
