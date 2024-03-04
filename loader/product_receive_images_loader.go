package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveImagesLoader struct {
	loaderByProductReceiveId dataloader.Loader
}

func (l *ProductReceiveImagesLoader) loadByProductReceiveId(id string) ([]model.ProductReceiveImage, error) {
	thunk := l.loaderByProductReceiveId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReceiveImage), nil
}

func (l *ProductReceiveImagesLoader) ProductReceiveFn(productReceive *model.ProductReceive) func() error {
	return func() error {
		if productReceive != nil {
			productReceiveImages, err := l.loadByProductReceiveId(productReceive.Id)
			if err != nil {
				return err
			}

			productReceive.ProductReceiveImages = productReceiveImages
		}

		return nil
	}
}

func NewProductReceiveImagesLoader(productReceiveImageRepository repository.ProductReceiveImageRepository) *ProductReceiveImagesLoader {
	batchByProductReceiveIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceiveImages, err := productReceiveImageRepository.FetchByProductReceiveIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveImagesByProductReceiveId := map[string][]model.ProductReceiveImage{}
		for _, productReceiveImage := range productReceiveImages {
			productReceiveImagesByProductReceiveId[productReceiveImage.ProductReceiveId] = append(productReceiveImagesByProductReceiveId[productReceiveImage.ProductReceiveId], productReceiveImage)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceiveImages []model.ProductReceiveImage
			if v, ok := productReceiveImagesByProductReceiveId[k.String()]; ok {
				productReceiveImages = v
			}

			result := &dataloader.Result{Data: productReceiveImages, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveImagesLoader{
		loaderByProductReceiveId: NewDataloader(batchByProductReceiveIdFn),
	}
}
