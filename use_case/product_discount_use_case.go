package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type productDiscountLoaderOptions struct {
}

type ProductDiscountUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.ProductDiscountCreateRequest) model.ProductDiscount

	// read
	Fetch(ctx context.Context, request dto_request.ProductDiscountFetchRequest) ([]model.ProductDiscount, int)
	Get(ctx context.Context, request dto_request.ProductDiscountGetRequest) model.ProductDiscount

	//  update
	Update(ctx context.Context, request dto_request.ProductDiscountUpdateRequest) model.ProductDiscount

	//  delete
	Delete(ctx context.Context, request dto_request.ProductDiscountDeleteRequest)
}

type productDiscountUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase
}

func NewProductDiscountUseCase(
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) ProductDiscountUseCase {
	return &productDiscountUseCase{
		repositoryManager: repositoryManager,
		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
	}
}

func (u *productDiscountUseCase) mustLoadProductDiscountDatas(ctx context.Context, productDiscounts []*model.ProductDiscount, options productDiscountLoaderOptions) {
	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range productDiscounts {
			group.Go(productLoader.ProductDiscountFn(productDiscounts[i]))
		}
	}))

	productUnitsLoader := loader.NewProductUnitsLoader(u.repositoryManager.ProductUnitRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productDiscounts {
				group.Go(productUnitsLoader.ProductFn(productDiscounts[i].Product))

				group.Go(fileLoader.ProductFn(productDiscounts[i].Product))
			}
		}),
	)

	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range productDiscounts {
			product := productDiscounts[i].Product
			for j := range product.ProductUnits {
				group.Go(unitLoader.ProductUnitFn(&product.ProductUnits[j]))
				group.Go(unitLoader.ProductUnitToUnitIdFn(&product.ProductUnits[j]))
			}
		}
	}))

	for i := range productDiscounts {
		product := productDiscounts[i].Product
		if product.ImageFile != nil {
			product.ImageFile.SetLink(u.baseFileUseCase.mainFilesystem)
		}
	}
}

func (u *productDiscountUseCase) mustValidateProductIdNotDuplicate(ctx context.Context, productId string) {
	isExist, err := u.repositoryManager.ProductDiscountRepository().IsExistByProductId(ctx, productId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_DISCOUNT.PRODUCT_ALREADY_HAVE_DISCOUNT"))
	}
}

func (u *productDiscountUseCase) Create(ctx context.Context, request dto_request.ProductDiscountCreateRequest) model.ProductDiscount {
	mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)
	u.mustValidateProductIdNotDuplicate(ctx, request.ProductId)

	productDiscount := model.ProductDiscount{
		Id:                 util.NewUuid(),
		ProductId:          request.ProductId,
		MinimumQty:         request.MinimumQty,
		IsActive:           request.IsActive,
		DiscountPercentage: request.DiscountPercentage,
		DiscountAmount:     request.DiscountAmount,
	}

	panicIfErr(
		u.repositoryManager.ProductDiscountRepository().Insert(ctx, &productDiscount),
	)

	u.mustLoadProductDiscountDatas(ctx, []*model.ProductDiscount{&productDiscount}, productDiscountLoaderOptions{})

	return productDiscount
}

func (u *productDiscountUseCase) Fetch(ctx context.Context, request dto_request.ProductDiscountFetchRequest) ([]model.ProductDiscount, int) {
	queryOption := model.ProductDiscountQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(request.Page, request.Limit, model.Sorts(request.Sorts)),
		Phrase:      request.Phrase,
		ProductId:   request.ProductId,
		IsActive:    request.IsActive,
	}

	productDiscounts, err := u.repositoryManager.ProductDiscountRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductDiscountRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductDiscountDatas(ctx, util.SliceValueToSlicePointer(productDiscounts), productDiscountLoaderOptions{})

	return productDiscounts, total
}

func (u *productDiscountUseCase) Get(ctx context.Context, request dto_request.ProductDiscountGetRequest) model.ProductDiscount {
	productDiscount := mustGetProductDiscount(ctx, u.repositoryManager, request.ProductDiscountId, false)

	u.mustLoadProductDiscountDatas(ctx, []*model.ProductDiscount{&productDiscount}, productDiscountLoaderOptions{})

	return productDiscount
}

func (u *productDiscountUseCase) Update(ctx context.Context, request dto_request.ProductDiscountUpdateRequest) model.ProductDiscount {
	productDiscount := mustGetProductDiscount(ctx, u.repositoryManager, request.ProductDiscountId, false)

	productDiscount.MinimumQty = request.MinimumQty
	productDiscount.IsActive = request.IsActive
	productDiscount.DiscountAmount = request.DiscountAmount
	productDiscount.DiscountPercentage = request.DiscountPercentage

	panicIfErr(
		u.repositoryManager.ProductDiscountRepository().Update(ctx, &productDiscount),
	)

	u.mustLoadProductDiscountDatas(ctx, []*model.ProductDiscount{&productDiscount}, productDiscountLoaderOptions{})

	return productDiscount
}

func (u *productDiscountUseCase) Delete(ctx context.Context, request dto_request.ProductDiscountDeleteRequest) {
	productDiscount := mustGetProductDiscount(ctx, u.repositoryManager, request.ProductDiscountId, true)

	panicIfErr(
		u.repositoryManager.ProductDiscountRepository().Delete(ctx, &productDiscount),
	)
}
