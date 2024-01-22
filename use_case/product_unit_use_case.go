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
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"
)

type ProductUnitUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.ProductUnitCreateRequest) model.ProductUnit
	Upload(ctx context.Context, request dto_request.ProductUnitUploadRequest) string

	//  read
	Get(ctx context.Context, request dto_request.ProductUnitGetRequest) model.ProductUnit

	//  update
	Update(ctx context.Context, request dto_request.ProductUnitUpdateRequest) model.ProductUnit

	//  delete
	Delete(ctx context.Context, request dto_request.ProductUnitDeleteRequest)

	// option
	OptionForProductReceiveForm(ctx context.Context, request dto_request.ProductUnitOptionForProductReceiveFormRequest) ([]model.ProductUnit, int)
	OptionForDeliveryOrderForm(ctx context.Context, request dto_request.ProductUnitOptionForDeliveryOrderFormRequest) ([]model.ProductUnit, int)
}

type productUnitUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase
}

func NewProductUnitUseCase(
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) ProductUnitUseCase {
	return &productUnitUseCase{
		repositoryManager: repositoryManager,
		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
	}
}

func (u *productUnitUseCase) mustValidateProductUnitNotDuplicate(ctx context.Context, productId string, unitId string) {
	isExist, err := u.repositoryManager.ProductUnitRepository().IsExistByProductIdAndUnitId(ctx, productId, unitId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_UNIT.ALREADY_EXIST"))
	}
}

func (u *productUnitUseCase) mustValidateAllowDeleteProductUnit(ctx context.Context, productUnitId string) {

}

func (u *productUnitUseCase) Create(ctx context.Context, request dto_request.ProductUnitCreateRequest) model.ProductUnit {
	var (
		imageFile *model.File
	)

	mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)
	mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	u.mustValidateProductUnitNotDuplicate(ctx, request.ProductId, request.UnitId)

	productUnit := model.ProductUnit{
		Id:          util.NewUuid(),
		ToUnitId:    request.ToUnitId,
		ImageFileId: nil,
		UnitId:      request.UnitId,
		ProductId:   request.ProductId,
		Scale:       request.Scale,
		ScaleToBase: request.Scale,
	}

	if request.ToUnitId != nil {
		toProductUnit := mustGetProductUnitByProductIdAndUnitId(ctx, u.repositoryManager, request.ProductId, *request.ToUnitId, true)

		productUnit.ScaleToBase = toProductUnit.ScaleToBase
	}

	if request.ImageFilePath != nil {
		imageFile = &model.File{
			Id:   util.NewUuid(),
			Type: data_type.FileTypeProductUnitImage,
		}

		productUnit.ImageFileId = &imageFile.Id

		imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
			ctx,
			constant.ProductUnitImagePath,
			productUnit.Id,
			fmt.Sprintf("%s%s", imageFile.Id, path.Ext(*request.ImageFilePath)),
			*request.ImageFilePath,
			fileUploadTemporaryToMainParams{
				deleteTmpOnSuccess: false,
			},
		)
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error {
			productUnitRepository := repository.NewProductUnitRepository(tx, loggerStack)
			fileRepository := repository.NewFileRepository(tx, loggerStack)

			if imageFile != nil {
				err := fileRepository.Insert(ctx, imageFile)
				if err != nil {
					return err
				}
			}

			err := productUnitRepository.Insert(ctx, &productUnit)
			if err != nil {
				return err
			}

			return nil
		}),
	)

	return productUnit
}

func (u *productUnitUseCase) Upload(ctx context.Context, request dto_request.ProductUnitUploadRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.ProductUnitImagePath,
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

func (u *productUnitUseCase) Get(ctx context.Context, request dto_request.ProductUnitGetRequest) model.ProductUnit {
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	return productUnit
}

func (u *productUnitUseCase) Update(ctx context.Context, request dto_request.ProductUnitUpdateRequest) model.ProductUnit {
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	productUnit.ProductId = request.ProductId
	productUnit.UnitId = request.UnitId

	panicIfErr(
		u.repositoryManager.ProductUnitRepository().Update(ctx, &productUnit),
	)

	return productUnit
}

func (u *productUnitUseCase) Delete(ctx context.Context, request dto_request.ProductUnitDeleteRequest) {
	var (
		file *model.File
	)

	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	if productUnit.ImageFileId != nil {
		file = util.Pointer(mustGetFile(ctx, u.repositoryManager, *productUnit.ImageFileId, true))
	}

	u.mustValidateAllowDeleteProductUnit(ctx, request.ProductUnitId)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error {
			fileRepository := repository.NewFileRepository(tx, loggerStack)
			productUnitRepository := repository.NewProductUnitRepository(tx, loggerStack)

			err := productUnitRepository.Delete(ctx, &productUnit)
			if err != nil {
				return err
			}

			if file != nil {
				err := fileRepository.Delete(ctx, file)
				if err != nil {
					return err
				}

				err = u.mainFilesystem.Delete(file.Path)
				if err != nil {
					return err
				}
			}

			return nil
		}),
	)
}

func (u *productUnitUseCase) OptionForProductReceiveForm(ctx context.Context, request dto_request.ProductUnitOptionForProductReceiveFormRequest) ([]model.ProductUnit, int) {
	mustGetProductReceiveItem(ctx, u.repositoryManager, request.ProductReceiveId, true)

	productReceiveItems, err := u.repositoryManager.ProductReceiveItemRepository().FetchByProductReceiveIds(ctx, []string{request.ProductReceiveId})
	panicIfErr(err)

	excludeProductUnitIds := []string{}
	for _, productReceiveItem := range productReceiveItems {
		excludeProductUnitIds = append(excludeProductUnitIds, productReceiveItem.ProductUnitId)
	}

	queryOption := model.ProductUnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		ExcludeIds: excludeProductUnitIds,
		Phrase:     request.Phrase,
	}

	productUnits, err := u.repositoryManager.ProductUnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductUnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return productUnits, total
}

func (u *productUnitUseCase) OptionForDeliveryOrderForm(ctx context.Context, request dto_request.ProductUnitOptionForDeliveryOrderFormRequest) ([]model.ProductUnit, int) {
	mustGetDeliveryOrderItem(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	deliveryOrderItems, err := u.repositoryManager.DeliveryOrderItemRepository().FetchByDeliveryOrderIds(ctx, []string{request.DeliveryOrderId})
	panicIfErr(err)

	excludeProductUnitIds := []string{}
	for _, deliveryOrderItem := range deliveryOrderItems {
		excludeProductUnitIds = append(excludeProductUnitIds, deliveryOrderItem.ProductUnitId)
	}

	queryOption := model.ProductUnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		ExcludeIds: excludeProductUnitIds,
		Phrase:     request.Phrase,
	}

	productUnits, err := u.repositoryManager.ProductUnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductUnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return productUnits, total
}
