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
	// admin create
	AdminCreate(ctx context.Context, request dto_request.AdminProductUnitCreateRequest) model.ProductUnit
	AdminUpload(ctx context.Context, request dto_request.AdminProductUnitUploadRequest) string

	// admin read
	AdminGet(ctx context.Context, request dto_request.AdminProductUnitGetRequest) model.ProductUnit

	// admin update
	AdminUpdate(ctx context.Context, request dto_request.AdminProductUnitUpdateRequest) model.ProductUnit

	// admin delete
	AdminDelete(ctx context.Context, request dto_request.AdminProductUnitDeleteRequest)
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

func (u *productUnitUseCase) AdminCreate(ctx context.Context, request dto_request.AdminProductUnitCreateRequest) model.ProductUnit {
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

func (u *productUnitUseCase) AdminUpload(ctx context.Context, request dto_request.AdminProductUnitUploadRequest) string {
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

func (u *productUnitUseCase) AdminGet(ctx context.Context, request dto_request.AdminProductUnitGetRequest) model.ProductUnit {
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	return productUnit
}

func (u *productUnitUseCase) AdminUpdate(ctx context.Context, request dto_request.AdminProductUnitUpdateRequest) model.ProductUnit {
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	panicIfErr(
		u.repositoryManager.ProductUnitRepository().Update(ctx, &productUnit),
	)

	return productUnit
}

func (u *productUnitUseCase) AdminDelete(ctx context.Context, request dto_request.AdminProductUnitDeleteRequest) {
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
