package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type SupplierTypeLoader struct {
	loader dataloader.Loader
}

func (l *SupplierTypeLoader) load(id string) (*model.SupplierType, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.SupplierType), nil
}

func (l *SupplierTypeLoader) SupplierFn(supplier *model.Supplier) func() error {
	return func() error {
		supplierType, err := l.load(supplier.SupplierTypeId)
		if err != nil {
			return err
		}

		supplier.SupplierType = supplierType

		return nil
	}
}

func NewSupplierTypeLoader(supplierTypeRepository repository.SupplierTypeRepository) *SupplierTypeLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		supplierTypes, err := supplierTypeRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		supplierTypeById := map[string]model.SupplierType{}
		for _, supplierType := range supplierTypes {
			supplierTypeById[supplierType.Id] = supplierType
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var supplierType *model.SupplierType
			if v, ok := supplierTypeById[k.String()]; ok {
				supplierType = &v
			}

			result := &dataloader.Result{Data: supplierType, Error: nil}
			if supplierType == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &SupplierTypeLoader{
		loader: NewDataloader(batchFn),
	}
}
