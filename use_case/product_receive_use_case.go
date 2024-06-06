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

type productReceivesLoaderParams struct {
	supplier             bool
	productReceiveItems  bool
	productReceiveImages bool

	productReceiveProductStockMutation bool
	productReceiveProductStock         bool

	_return bool
}

type ProductReceiveUseCase interface {
	// create
	AddImage(ctx context.Context, request dto_request.ProductReceiveAddImageRequest) model.ProductReceive
	Upload(ctx context.Context, request dto_request.ProductReceiveUploadRequest) string

	// read
	Fetch(ctx context.Context, request dto_request.ProductReceiveFetchRequest) ([]model.ProductReceive, int)
	Get(ctx context.Context, request dto_request.ProductReceiveGetRequest) model.ProductReceive

	// update
	Update(ctx context.Context, request dto_request.ProductReceiveUpdateRequest) model.ProductReceive
	Cancel(ctx context.Context, request dto_request.ProductReceiveCancelRequest) model.ProductReceive
	MarkComplete(ctx context.Context, request dto_request.ProductReceiveMarkCompleteRequest) model.ProductReceive
	UpdateItem(ctx context.Context, request dto_request.ProductReceiveUpdateItemRequest) model.ProductReceive

	// delete
	Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest)
	DeleteImage(ctx context.Context, request dto_request.ProductReceiveDeleteImageRequest) model.ProductReceive
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

func (u *productReceiveUseCase) mustLoadProductReceivesData(ctx context.Context, productReceives []*model.ProductReceive, option productReceivesLoaderParams) {
	productReceiveItemsLoader := loader.NewProductReceiveItemsLoader(u.repositoryManager.ProductReceiveItemRepository())
	productReceiveImagesLoader := loader.NewProductReceiveImagesLoader(u.repositoryManager.ProductReceiveImageRepository())
	supplierLoader := loader.NewSupplierLoader(u.repositoryManager.SupplierRepository())
	productReceiveReturnLoader := loader.NewProductReceiveReturnLoader(u.repositoryManager.ProductReceiveReturnRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReceives {
				if option.productReceiveImages {
					group.Go(productReceiveImagesLoader.ProductReceiveFn(productReceives[i]))
				}

				if option.productReceiveItems {
					group.Go(productReceiveItemsLoader.ProductReceiveFn(productReceives[i]))
				}

				if option.supplier {
					group.Go(supplierLoader.ProductReceiveFn(productReceives[i]))
				}

				if option._return {
					group.Go(productReceiveReturnLoader.ProductReceiveFnNotStrict(productReceives[i]))
				}
			}
		}),
	)

	productStockMutationLoader := loader.NewProductStockMutationLoader(u.repositoryManager.ProductStockMutationRepository())
	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())
	productReceiveReturnImagesLoader := loader.NewProductReceiveReturnImagesLoader(u.repositoryManager.ProductReceiveReturnImageRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReceives {
				if option.productReceiveImages {
					for j := range productReceives[i].ProductReceiveImages {
						group.Go(fileLoader.ProductReceiveImageFn(&productReceives[i].ProductReceiveImages[j]))
					}
				}

				if option.productReceiveItems {
					for j := range productReceives[i].ProductReceiveItems {
						group.Go(productUnitLoader.ProductReceiveItemFn(&productReceives[i].ProductReceiveItems[j]))
						if option.productReceiveProductStockMutation {
							group.Go(productStockMutationLoader.ProductReceiveItemNotStrictFn(&productReceives[i].ProductReceiveItems[j]))
						}
					}
				}

				if option._return && productReceives[i].ProductReceiveReturn != nil {
					group.Go(productReceiveReturnImagesLoader.ProductReceiveReturnFn(productReceives[i].ProductReceiveReturn))
				}
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReceives {
				for j := range productReceives[i].ProductReceiveItems {
					group.Go(productLoader.ProductUnitFn(productReceives[i].ProductReceiveItems[j].ProductUnit))
					group.Go(unitLoader.ProductUnitFn(productReceives[i].ProductReceiveItems[j].ProductUnit))
					if option.productReceiveItems && option.productReceiveProductStock {
						group.Go(productStockLoader.ProductUnitFn(productReceives[i].ProductReceiveItems[j].ProductUnit))
					}
				}

				if option._return && productReceives[i].ProductReceiveReturn != nil {
					for j := range productReceives[i].ProductReceiveReturn.ProductReceiveReturnImages {
						group.Go(fileLoader.ProductReceiveReturnImageFn(&productReceives[i].ProductReceiveReturn.ProductReceiveReturnImages[j]))
					}
				}
			}
		}),
	)

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReceives {
				for j := range productReceives[i].ProductReceiveItems {
					group.Go(fileLoader.ProductFn(productReceives[i].ProductReceiveItems[j].ProductUnit.Product))
				}
			}
		}),
	)

	for i := range productReceives {
		for j := range productReceives[i].ProductReceiveImages {
			productReceives[i].ProductReceiveImages[j].File.SetLink(u.mainFilesystem)
		}

		for j := range productReceives[i].ProductReceiveItems {
			productReceives[i].ProductReceiveItems[j].ProductUnit.Product.ImageFile.SetLink(u.mainFilesystem)
		}

		if productReceives[i].ProductReceiveReturn != nil {
			for j := range productReceives[i].ProductReceiveReturn.ProductReceiveReturnImages {
				productReceives[i].ProductReceiveReturn.ProductReceiveReturnImages[j].File.SetLink(u.mainFilesystem)
			}
		}
	}
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

	productReceiveImage := model.ProductReceiveImage{
		Id:               util.NewUuid(),
		ProductReceiveId: productReceive.Id,
		FileId:           imageFile.Id,
		Description:      request.Description,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			productReceiveImageRepository := u.repositoryManager.ProductReceiveImageRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := productReceiveImageRepository.Insert(ctx, &productReceiveImage); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:  true,
		productReceiveImages: true,
		supplier:             true,
	})

	return productReceive
}

