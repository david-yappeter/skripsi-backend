package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type TransactionPaymentsLoader struct {
	loaderByTransactionId dataloader.Loader
}

func (l *TransactionPaymentsLoader) loadByTransactionId(id string) ([]model.TransactionPayment, error) {
	thunk := l.loaderByTransactionId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.TransactionPayment), nil
}

func (l *TransactionPaymentsLoader) TransactionFn(transaction *model.Transaction) func() error {
	return func() error {
		if transaction != nil {
			transactionPayments, err := l.loadByTransactionId(transaction.Id)
			if err != nil {
				return err
			}

			transaction.TransactionPayments = transactionPayments
		}

		return nil
	}
}

func NewTransactionPaymentsLoader(transactionPaymentRepository repository.TransactionPaymentRepository) *TransactionPaymentsLoader {
	batchByTransactionIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		transactionPayments, err := transactionPaymentRepository.FetchByTransactionIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		transactionPaymentByTransactionId := map[string][]model.TransactionPayment{}
		for _, transactionPayment := range transactionPayments {
			transactionPaymentByTransactionId[transactionPayment.TransactionId] = append(transactionPaymentByTransactionId[transactionPayment.TransactionId], transactionPayment)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var transactionPayments []model.TransactionPayment
			if v, ok := transactionPaymentByTransactionId[k.String()]; ok {
				transactionPayments = v
			}

			result := &dataloader.Result{Data: transactionPayments, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &TransactionPaymentsLoader{
		loaderByTransactionId: NewDataloader(batchByTransactionIdFn),
	}
}
