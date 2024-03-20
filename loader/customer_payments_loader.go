package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type CustomerPaymentsLoader struct {
	loaderByCustomerDebtId dataloader.Loader
}

func (l *CustomerPaymentsLoader) loadByCustomerDebtId(id string) ([]model.CustomerPayment, error) {
	thunk := l.loaderByCustomerDebtId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.CustomerPayment), nil
}

func (l *CustomerPaymentsLoader) CustomerDebtFn(customerDebt *model.CustomerDebt) func() error {
	return func() error {
		if customerDebt != nil {
			customerPayments, err := l.loadByCustomerDebtId(customerDebt.Id)
			if err != nil {
				return err
			}

			customerDebt.CustomerPayments = customerPayments
		}

		return nil
	}
}

func NewCustomerPaymentsLoader(customerPaymentRepository repository.CustomerPaymentRepository) *CustomerPaymentsLoader {
	batchByCustomerDebtIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		customerPayments, err := customerPaymentRepository.FetchByCustomerDebtIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		customerPaymentByCustomerDebtId := map[string][]model.CustomerPayment{}
		for _, customerPayment := range customerPayments {
			customerPaymentByCustomerDebtId[customerPayment.CustomerDebtId] = append(customerPaymentByCustomerDebtId[customerPayment.CustomerDebtId], customerPayment)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var customerPayments []model.CustomerPayment
			if v, ok := customerPaymentByCustomerDebtId[k.String()]; ok {
				customerPayments = v
			}

			result := &dataloader.Result{Data: customerPayments, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &CustomerPaymentsLoader{
		loaderByCustomerDebtId: NewDataloader(batchByCustomerDebtIdFn),
	}
}
