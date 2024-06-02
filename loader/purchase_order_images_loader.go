package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type PurchaseOrderImagesLoader struct {
	loaderByPurchaseOrderId dataloader.Loader
}

func (l *PurchaseOrderImagesLoader) loadByPurchaseOrderId(id string) ([]model.PurchaseOrderImage, error) {
	thunk := l.loaderByPurchaseOrderId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.PurchaseOrderImage), nil
}

func (l *PurchaseOrderImagesLoader) PurchaseOrderFn(purchaseOrder *model.PurchaseOrder) func() error {
	return func() error {
		if purchaseOrder != nil {
			purchaseOrderImages, err := l.loadByPurchaseOrderId(purchaseOrder.Id)
			if err != nil {
				return err
			}

			purchaseOrder.PurchaseOrderImages = purchaseOrderImages
		}

		return nil
	}
}

func NewPurchaseOrderImagesLoader(purchaseOrderImageRepository repository.PurchaseOrderImageRepository) *PurchaseOrderImagesLoader {
	batchByPurchaseOrderIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		purchaseOrderImages, err := purchaseOrderImageRepository.FetchByPurchaseOrderIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		purchaseOrderImagesByPurchaseOrderId := map[string][]model.PurchaseOrderImage{}
		for _, purchaseOrderImage := range purchaseOrderImages {
			purchaseOrderImagesByPurchaseOrderId[purchaseOrderImage.PurchaseOrderId] = append(purchaseOrderImagesByPurchaseOrderId[purchaseOrderImage.PurchaseOrderId], purchaseOrderImage)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var purchaseOrderImages []model.PurchaseOrderImage
			if v, ok := purchaseOrderImagesByPurchaseOrderId[k.String()]; ok {
				purchaseOrderImages = v
			}

			result := &dataloader.Result{Data: purchaseOrderImages, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &PurchaseOrderImagesLoader{
		loaderByPurchaseOrderId: NewDataloader(batchByPurchaseOrderIdFn),
	}
}