func (u *productReceiveUseCase) Upload(ctx context.Context, request dto_request.ProductReceiveUploadRequest) string {
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
		Phrase:     request.Phrase,
		SupplierId: request.SupplierId,
		Status:     request.Status,
	}

	productReceives, err := u.repositoryManager.ProductReceiveRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductReceiveRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductReceivesData(ctx, util.SliceValueToSlicePointer(productReceives), productReceivesLoaderParams{
		supplier: true,
	})

	return productReceives, total
}

func (u *productReceiveUseCase) Get(ctx context.Context, request dto_request.ProductReceiveGetRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:        true,
		productReceiveProductStock: true,
		productReceiveImages:       true,
		supplier:                   true,
		_return:                    true,
	})

	return productReceive
}

func (u *productReceiveUseCase) Update(ctx context.Context, request dto_request.ProductReceiveUpdateRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS_MUST_BE_PENDING"))
	}

	productReceive.InvoiceNumber = request.InvoiceNumber
	productReceive.Date = request.Date

	panicIfErr(
		u.repositoryManager.ProductReceiveRepository().Update(ctx, &productReceive),
	)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:        true,
		productReceiveProductStock: true,
		productReceiveImages:       true,
		supplier:                   true,
	})

	return productReceive
}

func (u *productReceiveUseCase) Cancel(ctx context.Context, request dto_request.ProductReceiveCancelRequest) model.ProductReceive {
	var (
		toBeRemovedStockByProductId        map[string]float64            = nil
		toBeDeletedProductStockMutationIds                               = []string{}
		productStockByProductId            map[string]model.ProductStock = nil
	)

	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status == data_type.ProductReceiveStatusCanceled {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.ALREADY_CANCELED"))
	}

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:                true,
		productReceiveProductStock:         true,
		productReceiveProductStockMutation: true,
	})

	switch productReceive.Status {
	case data_type.ProductReceiveStatusCompleted:
		u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
			productReceiveProductStock: true,
		})

		if len(productReceive.ProductReceiveItems) > 0 {
			toBeRemovedStockByProductId = make(map[string]float64)
			productStockByProductId = make(map[string]model.ProductStock)
		}

		for _, productReceiveItem := range productReceive.ProductReceiveItems {
			productStockByProductId[productReceiveItem.ProductUnit.ProductId] = *productReceiveItem.ProductUnit.ProductStock
			toBeRemovedStockByProductId[productReceiveItem.ProductUnit.ProductId] += productReceiveItem.QtyReceived * productReceiveItem.ScaleToBase

			if productReceiveItem.ProductStockMutation != nil {
				productStockMutation := productReceiveItem.ProductStockMutation

				if productStockMutation.BaseQtyLeft != productStockMutation.Qty {
					panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.PRODUCT_ALREADY_SOLD_CANNOT_BE_CANCELED"))
				}

				toBeDeletedProductStockMutationIds = append(toBeDeletedProductStockMutationIds, productStockMutation.Id)
			}
		}

		// remove stock and remove from tiktok
		for productId, removedStock := range toBeRemovedStockByProductId {
			productStock := productStockByProductId[productId]
			productStock.Qty -= removedStock
			productStockByProductId[productId] = productStock

			tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, productId)

			if tiktokProduct != nil {
				mustUpdateTiktokProductInventory(ctx, u.repositoryManager, tiktokProduct.TiktokProductId, int(productStock.Qty))
			}
		}
	}

	productReceive.Status = data_type.ProductReceiveStatusCanceled

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReceiveRepository := u.repositoryManager.ProductReceiveRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			if err := productReceiveRepository.Update(ctx, &productReceive); err != nil {
				return err
			}

			for _, productStock := range productStockByProductId {
				if err := productStockRepository.Update(ctx, &productStock); err != nil {
					return err
				}
			}

			if len(toBeDeletedProductStockMutationIds) > 0 {
				if err := productStockMutationRepository.DeleteManyByIds(ctx, toBeDeletedProductStockMutationIds); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveImages: true,
		supplier:             true,
	})

	return productReceive
}

