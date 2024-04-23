package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type DeliveryOrderReviewLoader struct {
	loaderByDeliveryOrderId dataloader.Loader
}

func (l *DeliveryOrderReviewLoader) loadByDeliveryOrderId(id string) (*model.DeliveryOrderReview, error) {
	thunk := l.loaderByDeliveryOrderId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.DeliveryOrderReview), nil
}

func (l *DeliveryOrderReviewLoader) DeliveryOrderFn(deliveryOrder *model.DeliveryOrder) func() error {
	return func() error {
		deliveryOrderReview, err := l.loadByDeliveryOrderId(deliveryOrder.Id)
		if err != nil {
			return err
		}

		deliveryOrder.DeliveryOrderReview = deliveryOrderReview

		return nil
	}
}

func NewDeliveryOrderReviewLoader(deliveryOrderReviewRepository repository.DeliveryOrderReviewRepository) *DeliveryOrderReviewLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		deliveryOrderReviews, err := deliveryOrderReviewRepository.FetchByDeliveryOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		deliveryOrderReviewByDeliveryOrderId := map[string]model.DeliveryOrderReview{}
		for _, deliveryOrderReview := range deliveryOrderReviews {
			deliveryOrderReviewByDeliveryOrderId[deliveryOrderReview.DeliveryOrderId] = deliveryOrderReview
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var deliveryOrderReview *model.DeliveryOrderReview
			if v, ok := deliveryOrderReviewByDeliveryOrderId[k.String()]; ok {
				deliveryOrderReview = &v
			}

			result := &dataloader.Result{Data: deliveryOrderReview, Error: nil}
			if deliveryOrderReview == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &DeliveryOrderReviewLoader{
		loaderByDeliveryOrderId: NewDataloader(batchFn),
	}
}
