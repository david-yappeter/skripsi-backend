package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type productLoaderParams struct {
	productStock bool
}

type ProductUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.ProductCreateRequest) model.Product

	//  read
	Fetch(ctx context.Context, request dto_request.ProductFetchRequest) ([]model.Product, int)
	Get(ctx context.Context, request dto_request.ProductGetRequest) model.Product

	//  update
	Update(ctx context.Context, request dto_request.ProductUpdateRequest) model.Product

	//  delete
	Delete(ctx context.Context, request dto_request.ProductDeleteRequest)
}

type productUseCase struct {
	repositoryManager repository.RepositoryManager
}

func (u *productUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.ProductRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.NAME.ALREADY_EXIST"))
	}
}

func (u *productUseCase) mustValidateAllowDeleteProduct(ctx context.Context, productId string) {
	isExist, err := u.repositoryManager.ProductStockRepository().IsExistByProductId(ctx, productId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.ALREADY_HAVE_STOCK"))
	}
}

func (u *productUseCase) mustLoadProductDatas(ctx context.Context, products []*model.Product, option productLoaderParams) {
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range products {
				group.Go(productStockLoader.ProductFn(products[i]))
			}
		}),
	)
}

func (u *productUseCase) Create(ctx context.Context, request dto_request.ProductCreateRequest) model.Product {
	u.mustValidateNameNotDuplicate(ctx, request.Name)

	product := model.Product{
		Id:          util.NewUuid(),
		Name:        request.Name,
		Description: request.Description,
		Price:       nil,
		IsActive:    false,
	}

	panicIfErr(
		u.repositoryManager.ProductRepository().Insert(ctx, &product),
	)

	return product
}

func (u *productUseCase) Fetch(ctx context.Context, request dto_request.ProductFetchRequest) ([]model.Product, int) {
	queryOption := model.ProductQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	products, err := u.repositoryManager.ProductRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductDatas(ctx, util.SliceValueToSlicePointer(products), productLoaderParams{
		productStock: true,
	})

	return products, total
}

func (u *productUseCase) Get(ctx context.Context, request dto_request.ProductGetRequest) model.Product {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	u.mustLoadProductDatas(ctx, []*model.Product{&product}, productLoaderParams{
		productStock: true,
	})

	return product
}

func (u *productUseCase) Update(ctx context.Context, request dto_request.ProductUpdateRequest) model.Product {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	if product.Name != request.Name {
		u.mustValidateNameNotDuplicate(ctx, request.Name)
	}

	if request.IsActive && request.Price == nil {
		panic(dto_response.NewBadRequestErrorResponse("ACTIVE_PRODUCT.MUST_HAVE_PRICE"))
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.IsActive = request.IsActive

	panicIfErr(
		u.repositoryManager.ProductRepository().Update(ctx, &product),
	)

	u.mustLoadProductDatas(ctx, []*model.Product{&product}, productLoaderParams{
		productStock: true,
	})

	return product
}

func (u *productUseCase) Delete(ctx context.Context, request dto_request.ProductDeleteRequest) {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	u.mustValidateAllowDeleteProduct(ctx, request.ProductId)

	panicIfErr(
		u.repositoryManager.ProductRepository().Delete(ctx, &product),
	)
}

func NewProductUseCase(
	repositoryManager repository.RepositoryManager,
) ProductUseCase {
	return &productUseCase{
		repositoryManager: repositoryManager,
	}
}
