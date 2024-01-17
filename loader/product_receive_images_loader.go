package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveImagesLoader struct {
	loader dataloader.Loader
}

func (l *ProductReceiveImagesLoader) load(id string) ([]model.ProductReceiveImage, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReceiveImage), nil
}

func (l *ProductReceiveImagesLoader) ProductReceiveFn(productReceive *model.ProductReceive) func() error {
	return func() error {
		if productReceive != nil {
			productReceiveImages, err := l.load(productReceive.Id)
			if err != nil {
				return err
			}

			productReceive.ProductReceiveImages = productReceiveImages
		}

		return nil
	}
}

func NewProductReceiveImagesLoader(productReceiveImageRepository repository.ProductReceiveImageRepository) *ProductReceiveImagesLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceives, err := productReceiveImageRepository.FetchByProductReceiveIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveImagesByProductReceiveId := map[string][]model.ProductReceiveImage{}
		for _, productReceive := range productReceives {
			productReceiveImagesByProductReceiveId[productReceive.Id] = append(productReceiveImagesByProductReceiveId[productReceive.Id], productReceive)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceiveImages []model.ProductReceiveImage
			if v, ok := productReceiveImagesByProductReceiveId[k.String()]; ok {
				productReceiveImages = v
			}

			result := &dataloader.Result{Data: productReceiveImages, Error: nil}
			if productReceiveImages == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveImagesLoader{
		loader: NewDataloader(batchFn),
	}
}