func (u *productReceiveUseCase) MarkComplete(ctx context.Context, request dto_request.ProductReceiveMarkCompleteRequest) model.ProductReceive {
	var (
		currentDateTime                                        = util.CurrentDateTime()
		productStockMutations                                  = []model.ProductStockMutation{}
		productStockByProductId map[string]*model.ProductStock = make(map[string]*model.ProductStock)
	)

	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:        true,
		productReceiveImages:       true,
		productReceiveProductStock: true,
		supplier:                   true,
	})

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	for _, productReceiveItem := range productReceive.ProductReceiveItems {
		// adjust base cost price
		productStockByProductId[productReceiveItem.ProductUnit.ProductId] = new(model.ProductStock)
		*productStockByProductId[productReceiveItem.ProductUnit.ProductId] = *productReceiveItem.ProductUnit.ProductStock

		productStockByProductId[productReceiveItem.ProductUnit.ProductId].BaseCostPrice = productStockByProductId[productReceiveItem.ProductUnit.ProductId].RecalculateBaseCostPrice(productReceiveItem.BaseEligibleQty(), productReceiveItem.PricePerUnit/productReceiveItem.ScaleToBase)
		productStockByProductId[productReceiveItem.ProductUnit.ProductId].Qty += productReceiveItem.BaseEligibleQty()

		productStockMutations = append(productStockMutations, model.ProductStockMutation{
			Id:            util.NewUuid(),
			ProductUnitId: productReceiveItem.ProductUnitId,
			Type:          data_type.ProductStockMutationTypeProductReceiveItem,
			IdentifierId:  productReceiveItem.Id,
			Qty:           productReceiveItem.QtyEligible,
			ScaleToBase:   productReceiveItem.ScaleToBase,
			BaseQtyLeft:   productReceiveItem.BaseEligibleQty(),
			BaseCostPrice: productReceiveItem.PricePerUnit / productReceiveItem.ScaleToBase,
			MutatedAt:     currentDateTime,
		})
	}

	// add stock and sync tiktok stock
	for productId, productStock := range productStockByProductId {
		tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, productId)

		if tiktokProduct != nil {
			mustUpdateTiktokProductInventory(ctx, u.repositoryManager, tiktokProduct.TiktokProductId, int(productStock.Qty))
		}

	}

	// change status
	productReceive.Status = data_type.ProductReceiveStatusCompleted

	// initialize debt
	debt := model.Debt{
		Id:                   util.NewUuid(),
		DebtSource:           data_type.DebtSourceProductReceive,
		DebtSourceIdentifier: productReceive.Id,
		DueDate:              data_type.NewNullDate(nil),
		Status:               data_type.DebtStatusUnpaid,
		Amount:               productReceive.TotalPrice,
		RemainingAmount:      productReceive.TotalPrice,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			debtRepository := u.repositoryManager.DebtRepository()
			productReceiveRepository := u.repositoryManager.ProductReceiveRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			if err := debtRepository.Insert(ctx, &debt); err != nil {
				return err
			}

			if err := productReceiveRepository.Update(ctx, &productReceive); err != nil {
				return err
			}

			for _, productStock := range productStockByProductId {
				if err := productStockRepository.Update(ctx, productStock); err != nil {
					return err
				}
			}

			if err := productStockMutationRepository.InsertMany(ctx, productStockMutations); err != nil {
				return err
			}

			return nil
		}),
	)

	return productReceive
}

