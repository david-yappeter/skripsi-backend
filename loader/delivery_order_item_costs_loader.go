package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderItemCostsLoader struct {
	loaderByDeliveryOrderItem dataloader.Loader
}

func (l *DeliveryOrderItemCostsLoader) loadByDeliveryOrderItem(id string) ([]model.DeliveryOrderItemCost, error) {
	thunk := l.loaderByDeliveryOrderItem.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.DeliveryOrderItemCost), nil
}

func (l *DeliveryOrderItemCostsLoader) DeliveryOrderItemFn(deliveryOrderItem *model.DeliveryOrderItem) func() error {
	return func() error {
		if deliveryOrderItem != nil {
			deliveryOrderItemCosts, err := l.loadByDeliveryOrderItem(deliveryOrderItem.Id)
			if err != nil {
				return err
			}

			deliveryOrderItem.DeliveryOrderItemCosts = deliveryOrderItemCosts
		}

		return nil
	}
}

func NewDeliveryOrderItemCostsLoader(deliveryOrderItemCostRepository repository.DeliveryOrderItemCostRepository) *DeliveryOrderItemCostsLoader {
	batchByDeliveryOrderItemIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceives, err := deliveryOrderItemCostRepository.FetchByDeliveryOrderItemIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderItemCostByUserId := map[string][]model.DeliveryOrderItemCost{}
		for _, deliveryOrderItemCost := range productReceives {
			deliveryOrderItemCostByUserId[deliveryOrderItemCost.DeliveryOrderItemId] = append(deliveryOrderItemCostByUserId[deliveryOrderItemCost.DeliveryOrderItemId], deliveryOrderItemCost)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderItemCosts []model.DeliveryOrderItemCost
			if v, ok := deliveryOrderItemCostByUserId[k.String()]; ok {
				deliveryOrderItemCosts = v
			}

			result := &dataloader.Result{Data: deliveryOrderItemCosts, Error: nil}
			if deliveryOrderItemCosts == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderItemCostsLoader{
		loaderByDeliveryOrderItem: NewDataloader(batchByDeliveryOrderItemIdFn),
	}
}
