package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type TransactionItemCostsLoader struct {
	loaderByTransactionItemId dataloader.Loader
}

func (l *TransactionItemCostsLoader) loadByTransactionItemId(id string) ([]model.TransactionItemCost, error) {
	thunk := l.loaderByTransactionItemId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.TransactionItemCost), nil
}

func (l *TransactionItemCostsLoader) TransactionItemFn(transactionItem *model.TransactionItem) func() error {
	return func() error {
		if transactionItem != nil {
			transactionItemCosts, err := l.loadByTransactionItemId(transactionItem.Id)
			if err != nil {
				return err
			}

			transactionItem.TransactionItemCosts = transactionItemCosts
		}

		return nil
	}
}

func NewTransactionItemCostsLoader(transactionItemCostRepository repository.TransactionItemCostRepository) *TransactionItemCostsLoader {
	batchByTransactionItemIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		transactionItemCosts, err := transactionItemCostRepository.FetchByTransactionItemIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		transactionItemCostByTransactionItemId := map[string][]model.TransactionItemCost{}
		for _, transactionItemCost := range transactionItemCosts {
			transactionItemCostByTransactionItemId[transactionItemCost.TransactionItemId] = append(transactionItemCostByTransactionItemId[transactionItemCost.TransactionItemId], transactionItemCost)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var transactionItemCosts []model.TransactionItemCost
			if v, ok := transactionItemCostByTransactionItemId[k.String()]; ok {
				transactionItemCosts = v
			}

			result := &dataloader.Result{Data: transactionItemCosts, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &TransactionItemCostsLoader{
		loaderByTransactionItemId: NewDataloader(batchByTransactionItemIdFn),
	}
}
