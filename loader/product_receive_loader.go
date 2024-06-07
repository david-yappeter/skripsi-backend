package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveLoader struct {
	loader dataloader.Loader
}

func (l *ProductReceiveLoader) load(id string) (*model.ProductReceive, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.ProductReceive), nil
}

func (l *ProductReceiveLoader) DebtFn(debt *model.Debt) func() error {
	return func() error {
		productReceive, err := l.load(debt.DebtSourceIdentifier)
		if err != nil {
			return err
		}

		debt.ProductReceive = productReceive

		return nil
	}
}

func NewProductReceiveLoader(productReceiveRepository repository.ProductReceiveRepository) *ProductReceiveLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceives, err := productReceiveRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveById := map[string]model.ProductReceive{}
		for _, productReceive := range productReceives {
			productReceiveById[productReceive.Id] = productReceive
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceive *model.ProductReceive
			if v, ok := productReceiveById[k.String()]; ok {
				productReceive = &v
			}

			result := &dataloader.Result{Data: productReceive, Error: nil}
			if productReceive == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveLoader{
		loader: NewDataloader(batchFn),
	}
}
