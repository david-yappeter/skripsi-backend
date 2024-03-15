package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type ProductDiscountUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.ProductDiscountCreateRequest) model.ProductDiscount

	//  update
	Update(ctx context.Context, request dto_request.ProductDiscountUpdateRequest) model.ProductDiscount

	//  delete
	Delete(ctx context.Context, request dto_request.ProductDiscountDeleteRequest)
}

type productDiscountUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewProductDiscountUseCase(
	repositoryManager repository.RepositoryManager,
) ProductDiscountUseCase {
	return &productDiscountUseCase{
		repositoryManager: repositoryManager,
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

	return productDiscount
}

func (u *productDiscountUseCase) Delete(ctx context.Context, request dto_request.ProductDiscountDeleteRequest) {
	productDiscount := mustGetProductDiscount(ctx, u.repositoryManager, request.ProductDiscountId, true)

	panicIfErr(
		u.repositoryManager.ProductDiscountRepository().Delete(ctx, &productDiscount),
	)
}
