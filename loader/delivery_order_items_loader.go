package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderItemsLoader struct {
	loader dataloader.Loader
}

func (l *DeliveryOrderItemsLoader) load(id string) ([]model.DeliveryOrderItem, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.DeliveryOrderItem), nil
}

func (l *DeliveryOrderItemsLoader) DeliveryOrderFn(deliveryOrder *model.DeliveryOrder) func() error {
	return func() error {
		if deliveryOrder != nil {
			deliveryOrderItems, err := l.load(deliveryOrder.Id)
			if err != nil {
				return err
			}

			deliveryOrder.DeliveryOrderItems = deliveryOrderItems
		}

		return nil
	}
}

func NewDeliveryOrderItemsLoader(deliveryOrderItemRepository repository.DeliveryOrderItemRepository) *DeliveryOrderItemsLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		deliveryOrderItems, err := deliveryOrderItemRepository.FetchByDeliveryOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderItemsByProductReceiveId := map[string][]model.DeliveryOrderItem{}
		for _, deliveryOrderItem := range deliveryOrderItems {
			deliveryOrderItemsByProductReceiveId[deliveryOrderItem.DeliveryOrderId] = append(deliveryOrderItemsByProductReceiveId[deliveryOrderItem.DeliveryOrderId], deliveryOrderItem)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderItems []model.DeliveryOrderItem
			if v, ok := deliveryOrderItemsByProductReceiveId[k.String()]; ok {
				deliveryOrderItems = v
			}

			result := &dataloader.Result{Data: deliveryOrderItems, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderItemsLoader{
		loader: NewDataloader(batchFn),
	}
}
