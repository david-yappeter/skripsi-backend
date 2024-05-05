package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReceiveReturnImagesLoader struct {
	loaderByProductReceiveReturnId dataloader.Loader
}

func (l *ProductReceiveReturnImagesLoader) loadByProductReceiveReturnId(id string) ([]model.ProductReceiveReturnImage, error) {
	thunk := l.loaderByProductReceiveReturnId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReceiveReturnImage), nil
}

func (l *ProductReceiveReturnImagesLoader) ProductReceiveReturnFn(productReceiveReturn *model.ProductReceiveReturn) func() error {
	return func() error {
		if productReceiveReturn != nil {
			productReceiveReturnImages, err := l.loadByProductReceiveReturnId(productReceiveReturn.Id)
			if err != nil {
				return err
			}

			productReceiveReturn.ProductReceiveReturnImages = productReceiveReturnImages
		}

		return nil
	}
}

func NewProductReceiveReturnImagesLoader(productReceiveReturnImageRepository repository.ProductReceiveReturnImageRepository) *ProductReceiveReturnImagesLoader {
	batchByProductReceiveReturnIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReceiveReturnImages, err := productReceiveReturnImageRepository.FetchByProductReceiveReturnIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReceiveReturnImageByProductReceiveReturnId := map[string][]model.ProductReceiveReturnImage{}
		for _, productReceiveReturnImage := range productReceiveReturnImages {
			productReceiveReturnImageByProductReceiveReturnId[productReceiveReturnImage.ProductReceiveReturnId] = append(productReceiveReturnImageByProductReceiveReturnId[productReceiveReturnImage.ProductReceiveReturnId], productReceiveReturnImage)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReceiveReturnImages []model.ProductReceiveReturnImage
			if v, ok := productReceiveReturnImageByProductReceiveReturnId[k.String()]; ok {
				productReceiveReturnImages = v
			}

			result := &dataloader.Result{Data: productReceiveReturnImages, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ProductReceiveReturnImagesLoader{
		loaderByProductReceiveReturnId: NewDataloader(batchByProductReceiveReturnIdFn),
	}
}
