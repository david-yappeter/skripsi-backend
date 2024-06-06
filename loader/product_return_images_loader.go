package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type ProductReturnImagesLoader struct {
	loaderByProductReturnId dataloader.Loader
}

func (l *ProductReturnImagesLoader) loadByProductReturnId(id string) ([]model.ProductReturnImage, error) {
	thunk := l.loaderByProductReturnId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.ProductReturnImage), nil
}

func (l *ProductReturnImagesLoader) ProductReturnFn(productReturn *model.ProductReturn) func() error {
	return func() error {
		if productReturn != nil {
			productReturnImages, err := l.loadByProductReturnId(productReturn.Id)
			if err != nil {
				return err
			}

			productReturn.ProductReturnImages = productReturnImages
		}

		return nil
	}
}

func NewProductReturnImagesLoader(productReturnImageRepository repository.ProductReturnImageRepository) *ProductReturnImagesLoader {
	batchByProductReturnIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		productReturnImages, err := productReturnImageRepository.FetchByProductReturnIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		productReturnImagesByProductReturnId := map[string][]model.ProductReturnImage{}
		for _, productReturnImage := range productReturnImages {
			productReturnImagesByProductReturnId[productReturnImage.ProductReturnId] = append(productReturnImagesByProductReturnId[productReturnImage.ProductReturnId], productReturnImage)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var productReturnImages []model.ProductReturnImage
			if v, ok := productReturnImagesByProductReturnId[k.String()]; ok {
				productReturnImages = v
			}

			result := &dataloader.Result{Data: productReturnImages, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &ProductReturnImagesLoader{
		loaderByProductReturnId: NewDataloader(batchByProductReturnIdFn),
	}
}
