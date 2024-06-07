package use_case

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"

	"golang.org/x/sync/errgroup"
)

type productReturnsLoaderParams struct {
	supplier            bool
	productReturnItems  bool
	productReturnImages bool

	productReturnProductStock bool
}

type ProductReturnUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.ProductReturnCreateRequest) model.ProductReturn
	AddItem(ctx context.Context, request dto_request.ProductReturnAddItemRequest) model.ProductReturn
	AddImage(ctx context.Context, request dto_request.ProductReturnAddImageRequest) model.ProductReturn
	Upload(ctx context.Context, request dto_request.ProductReturnUploadRequest) string

	// read
	Fetch(ctx context.Context, request dto_request.ProductReturnFetchRequest) ([]model.ProductReturn, int)
	Get(ctx context.Context, request dto_request.ProductReturnGetRequest) model.ProductReturn

	// update
	Update(ctx context.Context, request dto_request.ProductReturnUpdateRequest) model.ProductReturn
	MarkComplete(ctx context.Context, request dto_request.ProductReturnMarkCompleteRequest) model.ProductReturn

	// delete
	Delete(ctx context.Context, request dto_request.ProductReturnDeleteRequest)
	DeleteItem(ctx context.Context, request dto_request.ProductReturnDeleteItemRequest) model.ProductReturn
	DeleteImage(ctx context.Context, request dto_request.ProductReturnDeleteImageRequest) model.ProductReturn
}

type productReturnUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase baseFileUseCase

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func NewProductReturnUseCase(
	repositoryManager repository.RepositoryManager,

	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) ProductReturnUseCase {
	return &productReturnUseCase{
		repositoryManager: repositoryManager,

		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

func (u *productReturnUseCase) shouldGetProductReturnItemByProductReturnIdAndProductUnitId(ctx context.Context, productReturnId string, productUnitId string) *model.ProductReturnItem {
	productReturnItem, err := u.repositoryManager.ProductReturnItemRepository().GetByProductReturnIdAndProductUnitId(ctx, productReturnId, productUnitId)
	panicIfErr(err, constant.ErrNoData)

	return productReturnItem
}

func (u *productReturnUseCase) mustLoadProductReturnsData(ctx context.Context, productReturns []*model.ProductReturn, option productReturnsLoaderParams) {
	productReturnItemsLoader := loader.NewProductReturnItemsLoader(u.repositoryManager.ProductReturnItemRepository())
	productReturnImagesLoader := loader.NewProductReturnImagesLoader(u.repositoryManager.ProductReturnImageRepository())
	supplierLoader := loader.NewSupplierLoader(u.repositoryManager.SupplierRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReturns {
				if option.productReturnImages {
					group.Go(productReturnImagesLoader.ProductReturnFn(productReturns[i]))
				}

				if option.productReturnItems {
					group.Go(productReturnItemsLoader.ProductReturnFn(productReturns[i]))
				}

				if option.supplier {
					group.Go(supplierLoader.ProductReturnFn(productReturns[i]))
				}
			}
		}),
	)

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReturns {
				if option.productReturnImages {
					for j := range productReturns[i].ProductReturnImages {
						group.Go(fileLoader.ProductReturnImageFn(&productReturns[i].ProductReturnImages[j]))
					}
				}

				if option.productReturnItems {
					for j := range productReturns[i].ProductReturnItems {
						group.Go(productUnitLoader.ProductReturnItemFn(&productReturns[i].ProductReturnItems[j]))
					}
				}
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReturns {
				for j := range productReturns[i].ProductReturnItems {
					group.Go(productLoader.ProductUnitFn(productReturns[i].ProductReturnItems[j].ProductUnit))
					group.Go(unitLoader.ProductUnitFn(productReturns[i].ProductReturnItems[j].ProductUnit))
					if option.productReturnItems && option.productReturnProductStock {
						group.Go(productStockLoader.ProductUnitFn(productReturns[i].ProductReturnItems[j].ProductUnit))
					}
				}
			}
		}),
	)

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReturns {
				for j := range productReturns[i].ProductReturnItems {
					group.Go(fileLoader.ProductFn(productReturns[i].ProductReturnItems[j].ProductUnit.Product))
				}
			}
		}),
	)

	for i := range productReturns {
		for j := range productReturns[i].ProductReturnImages {
			productReturns[i].ProductReturnImages[j].File.SetLink(u.mainFilesystem)
		}

		for j := range productReturns[i].ProductReturnItems {
			productReturns[i].ProductReturnItems[j].ProductUnit.Product.ImageFile.SetLink(u.mainFilesystem)
		}
	}
}

func (u *productReturnUseCase) Create(ctx context.Context, request dto_request.ProductReturnCreateRequest) model.ProductReturn {
	var (
		authUser    = model.MustGetUserCtx(ctx)
		currentDate = util.CurrentDate()
	)

	mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	productReturn := model.ProductReturn{
		Id:            util.NewUuid(),
		SupplierId:    request.SupplierId,
		UserId:        authUser.Id,
		InvoiceNumber: request.InvoiceNumber,
		Date:          currentDate,
		Status:        data_type.ProductReturnStatusPending,
	}

	panicIfErr(
		u.repositoryManager.ProductReturnRepository().Insert(ctx, &productReturn),
	)

	return productReturn
}

