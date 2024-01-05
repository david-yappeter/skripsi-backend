package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type ProductUseCase interface {
	// admin create
	Create(ctx context.Context, request dto_request.AdminProductCreateRequest) model.Product

	// admin read
	Fetch(ctx context.Context, request dto_request.AdminProductFetchRequest) ([]model.Product, int)
	Get(ctx context.Context, request dto_request.AdminProductGetRequest) model.Product

	// admin update
	Update(ctx context.Context, request dto_request.AdminProductUpdateRequest) model.Product

	// admin delete
	Delete(ctx context.Context, request dto_request.AdminProductDeleteRequest)
}

type productUseCase struct {
	repositoryManager repository.RepositoryManager
}

func (u *productUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.ProductRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("Product name already exist"))
	}
}

func (u *productUseCase) mustValidateAllowDeleteProduct(ctx context.Context, productId string) {

}

func (u *productUseCase) Create(ctx context.Context, request dto_request.AdminProductCreateRequest) model.Product {
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

func (u *productUseCase) Fetch(ctx context.Context, request dto_request.AdminProductFetchRequest) ([]model.Product, int) {
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

	return products, total
}

func (u *productUseCase) Get(ctx context.Context, request dto_request.AdminProductGetRequest) model.Product {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	return product
}

func (u *productUseCase) Update(ctx context.Context, request dto_request.AdminProductUpdateRequest) model.Product {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	if product.Name != request.Name {
		u.mustValidateNameNotDuplicate(ctx, request.Name)
	}

	if request.IsActive && request.Price == nil {
		panic(dto_response.NewBadRequestErrorResponse("Active product must have a price"))
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.IsActive = request.IsActive

	panicIfErr(
		u.repositoryManager.ProductRepository().Update(ctx, &product),
	)

	return product
}

func (u *productUseCase) Delete(ctx context.Context, request dto_request.AdminProductDeleteRequest) {
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