func (u *productReceiveUseCase) UpdateItem(ctx context.Context, request dto_request.ProductReceiveUpdateItemRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS_MUST_BE_PENDING"))
	}

	productReceiveItem := mustGetProductReceiveItem(ctx, u.repositoryManager, request.ProductReceiveItemId, true)

	if request.QtyEligible > productReceiveItem.QtyReceived {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE_ITEM.INVALID_AMOUNT_QTY_MUST_BE_SMALLER_THAN_OR_EQUAL_RECEIVED_QTY"))
	}

	productReceiveItem.QtyEligible = request.QtyEligible

	panicIfErr(
		u.repositoryManager.ProductReceiveItemRepository().Update(ctx, &productReceiveItem),
	)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:        true,
		productReceiveProductStock: true,
		productReceiveImages:       true,
		supplier:                   true,
	})

	return productReceive
}

func (u *productReceiveUseCase) Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest) {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReceiveRepository := u.repositoryManager.ProductReceiveRepository()
			productReceiveItemRepository := u.repositoryManager.ProductReceiveItemRepository()
			productReceiveImageRepository := u.repositoryManager.ProductReceiveImageRepository()

			if err := productReceiveItemRepository.DeleteManyByProductReceiveId(ctx, productReceive.Id); err != nil {
				return err
			}

			if err := productReceiveImageRepository.DeleteManyByProductReceiveId(ctx, productReceive.Id); err != nil {
				return err
			}

			if err := productReceiveRepository.Delete(ctx, &productReceive); err != nil {
				return err
			}

			return nil
		}),
	)
}

// func (u *productReceiveUseCase) Returned(ctx context.Context, request dto_request.ProductReceiveReturnedRequest) model.ProductReceive {
// 	currentUser := model.MustGetUserCtx(ctx)
// 	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

// 	if productReceive.Status != data_type.ProductReceiveStatusCompleted {
// 		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_COMPLETED"))
// 	}

// 	productReceive.Status = data_type.ProductReceiveStatusReturned

// 	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
// 		productReceiveItems: true,
// 	})

// 	// initialize return
// 	productReceiveReturn := model.ProductReceiveReturn{
// 		Id:               util.NewUuid(),
// 		ProductReceiveId: productReceive.Id,
// 		UserId:           currentUser.Id,
// 		Description:      request.Description,
// 	}

// 	// initialize images
// 	files := []model.File{}
// 	productReceiveReturnImages := []model.ProductReceiveReturnImage{}

// 	for _, filepath := range request.FilePaths {
// 		imageFile := model.File{
// 			Id:   util.NewUuid(),
// 			Type: data_type.FileTypeProductReceiveReturnImage,
// 		}

// 		imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
// 			ctx,
// 			constant.ProductReceiveReturnPath,
// 			productReceiveReturn.Id,
// 			fmt.Sprintf("%s%s", imageFile.Id, path.Ext(filepath)),
// 			filepath,
// 			fileUploadTemporaryToMainParams{
// 				deleteTmpOnSuccess: false,
// 			},
// 		)

// 		productReceiveReturnImages = append(productReceiveReturnImages, model.ProductReceiveReturnImage{
// 			Id:                     util.NewUuid(),
// 			ProductReceiveReturnId: productReceiveReturn.Id,
// 			FileId:                 imageFile.Id,
// 		})
// 		files = append(files, imageFile)
// 	}

// 	// remove stock data
// 	productStockRemovedByProductId := map[string]float64{}
// 	productStockMutations := []model.ProductStockMutation{}

