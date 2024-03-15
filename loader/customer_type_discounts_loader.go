package loader

import (
	"context"
	"myapp/model"
	"myapp/repository"

	"github.com/graph-gophers/dataloader"
)

type CustomerTypeDiscountsLoader struct {
	loaderByCustomerTypeId dataloader.Loader
}

func (l *CustomerTypeDiscountsLoader) loadByCustomerTypeId(id string) ([]model.CustomerTypeDiscount, error) {
	thunk := l.loaderByCustomerTypeId.Load(context.TODO(), dataloader.StringKey(id))

	result, err := thunk()
	if err != nil {
		return nil, err
	}

	return result.([]model.CustomerTypeDiscount), nil
}

func (l *CustomerTypeDiscountsLoader) CustomerTypeFn(customerType *model.CustomerType) func() error {
	return func() error {
		if customerType != nil {
			customerTypeDiscounts, err := l.loadByCustomerTypeId(customerType.Id)
			if err != nil {
				return err
			}

			customerType.CustomerTypeDiscounts = customerTypeDiscounts
		}

		return nil
	}
}

func NewCustomerTypeDiscountsLoader(customerTypeDiscountRepository repository.CustomerTypeDiscountRepository) *CustomerTypeDiscountsLoader {
	batchByCustomerTypeIdFn := func(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
		ids := make([]string, len(keys))
		for idx, k := range keys {
			ids[idx] = k.String()
		}

		customerTypeDiscounts, err := customerTypeDiscountRepository.FetchByCustomerTypeIds(ctx, ids)
		if err != nil {
			panic(err)
		}

		customerTypeDiscountByCustomerTypeId := map[string][]model.CustomerTypeDiscount{}
		for _, customerTypeDiscount := range customerTypeDiscounts {
			customerTypeDiscountByCustomerTypeId[customerTypeDiscount.CustomerTypeId] = append(customerTypeDiscountByCustomerTypeId[customerTypeDiscount.CustomerTypeId], customerTypeDiscount)
		}

		results := make([]*dataloader.Result, len(keys))
		for idx, k := range keys {
			var userRoles []model.CustomerTypeDiscount
			if v, ok := customerTypeDiscountByCustomerTypeId[k.String()]; ok {
				userRoles = v
			}

			result := &dataloader.Result{Data: userRoles, Error: nil}
			results[idx] = result
		}

		return results
	}

	return &CustomerTypeDiscountsLoader{
		loaderByCustomerTypeId: NewDataloader(batchByCustomerTypeIdFn),
	}
}
