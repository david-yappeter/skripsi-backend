package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type CustomerTypeLoader struct {
	loader dataloader.Loader
}

func (l *CustomerTypeLoader) load(id string) (*model.CustomerType, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.CustomerType), nil
}

func (l *CustomerTypeLoader) CustomerFn(customer *model.Customer) func() error {
	return func() error {
		if customer.CustomerTypeId != nil {
			customerType, err := l.load(*customer.CustomerTypeId)
			if err != nil {
				return err
			}

			customer.CustomerType = customerType
		}

		return nil
	}
}

func NewCustomerTypeLoader(customerTypeRepository repository.CustomerTypeRepository) *CustomerTypeLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		customerTypes, err := customerTypeRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		customerTypeById := map[string]model.CustomerType{}
		for _, customerType := range customerTypes {
			customerTypeById[customerType.Id] = customerType
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var customerType *model.CustomerType
			if v, ok := customerTypeById[k.String()]; ok {
				customerType = &v
			}

			result := &dataloader.Result{Data: customerType, Error: nil}
			if customerType == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &CustomerTypeLoader{
		loader: NewDataloader(batchFn),
	}
}
