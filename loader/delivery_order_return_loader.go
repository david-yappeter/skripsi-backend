package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderReturnLoader struct {
	loaderByDeliveryOrderId dataloader.Loader
}

func (l *DeliveryOrderReturnLoader) loadByDeliveryOrderId(id string) (*model.DeliveryOrderReturn, error) {
	thunk := l.loaderByDeliveryOrderId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.DeliveryOrderReturn), nil
}

func (l *DeliveryOrderReturnLoader) DeliveryOrderFnNotStrict(deliveryOrder *model.DeliveryOrder) func() error {
	return func() error {
		deliveryOrderReturn, err := l.loadByDeliveryOrderId(deliveryOrder.Id)
		if err != nil && err != constant.ErrNoData {
			return err
		}

		deliveryOrder.DeliveryOrderReturn = deliveryOrderReturn

		return nil
	}
}

func NewDeliveryOrderReturnLoader(deliveryOrderReturnRepository repository.DeliveryOrderReturnRepository) *DeliveryOrderReturnLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		deliveryOrderReturns, err := deliveryOrderReturnRepository.FetchByDeliveryOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderReturnByDeliveryOrderId := map[string]model.DeliveryOrderReturn{}
		for _, deliveryOrderReturn := range deliveryOrderReturns {
			deliveryOrderReturnByDeliveryOrderId[deliveryOrderReturn.DeliveryOrderId] = deliveryOrderReturn
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderReturn *model.DeliveryOrderReturn
			if v, ok := deliveryOrderReturnByDeliveryOrderId[k.String()]; ok {
				deliveryOrderReturn = &v
			}

			result := &dataloader.Result{Data: deliveryOrderReturn, Error: nil}
			if deliveryOrderReturn == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderReturnLoader{
		loaderByDeliveryOrderId: NewDataloader(batchFn),
	}
}
