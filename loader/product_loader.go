package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductLoader struct {
	loaderById dataloader.Loader
}

func (l *ProductLoader) loadById(id string) (*model.Product, error) {
	thunk := l.loaderById.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Product), nil
}

func (l *ProductLoader) ProductStockFn(productStock *model.ProductStock) func() error {
	return func() error {
		product, err := l.loadById(productStock.ProductId)
		if err != nil {
			return err
		}

		productStock.Product = product

		return nil
	}
}

func NewProductLoader(productRepository repository.ProductRepository) *ProductLoader {
	batchByProductIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		products, err := productRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productByProductId := map[string]model.Product{}
		for _, product := range products {
			productByProductId[product.Id] = product
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var product *model.Product
			if v, ok := productByProductId[k.String()]; ok {
				product = &v
			}

			result := &dataloader.Result{Data: product, Error: nil}
			if product == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductLoader{
		loaderById: NewDataloader(batchByProductIdFn),
	}
}
