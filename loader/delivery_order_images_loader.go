package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderImagesLoader struct {
	loader dataloader.Loader
}

func (l *DeliveryOrderImagesLoader) load(id string) ([]model.DeliveryOrderImage, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.DeliveryOrderImage), nil
}

func (l *DeliveryOrderImagesLoader) DeliveryOrderFn(deliveryOrder *model.DeliveryOrder) func() error {
	return func() error {
		if deliveryOrder != nil {
			deliveryOrderImages, err := l.load(deliveryOrder.Id)
			if err != nil {
				return err
			}

			deliveryOrder.DeliveryOrderImages = deliveryOrderImages
		}

		return nil
	}
}

func NewDeliveryOrderImagesLoader(deliveryOrderImageRepository repository.DeliveryOrderImageRepository) *DeliveryOrderImagesLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		deliveryOrderImages, err := deliveryOrderImageRepository.FetchByDeliveryOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderImagesByProductReceiveId := map[string][]model.DeliveryOrderImage{}
		for _, deliveryOrder := range deliveryOrderImages {
			deliveryOrderImagesByProductReceiveId[deliveryOrder.Id] = append(deliveryOrderImagesByProductReceiveId[deliveryOrder.Id], deliveryOrder)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderImages []model.DeliveryOrderImage
			if v, ok := deliveryOrderImagesByProductReceiveId[k.String()]; ok {
				deliveryOrderImages = v
			}

			result := &dataloader.Result{Data: deliveryOrderImages, Error: nil}
			if deliveryOrderImages == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderImagesLoader{
		loader: NewDataloader(batchFn),
	}
}
