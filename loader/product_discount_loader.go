package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductDiscountLoader struct {
	loaderByProductId dataloader.Loader
}

func (l *ProductDiscountLoader) loadByProductId(id string) (*model.ProductDiscount, error) {
	thunk := l.loaderByProductId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductDiscount), nil
}

func (l *ProductDiscountLoader) ProductFnNotStrict(product *model.Product) func() error {
	return func() error {
		productDiscount, err := l.loadByProductId(product.Id)
		if err != nil && err != constant.ErrNoData {
			return err
		}

		product.ProductDiscount = productDiscount

		return nil
	}
}

func NewProductDiscountLoader(productDiscountRepository repository.ProductDiscountRepository) *ProductDiscountLoader {
	batchFnByProductId := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		productIds := make([]string, len(keys))
		for idx, k := range keys {
			productIds[idx] = k.String()
		}

		productDiscounts, err := productDiscountRepository.FetchByProductIds(ctx, productIds)
		if err != nil {
			panic(err)
		}

		productDiscountByProductId := map[string]model.ProductDiscount{}
		for _, productDiscount := range productDiscounts {
			productDiscountByProductId[productDiscount.ProductId] = productDiscount
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productDiscount *model.ProductDiscount
			if v, ok := productDiscountByProductId[k.String()]; ok {
				productDiscount = &v
			}

			result := &dataloader.Result{Data: productDiscount, Error: nil}
			if productDiscount == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductDiscountLoader{
		loaderByProductId: NewDataloader(batchFnByProductId),
	}
}
