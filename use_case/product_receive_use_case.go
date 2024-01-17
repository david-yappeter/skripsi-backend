package use_case

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/infrastructure"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"

	"golang.org/x/sync/errgroup"
)

type productReceivesLoaderParams struct {
	productReceiveItems  bool
	productReceiveImages bool
}

type ProductReceiveUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.ProductReceiveCreateRequest) model.ProductReceive
	AddItem(ctx context.Context, request dto_request.ProductReceiveAddItemRequest) model.ProductReceive
	AddImage(ctx context.Context, request dto_request.ProductReceiveAddImageRequest) model.ProductReceive
	Upload(ctx context.Context, request dto_request.ProductReceiveUploadRequest) string

	// read
	Fetch(ctx context.Context, request dto_request.ProductReceiveFetchRequest) ([]model.ProductReceive, int)
	Get(ctx context.Context, request dto_request.ProductReceiveGetRequest) model.ProductReceive

	// update
	MarkCompleted(ctx context.Context, request dto_request.ProductReceiveMarkCompletedRequest) model.ProductReceive

	// delete
	Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest)
}

type productReceiveUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase baseFileUseCase

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func NewProductReceiveUseCase(
	repositoryManager repository.RepositoryManager,

	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) ProductReceiveUseCase {
	return &productReceiveUseCase{
		repositoryManager: repositoryManager,

		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

func (u *productReceiveUseCase) mustValidateAllowDeleteProductReceive(ctx context.Context, productReceiveId string) {

}

func (u *productReceiveUseCase) mustLoadProductReceivesData(ctx context.Context, productReceives []*model.ProductReceive, option productReceivesLoaderParams) {
	productReceiveItemsLoader := loader.NewProductReceiveItemsLoader(u.repositoryManager.ProductReceiveItemRepository())
	productReceiveImagesLoader := loader.NewProductReceiveImagesLoader(u.repositoryManager.ProductReceiveImageRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReceives {
				if option.productReceiveImages {
					group.Go(productReceiveImagesLoader.ProductReceiveFn(productReceives[i]))
				}

				if option.productReceiveItems {
					group.Go(productReceiveItemsLoader.ProductReceiveFn(productReceives[i]))
				}
			}
		}),
	)
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

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:  true,
		productReceiveImages: true,
	})

	return productReceive
}

func (u *productReceiveUseCase) AddImage(ctx context.Context, request dto_request.ProductReceiveAddImageRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, false)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	imageFile := model.File{
		Id:   util.NewUuid(),
		Type: data_type.FileTypeProductReceiveImage,
	}

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.ProductReceiveImagePath,
		productReceive.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.FilePath)),
		request.FilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: true,
		},
	)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:  true,
		productReceiveImages: true,
	})

	return productReceive
}

func (u *productReceiveUseCase) Upload(ctx context.Context, request dto_request.ProductReceiveUploadRequest) string {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, false)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.ProductReceiveImagePath,
		request.File.Filename,
		request.File,
		fileUploadTemporaryParams{
			supportedExtensions: listSupportedExtension([]string{
				extensionTypeImage,
			}),
			maxFileSizeInBytes: util.Pointer[int64](2 << 20),
		},
	)
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

func (u *productReceiveUseCase) MarkCompleted(ctx context.Context, request dto_request.ProductReceiveMarkCompletedRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	productReceive.Status = data_type.ProductReceiveStatusCompleted

	panicIfErr(
		u.repositoryManager.ProductReceiveRepository().Update(ctx, &productReceive),
	)

	return productReceive
}

func (u *productReceiveUseCase) Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest) {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	u.mustValidateAllowDeleteProductReceive(ctx, request.ProductReceiveId)

	panicIfErr(
		u.repositoryManager.ProductReceiveRepository().Delete(ctx, &productReceive),
	)
}
