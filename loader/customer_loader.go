package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type CustomerLoader struct {
	loader dataloader.Loader
}

func (l *CustomerLoader) load(id string) (*model.Customer, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.Customer), nil
}

func (l *CustomerLoader) CustomerDebtFn(customerDebt *model.CustomerDebt) func() error {
	return func() error {
		customer, err := l.load(customerDebt.CustomerId)
		if err != nil {
			return err
		}

		customerDebt.Customer = customer

		return nil
	}
}

func NewCustomerLoader(customerRepository repository.CustomerRepository) *CustomerLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		customers, err := customerRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		customerById := map[string]model.Customer{}
		for _, customer := range customers {
			customerById[customer.Id] = customer
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var customer *model.Customer
			if v, ok := customerById[k.String()]; ok {
				customer = &v
			}

			result := &dataloader.Result{Data: customer, Error: nil}
			if customer == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &CustomerLoader{
		loader: NewDataloader(batchFn),
	}
}
