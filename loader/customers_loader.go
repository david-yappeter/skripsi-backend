package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type CustomersLoader struct {
	loaderByCustomerTypeId dataloader.Loader
}

func (l *CustomersLoader) loadByCustomerTypeId(id string) ([]model.Customer, error) {
	thunk := l.loaderByCustomerTypeId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.Customer), nil
}

func (l *CustomersLoader) CustomerTypeFn(customerType *model.CustomerType) func() error {
	return func() error {
		if customerType != nil {
			customers, err := l.loadByCustomerTypeId(customerType.Id)
			if err != nil {
				return err
			}

			customerType.Customers = customers
		}

		return nil
	}
}

func NewCustomersLoader(customerRepository repository.CustomerRepository) *CustomersLoader {
	batchByCustomerTypeIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		customers, err := customerRepository.FetchByCustomerTypeIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		customerByCustomerTypeId := map[string][]model.Customer{}
		for _, customer := range customers {
			customerByCustomerTypeId[*customer.CustomerTypeId] = append(customerByCustomerTypeId[*customer.CustomerTypeId], customer)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var customers []model.Customer
			if v, ok := customerByCustomerTypeId[k.String()]; ok {
				customers = v
			}

			result := &dataloader.Result{Data: customers, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &CustomersLoader{
		loaderByCustomerTypeId: NewDataloader(batchByCustomerTypeIdFn),
	}
}
