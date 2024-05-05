package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderReturnImagesLoader struct {
	loaderByDeliveryOrderReturnId dataloader.Loader
}

func (l *DeliveryOrderReturnImagesLoader) loadByDeliveryOrderReturnId(id string) ([]model.DeliveryOrderReturnImage, error) {
	thunk := l.loaderByDeliveryOrderReturnId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.DeliveryOrderReturnImage), nil
}

func (l *DeliveryOrderReturnImagesLoader) DeliveryOrderReturnFn(deliveryOrderReturn *model.DeliveryOrderReturn) func() error {
	return func() error {
		if deliveryOrderReturn != nil {
			deliveryOrderReturnImages, err := l.loadByDeliveryOrderReturnId(deliveryOrderReturn.Id)
			if err != nil {
				return err
			}

			deliveryOrderReturn.DeliveryOrderReturnImages = deliveryOrderReturnImages
		}

		return nil
	}
}

func NewDeliveryOrderReturnImagesLoader(deliveryOrderReturnImageRepository repository.DeliveryOrderReturnImageRepository) *DeliveryOrderReturnImagesLoader {
	batchByDeliveryOrderReturnIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		deliveryOrderReturnImages, err := deliveryOrderReturnImageRepository.FetchByDeliveryOrderReturnIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderReturnImageByDeliveryOrderReturnId := map[string][]model.DeliveryOrderReturnImage{}
		for _, deliveryOrderReturnImage := range deliveryOrderReturnImages {
			deliveryOrderReturnImageByDeliveryOrderReturnId[deliveryOrderReturnImage.DeliveryOrderReturnId] = append(deliveryOrderReturnImageByDeliveryOrderReturnId[deliveryOrderReturnImage.DeliveryOrderReturnId], deliveryOrderReturnImage)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderReturnImages []model.DeliveryOrderReturnImage
			if v, ok := deliveryOrderReturnImageByDeliveryOrderReturnId[k.String()]; ok {
				deliveryOrderReturnImages = v
			}

			result := &dataloader.Result{Data: deliveryOrderReturnImages, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderReturnImagesLoader{
		loaderByDeliveryOrderReturnId: NewDataloader(batchByDeliveryOrderReturnIdFn),
	}
}