// 	for _, productReceiveItem := range productReceive.ProductReceiveItems {
// 		productStockRemovedByProductId[productReceiveItem.ProductUnit.ProductId] += productReceiveItem.BaseEligibleQty()
// 	}

// 	// check sufficient amoun of stock to returned
// 	for productId, stockRemoveCount := range productStockRemovedByProductId {
// 		productStock, err := u.repositoryManager.ProductStockRepository().GetByProductId(ctx, productId)
// 		panicIfErr(err)

// 		if productStock.Qty < stockRemoveCount {
// 			panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.INSUFFICIENT_PRODUCT_STOCK"))
// 		}
// 	}

// 	// change debt status to returned
// 	debt, err := u.repositoryManager.DebtRepository().GetByDebtSourceAndDebtSourceId(ctx, data_type.DebtSourceProductReceive, productReceive.Id)
// 	panicIfErr(err)

// 	debt.Status = data_type.DebtStatusReturned

// 	panicIfErr(
// 		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
// 			debtRepository := u.repositoryManager.DebtRepository()
// 			productReceiveRepository := u.repositoryManager.ProductReceiveRepository()
// 			fileRepository := u.repositoryManager.FileRepository()
// 			deliveryOrderReturnRepository := u.repositoryManager.ProductReceiveReturnRepository()
// 			deliveryOrderReturnImageRepository := u.repositoryManager.ProductReceiveReturnImageRepository()
// 			productStockRepository := u.repositoryManager.ProductStockRepository()
// 			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

// 			if err := fileRepository.InsertMany(ctx, files); err != nil {
// 				return err
// 			}

// 			if err := debtRepository.Update(ctx, debt); err != nil {
// 				return err
// 			}

// 			if err := productReceiveRepository.Update(ctx, &productReceive); err != nil {
// 				return err
// 			}

// 			if err := deliveryOrderReturnRepository.Insert(ctx, &productReceiveReturn); err != nil {
// 				return err
// 			}

// 			if err := deliveryOrderReturnImageRepository.InsertMany(ctx, productReceiveReturnImages); err != nil {
// 				return err
// 			}

// 			for productId, stockRemoveCount := range productStockRemovedByProductId {
// 				if err := productStockRepository.UpdateDecrementQtyByProductId(ctx, productId, stockRemoveCount); err != nil {
// 					return err
// 				}
// 			}

// 			if err := productStockMutationRepository.InsertMany(ctx, productStockMutations); err != nil {
// 				return err
// 			}

// 			for _, productReceiveItem := range productReceive.ProductReceiveItems {
// 				deductQtyLeft := productReceiveItem.QtyReceived

// 				for deductQtyLeft > 0 {
// 					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx, productReceiveItem.ProductUnit.ProductId)
// 					if err != nil {
// 						return err
// 					}

// 					if deductQtyLeft > productStockMutation.BaseQtyLeft {
// 						deductQtyLeft -= productStockMutation.BaseQtyLeft
// 						productStockMutation.BaseQtyLeft = 0
// 					} else {
// 						productStockMutation.BaseQtyLeft -= deductQtyLeft
// 						deductQtyLeft = 0
// 					}

// 					if err := productStockMutationRepository.Update(ctx, productStockMutation); err != nil {
// 						return err
// 					}
// 				}
// 			}

// 			return nil
// 		}),
// 	)

// 	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
// 		productReceiveProductStock: true,
// 		productReceiveImages:       true,
// 		supplier:                   true,
// 		_return:                    true,
// 	})

// 	return productReceive
// }

func (u *productReceiveUseCase) DeleteImage(ctx context.Context, request dto_request.ProductReceiveDeleteImageRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	productReceiveImage := mustGetProductReceiveImage(ctx, u.repositoryManager, request.ProductReceiveImageId, true)

	if productReceiveImage.ProductReceiveId != productReceive.Id {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE_ITEM.NOT_FOUND"))
	}

	file := mustGetFile(ctx, u.repositoryManager, productReceiveImage.FileId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			productReceiveImageRepository := u.repositoryManager.ProductReceiveImageRepository()

			if err := productReceiveImageRepository.Delete(ctx, &productReceiveImage); err != nil {
				return err
			}

			if err := fileRepository.Delete(ctx, &file); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mainFilesystem.Delete(file.Path)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:  true,
		productReceiveImages: true,
		supplier:             true,
	})

	return productReceive
}