func (u *productReturnUseCase) AddItem(ctx context.Context, request dto_request.ProductReturnAddItemRequest) model.ProductReturn {
	var (
		productReturn = mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, false)
		productUnit   = mustGetBaseProductUnitByProductId(ctx, u.repositoryManager, request.ProductId, true)
		product       = mustGetProduct(ctx, u.repositoryManager, request.ProductId, false)
		productStock  = shouldGetProductStockByProductId(ctx, u.repositoryManager, product.Id)

		totalSmallestQty = request.Qty * productUnit.ScaleToBase
	)

	if productReturn.Status != data_type.ProductReturnStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.STATUS.MUST_BE_PENDING"))
	}

	if !product.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.NOT_FOUND"))
	}

	if productStock.Qty < totalSmallestQty {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.INSUFFICIENT_STOCK"))
	}

	// deduct product stock
	productStock.Qty -= totalSmallestQty

	// add product return item
	productReturnItem := u.shouldGetProductReturnItemByProductReturnIdAndProductUnitId(ctx, productReturn.Id, productUnit.Id)
	isNewProductReturnItem := productReturnItem == nil

	if isNewProductReturnItem {
		productReturnItem = &model.ProductReturnItem{
			Id:              util.NewUuid(),
			ProductReturnId: productReturn.Id,
			ProductUnitId:   productUnit.Id,
			ScaleToBase:     productUnit.ScaleToBase,
			Qty:             request.Qty,
			BaseCostPrice:   productStock.BaseCostPrice,
		}
	} else {
		productReturnItem.Qty += request.Qty
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReturnRepository := u.repositoryManager.ProductReturnRepository()
			productReturnItemRepository := u.repositoryManager.ProductReturnItemRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()

			if err := productStockRepository.Update(ctx, productStock); err != nil {
				return err
			}

			if err := productReturnRepository.Update(ctx, &productReturn); err != nil {
				return err
			}

			if isNewProductReturnItem {
				if err := productReturnItemRepository.Insert(ctx, productReturnItem); err != nil {
					return err
				}
			} else {
				if err := productReturnItemRepository.Update(ctx, productReturnItem); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		productReturnItems:  true,
		productReturnImages: true,
		supplier:            true,
	})

	return productReturn
}

func (u *productReturnUseCase) AddImage(ctx context.Context, request dto_request.ProductReturnAddImageRequest) model.ProductReturn {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, false)

	if productReturn.Status != data_type.ProductReturnStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.STATUS.MUST_BE_PENDING"))
	}

	imageFile := model.File{
		Id:   util.NewUuid(),
		Type: data_type.FileTypeProductReturnImage,
	}

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.ProductReturnImagePath,
		productReturn.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.FilePath)),
		request.FilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: true,
		},
	)

	productReturnImage := model.ProductReturnImage{
		Id:              util.NewUuid(),
		ProductReturnId: productReturn.Id,
		FileId:          imageFile.Id,
		Description:     request.Description,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			productReturnImageRepository := u.repositoryManager.ProductReturnImageRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := productReturnImageRepository.Insert(ctx, &productReturnImage); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		productReturnItems:  true,
		productReturnImages: true,
		supplier:            true,
	})

	return productReturn
}

func (u *productReturnUseCase) Upload(ctx context.Context, request dto_request.ProductReturnUploadRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.ProductReturnImagePath,
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

func (u *productReturnUseCase) Fetch(ctx context.Context, request dto_request.ProductReturnFetchRequest) ([]model.ProductReturn, int) {
	queryOption := model.ProductReturnQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase:     request.Phrase,
		SupplierId: request.SupplierId,
		Status:     request.Status,
	}

	productReturns, err := u.repositoryManager.ProductReturnRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductReturnRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductReturnsData(ctx, util.SliceValueToSlicePointer(productReturns), productReturnsLoaderParams{
		supplier: true,
	})

	return productReturns, total
}

func (u *productReturnUseCase) Get(ctx context.Context, request dto_request.ProductReturnGetRequest) model.ProductReturn {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, true)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		productReturnItems:        true,
		productReturnProductStock: true,
		productReturnImages:       true,
		supplier:                  true,
	})

	return productReturn
}

func (u *productReturnUseCase) Update(ctx context.Context, request dto_request.ProductReturnUpdateRequest) model.ProductReturn {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, true)

	if productReturn.Status != data_type.ProductReturnStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.STATUS_MUST_BE_PENDING"))
	}

	productReturn.InvoiceNumber = request.InvoiceNumber
	productReturn.Date = request.Date

	panicIfErr(
		u.repositoryManager.ProductReturnRepository().Update(ctx, &productReturn),
	)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		productReturnItems:        true,
		productReturnProductStock: true,
		productReturnImages:       true,
		supplier:                  true,
	})

	return productReturn
}

