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

	gotiktok "github.com/david-yappeter/go-tiktok"
	"golang.org/x/sync/errgroup"
)

type productReceivesLoaderParams struct {
	productReceiveItems  bool
	productReceiveImages bool

	productReceiveProductStockMutation bool
	productReceiveProductStock         bool
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
	Cancel(ctx context.Context, request dto_request.ProductReceiveCancelRequest) model.ProductReceive
	MarkComplete(ctx context.Context, request dto_request.ProductReceiveMarkCompleteRequest) model.ProductReceive

	// delete
	Delete(ctx context.Context, request dto_request.ProductReceiveDeleteRequest)
	DeleteImage(ctx context.Context, request dto_request.ProductReceiveDeleteImageRequest) model.ProductReceive
	DeleteItem(ctx context.Context, request dto_request.ProductReceiveDeleteItemRequest) model.ProductReceive
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

	productStockMutationLoader := loader.NewProductStockMutationLoader(u.repositoryManager.ProductStockMutationRepository())
	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())
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
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productReceives {
				if option.productReceiveItems && option.productReceiveProductStock {
					for j := range productReceives[i].ProductReceiveItems {
						group.Go(productStockLoader.ProductUnitFn(productReceives[i].ProductReceiveItems[j].ProductUnit))
						group.Go(productLoader.ProductUnitFn(productReceives[i].ProductReceiveItems[j].ProductUnit))
					}
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
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.NOT_FOUND"))
	}

	// add total to product receive
	productReceive.TotalPrice += request.Qty * request.PricePerUnit

	// add product receive item
	productReceiveItem := model.ProductReceiveItem{
		Id:               util.NewUuid(),
		ProductReceiveId: productReceive.Id,
		ProductUnitId:    productUnit.Id,
		UserId:           authUser.Id,
		Qty:              request.Qty,
		PricePerUnit:     request.PricePerUnit,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReceiveRepository := u.repositoryManager.ProductReceiveRepository()
			productReceiveItemRepository := u.repositoryManager.ProductReceiveItemRepository()

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

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:        true,
		productReceiveProductStock: true,
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
			toBeRemovedStockByProductId[productReceiveItem.ProductUnit.ProductId] += productReceiveItem.Qty * productReceiveItem.ProductUnit.ScaleToBase

			if productReceiveItem.ProductStockMutation != nil {
				productStockMutation := productReceiveItem.ProductStockMutation

				if productStockMutation.BaseQtyLeft != productStockMutation.Qty {
					panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.PRODUCT_ALREADY_SOLD_CANNOT_BE_CANCELED"))
				}

				toBeDeletedProductStockMutationIds = append(toBeDeletedProductStockMutationIds, productStockMutation.Id)
			}
		}

		for productId, removedStock := range toBeRemovedStockByProductId {
			productStock := productStockByProductId[productId]
			productStock.Qty -= removedStock
			productStockByProductId[productId] = productStock
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
	})

	return productReceive
}

