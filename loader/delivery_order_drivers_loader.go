package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderDriversLoader struct {
	loaderByDeliveryOrderId dataloader.Loader
}

func (l *DeliveryOrderDriversLoader) loadByDeliveryOrderId(id string) ([]model.DeliveryOrderDriver, error) {
	thunk := l.loaderByDeliveryOrderId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.DeliveryOrderDriver), nil
}

func (l *DeliveryOrderDriversLoader) DeliveryOrderFn(deliveryOrder *model.DeliveryOrder) func() error {
	return func() error {
		if deliveryOrder != nil {
			deliveryOrderDrivers, err := l.loadByDeliveryOrderId(deliveryOrder.Id)
			if err != nil {
				return err
			}

			deliveryOrder.DeliveryOrderDrivers = deliveryOrderDrivers
		}

		return nil
	}
}

func NewDeliveryOrderDriversLoader(deliveryOrderDriverRepository repository.DeliveryOrderDriverRepository) *DeliveryOrderDriversLoader {
	batchByDeliveryOrderIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		deliveryOrderDrivers, err := deliveryOrderDriverRepository.FetchByDeliveryOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderDriverByDeliveryOrderId := map[string][]model.DeliveryOrderDriver{}
		for _, deliveryOrderDriver := range deliveryOrderDrivers {
			deliveryOrderDriverByDeliveryOrderId[deliveryOrderDriver.DeliveryOrderId] = append(deliveryOrderDriverByDeliveryOrderId[deliveryOrderDriver.DeliveryOrderId], deliveryOrderDriver)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderDrivers []model.DeliveryOrderDriver
			if v, ok := deliveryOrderDriverByDeliveryOrderId[k.String()]; ok {
				deliveryOrderDrivers = v
			}

			result := &dataloader.Result{Data: deliveryOrderDrivers, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderDriversLoader{
		loaderByDeliveryOrderId: NewDataloader(batchByDeliveryOrderIdFn),
	}
}
