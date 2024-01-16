package use_case

import (
	"context"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/infrastructure"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type productReceivesLoaderParams struct {
	productReceiveItems bool
}

type ProductReceiveUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.ProductReceiveCreateRequest) model.ProductReceive
	AddItem(ctx context.Context, request dto_request.ProductReceiveAddItemRequest) model.ProductReceive
	Upload(ctx context.Context, request dto_request.ProductReceiveUploadRequest) model.ProductReceive

	// read
	Fetch(ctx context.Context, request dto_request.ProductReceiveFetchRequest) ([]model.ProductReceive, int)
	Get(ctx context.Context, request dto_request.ProductReceiveGetRequest) model.ProductReceive

	// delete
	Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest)
}

type productReceiveUseCase struct {
	repositoryManager repository.RepositoryManager
}

func (u *productReceiveUseCase) mustValidateAllowDeleteProductReceive(ctx context.Context, productReceiveId string) {

}

func (u *productReceiveUseCase) mustLoadProductReceivesData(ctx context.Context, productReceives []*model.ProductReceive, option productReceivesLoaderParams) {

}

func (u *productReceiveUseCase) Create(ctx context.Context, request dto_request.ProductReceiveCreateRequest) model.ProductReceive {
	var (
		authUser    = model.MustGetUserCtx(ctx)
		currentDate = util.CurrentDate()
	)

	productReceive := model.ProductReceive{
		Id:            util.NewUuid(),
		SupplierId:    request.SupplierId,
		UserId:        authUser.Id,
		InvoiceNumber: "",
		Date:          currentDate,
		Status:        data_type.ProductReceiveStatusPending,
		TotalPrice:    0,
	}

	panicIfErr(
		u.repositoryManager.ProductReceiveRepository().Insert(ctx, &productReceive),
	)

	return productReceive
}

func (u *productReceiveUseCase) AddItem(ctx context.Context, request dto_request.ProductReceiveAddItemRequest) model.ProductReceive {
	var (
		authUser       = model.MustGetUserCtx(ctx)
		productReceive = mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, false)
		productUnit    = mustGetProductUnitByProductIdAndUnitId(ctx, u.repositoryManager, request.ProductId, request.UnitId, true)
		product        = mustGetProduct(ctx, u.repositoryManager, request.ProductId, false)
	)

	if !product.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("Product data not found"))
	}

	// add total to product receive
	productReceive.TotalPrice += request.Qty * *product.Price

	// add product receive item
	productReceiveItem := model.ProductReceiveItem{
		Id:               util.NewUuid(),
		ProductReceiveId: productReceive.Id,
		ProductUnitId:    productUnit.Id,
		UserId:           authUser.Id,
		Qty:              request.Qty,
		PricePerUnit:     *product.Price,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error {
			productReceiveRepository := repository.NewProductReceiveRepository(tx, loggerStack)
			productReceiveItemRepository := repository.NewProductReceiveItemRepository(tx, loggerStack)

			if err := productReceiveRepository.Update(ctx, &productReceive); err != nil {
				return err
			}

			if err := productReceiveItemRepository.Insert(ctx, &productReceiveItem); err != nil {
				return err
			}

			return nil
		}),
	)

	return productReceive
}

func (u *productReceiveUseCase) Upload(ctx context.Context, request dto_request.ProductReceiveUploadRequest) model.ProductReceive {
	var (
		productReceive = mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, false)
	)

	return productReceive
}

func (u *productReceiveUseCase) Fetch(ctx context.Context, request dto_request.ProductReceiveFetchRequest) ([]model.ProductReceive, int) {
	queryOption := model.ProductReceiveQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	productReceives, err := u.repositoryManager.ProductReceiveRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductReceiveRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return productReceives, total
}

func (u *productReceiveUseCase) Get(ctx context.Context, request dto_request.ProductReceiveGetRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	return productReceive
}

func (u *productReceiveUseCase) Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest) {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	u.mustValidateAllowDeleteProductReceive(ctx, request.ProductReceiveId)

	panicIfErr(
		u.repositoryManager.ProductReceiveRepository().Delete(ctx, &productReceive),
	)
}

func NewProductReceiveUseCase(
	repositoryManager repository.RepositoryManager,
) ProductReceiveUseCase {
	return &productReceiveUseCase{
		repositoryManager: repositoryManager,
	}
}
