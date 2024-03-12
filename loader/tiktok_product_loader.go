package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type TiktokProductLoader struct {
	loaderByProductId dataloader.Loader
}

func (l *TiktokProductLoader) loadByProductId(id string) (*model.TiktokProduct, error) {
	thunk := l.loaderByProductId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.TiktokProduct), nil
}

func (l *TiktokProductLoader) ProductFnNotStrict(product *model.Product) func() error {
	return func() error {
		tiktokProduct, err := l.loadByProductId(product.Id)
		if err != nil && err != constant.ErrNoData {
			return err
		}

		product.TiktokProduct = tiktokProduct

		return nil
	}
}

func NewTiktokProductLoader(tiktokProductRepository repository.TiktokProductRepository) *TiktokProductLoader {
	batchByTiktokProductIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		tiktokProducts, err := tiktokProductRepository.FetchByProductIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		tiktokProductByTiktokProductId := map[string]model.TiktokProduct{}
		for _, tiktokProduct := range tiktokProducts {
			tiktokProductByTiktokProductId[tiktokProduct.ProductId] = tiktokProduct
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var tiktokProduct *model.TiktokProduct
			if v, ok := tiktokProductByTiktokProductId[k.String()]; ok {
				tiktokProduct = &v
			}

			result := &dataloader.Result{Data: tiktokProduct, Error: nil}
			if tiktokProduct == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &TiktokProductLoader{
		loaderByProductId: NewDataloader(batchByTiktokProductIdFn),
	}
}