func (u *productReturnUseCase) MarkComplete(ctx context.Context, request dto_request.ProductReturnMarkCompleteRequest) model.ProductReturn {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, true)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		productReturnItems:        true,
		productReturnImages:       true,
		productReturnProductStock: true,
		supplier:                  true,
	})

	if productReturn.Status != data_type.ProductReturnStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.STATUS.MUST_BE_PENDING"))
	}

	// change status
	productReturn.Status = data_type.ProductReturnStatusCompleted

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReturnRepository := u.repositoryManager.ProductReturnRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			for _, productReturnItem := range productReturn.ProductReturnItems {
				deductQtyLeft := productReturnItem.Qty

				for deductQtyLeft > 0 {
					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx, productReturnItem.ProductUnit.ProductId)
					if err != nil {
						return err
					}

					if deductQtyLeft > productStockMutation.BaseQtyLeft {
						deductQtyLeft -= productStockMutation.BaseQtyLeft
						productStockMutation.BaseQtyLeft = 0
					} else {
						productStockMutation.BaseQtyLeft -= deductQtyLeft
						deductQtyLeft = 0
					}

					if err := productStockMutationRepository.Update(ctx, productStockMutation); err != nil {
						return err
					}
				}
			}

			if err := productReturnRepository.Update(ctx, &productReturn); err != nil {
				return err
			}

			return nil
		}),
	)

	return productReturn
}

func (u *productReturnUseCase) Delete(ctx context.Context, request dto_request.ProductReturnDeleteRequest) {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReturnRepository := u.repositoryManager.ProductReturnRepository()
			productReturnItemRepository := u.repositoryManager.ProductReturnItemRepository()
			productReturnImageRepository := u.repositoryManager.ProductReturnImageRepository()

			if err := productReturnItemRepository.DeleteManyByProductReturnId(ctx, productReturn.Id); err != nil {
				return err
			}

			if err := productReturnImageRepository.DeleteManyByProductReturnId(ctx, productReturn.Id); err != nil {
				return err
			}

			if err := productReturnRepository.Delete(ctx, &productReturn); err != nil {
				return err
			}

			return nil
		}),
	)
}

func (u *productReturnUseCase) DeleteItem(ctx context.Context, request dto_request.ProductReturnDeleteItemRequest) model.ProductReturn {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, true)

	if productReturn.Status != data_type.ProductReturnStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.STATUS.MUST_BE_PENDING"))
	}

	// check delivery order item
	productReturnItem := mustGetProductReturnItem(ctx, u.repositoryManager, request.ProductReturnItemId, true)
	if productReturnItem.ProductReturnId != productReturn.Id {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN_ITEM.NOT_FOUND"))
	}

	// get product unit and product
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, productReturnItem.ProductUnitId, true)
	product := mustGetProduct(ctx, u.repositoryManager, productUnit.ProductId, true)

	// calculate qty
	totalSmallestQty := productReturnItem.Qty * productUnit.ScaleToBase

	productStock := shouldGetProductStockByProductId(ctx, u.repositoryManager, product.Id)

	// add product stock back
	productStock.Qty += totalSmallestQty

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReturnRepository := u.repositoryManager.ProductReturnRepository()
			productReturnItemRepository := u.repositoryManager.ProductReturnItemRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()

			if err := productReturnRepository.Update(ctx, &productReturn); err != nil {
				return err
			}

			if err := productReturnItemRepository.Delete(ctx, &productReturnItem); err != nil {
				return err
			}

			if err := productStockRepository.Update(ctx, productStock); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		supplier:            true,
		productReturnItems:  true,
		productReturnImages: true,
	})

	return productReturn
}

func (u *productReturnUseCase) DeleteImage(ctx context.Context, request dto_request.ProductReturnDeleteImageRequest) model.ProductReturn {
	productReturn := mustGetProductReturn(ctx, u.repositoryManager, request.ProductReturnId, true)

	if productReturn.Status != data_type.ProductReturnStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RETURN.STATUS.MUST_BE_PENDING"))
	}

	productReturnImage := mustGetProductReturnImage(ctx, u.repositoryManager, request.ProductReturnImageId, true)

	if productReturnImage.ProductReturnId != productReturn.Id {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE_ITEM.NOT_FOUND"))
	}

	file := mustGetFile(ctx, u.repositoryManager, productReturnImage.FileId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			productReturnImageRepository := u.repositoryManager.ProductReturnImageRepository()

			if err := productReturnImageRepository.Delete(ctx, &productReturnImage); err != nil {
				return err
			}

			if err := fileRepository.Delete(ctx, &file); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mainFilesystem.Delete(file.Path)

	u.mustLoadProductReturnsData(ctx, []*model.ProductReturn{&productReturn}, productReturnsLoaderParams{
		productReturnItems:  true,
		productReturnImages: true,
		supplier:            true,
	})

	return productReturn
}
