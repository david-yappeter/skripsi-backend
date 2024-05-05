package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveReturnLoader struct {
	loaderByProductReceiveId dataloader.Loader
}

func (l *ProductReceiveReturnLoader) loadByProductReceiveId(id string) (*model.ProductReceiveReturn, error) {
	thunk := l.loaderByProductReceiveId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductReceiveReturn), nil
}

func (l *ProductReceiveReturnLoader) ProductReceiveFnNotStrict(productReceive *model.ProductReceive) func() error {
	return func() error {
		productReceiveReturn, err := l.loadByProductReceiveId(productReceive.Id)
		if err != nil && err != constant.ErrNoData {
			return err
		}

		productReceive.ProductReceiveReturn = productReceiveReturn

		return nil
	}
}

func NewProductReceiveReturnLoader(productReceiveReturnRepository repository.ProductReceiveReturnRepository) *ProductReceiveReturnLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceiveReturns, err := productReceiveReturnRepository.FetchByProductReceiveIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveReturnByProductReceiveId := map[string]model.ProductReceiveReturn{}
		for _, productReceiveReturn := range productReceiveReturns {
			productReceiveReturnByProductReceiveId[productReceiveReturn.ProductReceiveId] = productReceiveReturn
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceiveReturn *model.ProductReceiveReturn
			if v, ok := productReceiveReturnByProductReceiveId[k.String()]; ok {
				productReceiveReturn = &v
			}

			result := &dataloader.Result{Data: productReceiveReturn, Error: nil}
			if productReceiveReturn == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveReturnLoader{
		loaderByProductReceiveId: NewDataloader(batchFn),
	}
}
