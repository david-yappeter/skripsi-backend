package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DebtPaymentsLoader struct {
	loaderByDebtId dataloader.Loader
}

func (l *DebtPaymentsLoader) loadByDebtId(id string) ([]model.DebtPayment, error) {
	thunk := l.loaderByDebtId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.DebtPayment), nil
}

func (l *DebtPaymentsLoader) DebtFn(debt *model.Debt) func() error {
	return func() error {
		if debt != nil {
			debtPayments, err := l.loadByDebtId(debt.Id)
			if err != nil {
				return err
			}

			debt.DebtPayments = debtPayments
		}

		return nil
	}
}

func NewDebtPaymentsLoader(debtPaymentRepository repository.DebtPaymentRepository) *DebtPaymentsLoader {
	batchByDebtIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		debtPayments, err := debtPaymentRepository.FetchByDebtIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		debtPaymentByDebtId := map[string][]model.DebtPayment{}
		for _, debtPayment := range debtPayments {
			debtPaymentByDebtId[debtPayment.DebtId] = append(debtPaymentByDebtId[debtPayment.DebtId], debtPayment)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var debtPayments []model.DebtPayment
			if v, ok := debtPaymentByDebtId[k.String()]; ok {
				debtPayments = v
			}

			result := &dataloader.Result{Data: debtPayments, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &DebtPaymentsLoader{
		loaderByDebtId: NewDataloader(batchByDebtIdFn),
	}
}
