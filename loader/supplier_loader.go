package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type SupplierLoader struct {
	loader dataloader.Loader
}

func (l *SupplierLoader) load(id string) (*model.Supplier, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Supplier), nil
}

func (l *SupplierLoader) ProductReceiveFn(productReceive *model.ProductReceive) func() error {
	return func() error {
		supplier, err := l.load(productReceive.SupplierId)
		if err != nil {
			return err
		}

		productReceive.Supplier = supplier

		return nil
	}
}

func (l *SupplierLoader) PurchaseOrderFn(purchaseOrder *model.PurchaseOrder) func() error {
	return func() error {
		supplier, err := l.load(purchaseOrder.SupplierId)
		if err != nil {
			return err
		}

		purchaseOrder.Supplier = supplier

		return nil
	}
}

func (l *SupplierLoader) ProductReturnFn(productReturn *model.ProductReturn) func() error {
	return func() error {
		supplier, err := l.load(productReturn.SupplierId)
		if err != nil {
			return err
		}

		productReturn.Supplier = supplier

		return nil
	}
}

func NewSupplierLoader(supplierRepository repository.SupplierRepository) *SupplierLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		suppliers, err := supplierRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		supplierById := map[string]model.Supplier{}
		for _, supplier := range suppliers {
			supplierById[supplier.Id] = supplier
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var supplier *model.Supplier
			if v, ok := supplierById[k.String()]; ok {
				supplier = &v
			}

			result := &dataloader.Result{Data: supplier, Error: nil}
			if supplier == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &SupplierLoader{
		loader: NewDataloader(batchFn),
	}
}
