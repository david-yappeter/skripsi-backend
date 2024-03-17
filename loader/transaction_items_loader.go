package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type TransactionItemsLoader struct {
	loaderByTransactionId dataloader.Loader
}

func (l *TransactionItemsLoader) loadByTransactionId(id string) ([]model.TransactionItem, error) {
	thunk := l.loaderByTransactionId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.TransactionItem), nil
}

func (l *TransactionItemsLoader) TransactionFn(transaction *model.Transaction) func() error {
	return func() error {
		if transaction != nil {
			transactionItems, err := l.loadByTransactionId(transaction.Id)
			if err != nil {
				return err
			}

			transaction.TransactionItems = transactionItems
		}

		return nil
	}
}

func NewTransactionItemsLoader(transactionItemRepository repository.TransactionItemRepository) *TransactionItemsLoader {
	batchByTransactionIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		transactionItems, err := transactionItemRepository.FetchByTransactionIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		transactionItemByTransactionId := map[string][]model.TransactionItem{}
		for _, transactionItem := range transactionItems {
			transactionItemByTransactionId[transactionItem.TransactionId] = append(transactionItemByTransactionId[transactionItem.TransactionId], transactionItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var transactionItems []model.TransactionItem
			if v, ok := transactionItemByTransactionId[k.String()]; ok {
				transactionItems = v
			}

			result := &dataloader.Result{Data: transactionItems, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &TransactionItemsLoader{
		loaderByTransactionId: NewDataloader(batchByTransactionIdFn),
	}
}
