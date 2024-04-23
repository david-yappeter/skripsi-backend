package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderLoader struct {
	loader dataloader.Loader
}

func (l *DeliveryOrderLoader) load(id string) (*model.DeliveryOrder, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.DeliveryOrder), nil
}

func (l *DeliveryOrderLoader) DeliveryOrderReviewFn(deliveryOrderReview *model.DeliveryOrderReview) func() error {
	return func() error {
		DeliveryOrder, err := l.load(deliveryOrderReview.DeliveryOrderId)
		if err != nil {
			return err
		}

		deliveryOrderReview.DeliveryOrder = DeliveryOrder

		return nil
	}
}

func NewDeliveryOrderLoader(deliveryOrderRepository repository.DeliveryOrderRepository) *DeliveryOrderLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		DeliveryOrders, err := deliveryOrderRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		DeliveryOrderById := map[string]model.DeliveryOrder{}
		for _, DeliveryOrder := range DeliveryOrders {
			DeliveryOrderById[DeliveryOrder.Id] = DeliveryOrder
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var DeliveryOrder *model.DeliveryOrder
			if v, ok := DeliveryOrderById[k.String()]; ok {
				DeliveryOrder = &v
			}

			result := &dataloader.Result{Data: DeliveryOrder, Error: nil}
			if DeliveryOrder == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderLoader{
		loader: NewDataloader(batchFn),
	}
}
