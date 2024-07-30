package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type customerTypeDiscountLoaderParams struct {
	product bool
}

type CustomerTypeDiscountUseCase interface {
	// option
	OptionForWhatsappCustomerTypeDiscountChangeBroadcastForm(ctx context.Context, request dto_request.CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest) ([]model.CustomerTypeDiscount, int)
}

type customerTypeDiscountUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewCustomerTypeDiscountUseCase(
	repositoryManager repository.RepositoryManager,
) CustomerTypeDiscountUseCase {
	return &customerTypeDiscountUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *customerTypeDiscountUseCase) mustLoadCustomerTypeDiscountsData(ctx context.Context, customerTypeDiscounts []*model.CustomerTypeDiscount, option customerTypeDiscountLoaderParams) {
	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			if option.product {
				for i := range customerTypeDiscounts {
					group.Go(productLoader.CustomerTypeDiscountFn(customerTypeDiscounts[i]))
				}
			}
		}),
	)
}
func (u *customerTypeDiscountUseCase) OptionForWhatsappCustomerTypeDiscountChangeBroadcastForm(ctx context.Context, request dto_request.CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest) ([]model.CustomerTypeDiscount, int) {
	queryOption := model.CustomerTypeDiscountQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	customerTypeDiscounts, err := u.repositoryManager.CustomerTypeDiscountRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CustomerTypeDiscountRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadCustomerTypeDiscountsData(ctx, util.SliceValueToSlicePointer(customerTypeDiscounts), customerTypeDiscountLoaderParams{
		product: true,
	})

	return customerTypeDiscounts, total
}