func (u *productReceiveUseCase) MarkComplete(ctx context.Context, request dto_request.ProductReceiveMarkCompleteRequest) model.ProductReceive {
	var (
		currentDateTime                                         = util.CurrentDateTime()
		toBeAddedStockByProductId map[string]float64            = make(map[string]float64)
		productStockMutations                                   = []model.ProductStockMutation{}
		productStockByProductId   map[string]model.ProductStock = make(map[string]model.ProductStock)
	)

	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:        true,
		productReceiveImages:       true,
		productReceiveProductStock: true,
	})

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	for _, productReceiveItem := range productReceive.ProductReceiveItems {
		productStockByProductId[productReceiveItem.ProductUnit.ProductId] = *productReceiveItem.ProductUnit.ProductStock
		toBeAddedStockByProductId[productReceiveItem.ProductUnit.ProductId] += productReceiveItem.Qty * productReceiveItem.ProductUnit.ScaleToBase

		productStockMutations = append(productStockMutations, model.ProductStockMutation{
			Id:            util.NewUuid(),
			ProductUnitId: productReceiveItem.ProductUnitId,
			Type:          data_type.ProductStockMutationTypeProductReceiveItem,
			IdentifierId:  productReceiveItem.Id,
			Qty:           productReceiveItem.Qty,
			ScaleToBase:   productReceiveItem.ProductUnit.ScaleToBase,
			BaseQtyLeft:   productReceiveItem.Qty * productReceiveItem.ProductUnit.ScaleToBase,
			BaseCostPrice: productReceiveItem.PricePerUnit / productReceiveItem.ProductUnit.ScaleToBase,
			MutatedAt:     currentDateTime,
		})
	}

	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	// add stock and sync tiktok stock
	for productId, addedStock := range toBeAddedStockByProductId {
		productStock := productStockByProductId[productId]
		productStock.Qty += addedStock
		productStockByProductId[productId] = productStock

		tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, productId)

		if tiktokProduct != nil {
			tiktokProductDetail := mustGetTiktokProductDetail(ctx, u.repositoryManager, tiktokProduct.TiktokProductId)

			_, err := client.UpdateProductInventory(
				ctx,
				gotiktok.CommonParam{
					AccessToken: *tiktokConfig.AccessToken,
					ShopCipher:  tiktokConfig.ShopCipher,
					ShopId:      tiktokConfig.ShopId,
				},
				tiktokProduct.TiktokProductId,
				gotiktok.UpdateProductInventoryRequest{
					Skus: []gotiktok.UpdateProductInventoryRequestSku{
						{
							Id: tiktokProductDetail.Skus[0].Id,
							Inventory: []gotiktok.UpdateProductInventoryRequestSkuInventory{
								{
									WarehouseId: tiktokProductDetail.Skus[0].Inventory[0].WarehouseId,
									Quantity:    int(productStock.Qty),
								},
							},
						},
					},
				},
			)
			panicIfErr(err)
		}

	}

	productReceive.Status = data_type.ProductReceiveStatusCompleted

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

			if err := productStockMutationRepository.InsertMany(ctx, productStockMutations); err != nil {
				return err
			}

			return nil
		}),
	)

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

func (u *productReceiveUseCase) DeleteImage(ctx context.Context, request dto_request.ProductReceiveDeleteImageRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	file := mustGetFile(ctx, u.repositoryManager, request.FileId, true)
	productReceiveImage := mustGetProductReceiveImageByProductReceiveIdAndFileId(ctx, u.repositoryManager, request.ProductReceiveId, request.FileId, true)

	panicIfErr(
		u.repositoryManager.ProductReceiveImageRepository().Delete(ctx, &productReceiveImage),
	)

	panicIfErr(u.mainFilesystem.Delete(file.Path))

	u.mustLoadProductReceivesData(ctx, []*model.ProductReceive{&productReceive}, productReceivesLoaderParams{
		productReceiveItems:  true,
		productReceiveImages: true,
	})

	return productReceive
}

func (u *productReceiveUseCase) DeleteItem(ctx context.Context, request dto_request.ProductReceiveDeleteItemRequest) model.ProductReceive {
	productReceive := mustGetProductReceive(ctx, u.repositoryManager, request.ProductReceiveId, true)

	if productReceive.Status != data_type.ProductReceiveStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE.STATUS.MUST_BE_PENDING"))
	}

	productReceiveItem := mustGetProductReceiveItem(ctx, u.repositoryManager, request.ProductReceiveItemId, true)

	if productReceiveItem.ProductReceiveId != productReceive.Id {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_RECEIVE_ITEM.NOT_FOUND"))
	}

	// deduct total from product_receive
	productReceive.TotalPrice -= productReceiveItem.Qty * productReceiveItem.PricePerUnit

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productReceiveRepository := u.repositoryManager.ProductReceiveRepository()
			productReceiveItemRepository := u.repositoryManager.ProductReceiveItemRepository()

			if err := productReceiveRepository.Update(ctx, &productReceive); err != nil {
				return err
			}

			if err := productReceiveItemRepository.Delete(ctx, &productReceiveItem); err != nil {
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
