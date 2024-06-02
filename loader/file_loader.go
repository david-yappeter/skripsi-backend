package loader

import (
	"context"
	"myapp/constant"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type FileLoader struct {
	loader dataloader.Loader
}

func (l *FileLoader) load(id string) (*model.File, error) {
	thunk := l.loader.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.(*model.File), nil
}

func (l *FileLoader) ProductReceiveImageFn(productReceiveImage *model.ProductReceiveImage) func() error {
	return func() error {
		file, err := l.load(productReceiveImage.FileId)
		if err != nil {
			return err
		}

		productReceiveImage.File = file

		return nil
	}
}

func (l *FileLoader) PurchaseOrderImageFn(purchaseOrderImage *model.PurchaseOrderImage) func() error {
	return func() error {
		file, err := l.load(purchaseOrderImage.FileId)
		if err != nil {
			return err
		}

		purchaseOrderImage.File = file

		return nil
	}
}

func (l *FileLoader) DeliveryOrderImageFn(deliveryOrderImage *model.DeliveryOrderImage) func() error {
	return func() error {
		file, err := l.load(deliveryOrderImage.FileId)
		if err != nil {
			return err
		}

		deliveryOrderImage.File = file

		return nil
	}
}

func (l *FileLoader) DeliveryOrderReturnImageFn(deliveryOrderReturnImage *model.DeliveryOrderReturnImage) func() error {
	return func() error {
		file, err := l.load(deliveryOrderReturnImage.FileId)
		if err != nil {
			return err
		}

		deliveryOrderReturnImage.File = file

		return nil
	}
}

func (l *FileLoader) ProductReceiveReturnImageFn(productReceiveReturnImage *model.ProductReceiveReturnImage) func() error {
	return func() error {
		file, err := l.load(productReceiveReturnImage.FileId)
		if err != nil {
			return err
		}

		productReceiveReturnImage.File = file

		return nil
	}
}

func (l *FileLoader) CustomerPaymentFn(customerPayment *model.CustomerPayment) func() error {
	return func() error {
		file, err := l.load(customerPayment.ImageFileId)
		if err != nil {
			return err
		}

		customerPayment.ImageFile = file

		return nil
	}
}

func (l *FileLoader) DebtPaymentFn(debtPayment *model.DebtPayment) func() error {
	return func() error {
		file, err := l.load(debtPayment.ImageFileId)
		if err != nil {
			return err
		}

		debtPayment.ImageFile = file

		return nil
	}
}

func (l *FileLoader) ProductFn(product *model.Product) func() error {
	return func() error {
		file, err := l.load(product.ImageFileId)
		if err != nil {
			return err
		}

		product.ImageFile = file

		return nil
	}
}

func NewFileLoader(fileRepository repository.FileRepository) *FileLoader {
	batchFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		files, err := fileRepository.FetchByIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		fileById := map[string]model.File{}
		for _, file := range files {
			fileById[file.Id] = file
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var file *model.File
			if v, ok := fileById[k.String()]; ok {
				file = &v
			}

			result := &dataloader.Result{Data: file, Error: nil}
			if file == nil {
				result.Error = constant.ErrNoData
			}
			results[idx] = result
		}

		return results
	}

	return &FileLoader{
		loader: NewDataloader(batchFn),
	}
}
