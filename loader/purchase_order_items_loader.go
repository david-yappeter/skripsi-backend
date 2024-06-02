package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type PurchaseOrderItemsLoader struct {
	loaderByPurchaseOrderId dataloader.Loader
}

func (l *PurchaseOrderItemsLoader) loadByPurchaseOrderId(id string) ([]model.PurchaseOrderItem, error) {
	thunk := l.loaderByPurchaseOrderId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.PurchaseOrderItem), nil
}

func (l *PurchaseOrderItemsLoader) PurchaseOrderFn(purchaseOrder *model.PurchaseOrder) func() error {
	return func() error {
		if purchaseOrder != nil {
			purchaseOrderItems, err := l.loadByPurchaseOrderId(purchaseOrder.Id)
			if err != nil {
				return err
			}

			purchaseOrder.PurchaseOrderItems = purchaseOrderItems
		}

		return nil
	}
}

func NewPurchaseOrderItemsLoader(purchaseOrderItemRepository repository.PurchaseOrderItemRepository) *PurchaseOrderItemsLoader {
	batchByPurchaseOrderIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		purchaseOrderItems, err := purchaseOrderItemRepository.FetchByPurchaseOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		purchaseOrderItemsByPurchaseOrderId := map[string][]model.PurchaseOrderItem{}
		for _, purchaseOrderItem := range purchaseOrderItems {
			purchaseOrderItemsByPurchaseOrderId[purchaseOrderItem.PurchaseOrderId] = append(purchaseOrderItemsByPurchaseOrderId[purchaseOrderItem.PurchaseOrderId], purchaseOrderItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var purchaseOrderItems []model.PurchaseOrderItem
			if v, ok := purchaseOrderItemsByPurchaseOrderId[k.String()]; ok {
				purchaseOrderItems = v
			}

			result := &dataloader.Result{Data: purchaseOrderItems, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &PurchaseOrderItemsLoader{
		loaderByPurchaseOrderId: NewDataloader(batchByPurchaseOrderIdFn),
	}
}
