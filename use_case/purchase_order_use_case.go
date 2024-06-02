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

type purchaseOrdersLoaderParams struct {
	supplier            bool
	purchaseOrderItems  bool
	purchaseOrderImages bool
}

type PurchaseOrderUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.PurchaseOrderCreateRequest) model.PurchaseOrder
	AddItem(ctx context.Context, request dto_request.PurchaseOrderAddItemRequest) model.PurchaseOrder
	AddImage(ctx context.Context, request dto_request.PurchaseOrderAddImageRequest) model.PurchaseOrder
	Upload(ctx context.Context, request dto_request.PurchaseOrderUploadRequest) string

	// read
	Fetch(ctx context.Context, request dto_request.PurchaseOrderFetchRequest) ([]model.PurchaseOrder, int)
	Get(ctx context.Context, request dto_request.PurchaseOrderGetRequest) model.PurchaseOrder

	// update
	Update(ctx context.Context, request dto_request.PurchaseOrderUpdateRequest) model.PurchaseOrder
	Ongoing(ctx context.Context, request dto_request.PurchaseOrderOngoingRequest) model.PurchaseOrder
	Cancel(ctx context.Context, request dto_request.PurchaseOrderCancelRequest) model.PurchaseOrder
	MarkComplete(ctx context.Context, request dto_request.PurchaseOrderMarkCompleteRequest) model.PurchaseOrder

	// delete
	Delete(ctx context.Context, request dto_request.PurchaseOrderDeleteRequest)
	DeleteImage(ctx context.Context, request dto_request.PurchaseOrderDeleteImageRequest) model.PurchaseOrder
	DeleteItem(ctx context.Context, request dto_request.PurchaseOrderDeleteItemRequest) model.PurchaseOrder
}

type purchaseOrderUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase baseFileUseCase

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func NewPurchaseOrderUseCase(
	repositoryManager repository.RepositoryManager,

	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) PurchaseOrderUseCase {
	return &purchaseOrderUseCase{
		repositoryManager: repositoryManager,

		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

func (u *purchaseOrderUseCase) mustLoadPurchaseOrdersData(ctx context.Context, purchaseOrders []*model.PurchaseOrder, option purchaseOrdersLoaderParams) {
	purchaseOrderItemsLoader := loader.NewPurchaseOrderItemsLoader(u.repositoryManager.PurchaseOrderItemRepository())
	purchaseOrderImagesLoader := loader.NewPurchaseOrderImagesLoader(u.repositoryManager.PurchaseOrderImageRepository())
	supplierLoader := loader.NewSupplierLoader(u.repositoryManager.SupplierRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range purchaseOrders {
				if option.purchaseOrderImages {
					group.Go(purchaseOrderImagesLoader.PurchaseOrderFn(purchaseOrders[i]))
				}

				if option.purchaseOrderItems {
					group.Go(purchaseOrderItemsLoader.PurchaseOrderFn(purchaseOrders[i]))
				}

				if option.supplier {
					group.Go(supplierLoader.PurchaseOrderFn(purchaseOrders[i]))
				}
			}
		}),
	)

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range purchaseOrders {
				if option.purchaseOrderImages {
					for j := range purchaseOrders[i].PurchaseOrderImages {
						group.Go(fileLoader.PurchaseOrderImageFn(&purchaseOrders[i].PurchaseOrderImages[j]))
					}
				}

				if option.purchaseOrderItems {
					for j := range purchaseOrders[i].PurchaseOrderItems {
						group.Go(productUnitLoader.PurchaseOrderItemFn(&purchaseOrders[i].PurchaseOrderItems[j]))
					}
				}
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range purchaseOrders {
				for j := range purchaseOrders[i].PurchaseOrderItems {
					group.Go(productLoader.ProductUnitFn(purchaseOrders[i].PurchaseOrderItems[j].ProductUnit))
					group.Go(unitLoader.ProductUnitFn(purchaseOrders[i].PurchaseOrderItems[j].ProductUnit))
				}
			}
		}),
	)

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range purchaseOrders {
				for j := range purchaseOrders[i].PurchaseOrderItems {
					group.Go(fileLoader.ProductFn(purchaseOrders[i].PurchaseOrderItems[j].ProductUnit.Product))
				}
			}
		}),
	)

	for i := range purchaseOrders {
		for j := range purchaseOrders[i].PurchaseOrderImages {
			purchaseOrders[i].PurchaseOrderImages[j].File.SetLink(u.mainFilesystem)
		}

		for j := range purchaseOrders[i].PurchaseOrderItems {
			purchaseOrders[i].PurchaseOrderItems[j].ProductUnit.Product.ImageFile.SetLink(u.mainFilesystem)
		}
	}
}

func (u *purchaseOrderUseCase) Create(ctx context.Context, request dto_request.PurchaseOrderCreateRequest) model.PurchaseOrder {
	var (
		authUser    = model.MustGetUserCtx(ctx)
		currentDate = util.CurrentDate()
	)

	mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	purchaseOrder := model.PurchaseOrder{
		Id:                  util.NewUuid(),
		SupplierId:          request.SupplierId,
		UserId:              authUser.Id,
		InvoiceNumber:       request.InvoiceNumber,
		Date:                currentDate,
		Status:              data_type.PurchaseOrderStatusPending,
		TotalEstimatedPrice: 0,
	}

	panicIfErr(
		u.repositoryManager.PurchaseOrderRepository().Insert(ctx, &purchaseOrder),
	)

	return purchaseOrder
}

func (u *purchaseOrderUseCase) AddItem(ctx context.Context, request dto_request.PurchaseOrderAddItemRequest) model.PurchaseOrder {
	var (
		authUser      = model.MustGetUserCtx(ctx)
		purchaseOrder = mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, false)
		productUnit   = mustGetProductUnitByProductIdAndUnitId(ctx, u.repositoryManager, request.ProductId, request.UnitId, true)
		product       = mustGetProduct(ctx, u.repositoryManager, request.ProductId, false)
	)

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS_MUST_BE_PENDING"))
	}

	if !product.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.NOT_FOUND"))
	}

	// add total to product receive
	purchaseOrder.TotalEstimatedPrice += request.Qty * request.PricePerUnit

	// add product receive item
	purchaseOrderItem := model.PurchaseOrderItem{
		Id:              util.NewUuid(),
		PurchaseOrderId: purchaseOrder.Id,
		ProductUnitId:   productUnit.Id,
		UserId:          authUser.Id,
		Qty:             request.Qty,
		ScaleToBase:     productUnit.ScaleToBase,
		PricePerUnit:    request.PricePerUnit,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			purchaseOrderRepository := u.repositoryManager.PurchaseOrderRepository()
			purchaseOrderItemRepository := u.repositoryManager.PurchaseOrderItemRepository()

			if err := purchaseOrderRepository.Update(ctx, &purchaseOrder); err != nil {
				return err
			}

			if err := purchaseOrderItemRepository.Insert(ctx, &purchaseOrderItem); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems:  true,
		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}

func (u *purchaseOrderUseCase) AddImage(ctx context.Context, request dto_request.PurchaseOrderAddImageRequest) model.PurchaseOrder {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, false)

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS.MUST_BE_PENDING"))
	}

	imageFile := model.File{
		Id:   util.NewUuid(),
		Type: data_type.FileTypePurchaseOrderImage,
	}

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.PurchaseOrderImagePath,
		purchaseOrder.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.FilePath)),
		request.FilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: true,
		},
	)

	purchaseOrderImage := model.PurchaseOrderImage{
		Id:              util.NewUuid(),
		PurchaseOrderId: purchaseOrder.Id,
		FileId:          imageFile.Id,
		Description:     request.Description,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			purchaseOrderImageRepository := u.repositoryManager.PurchaseOrderImageRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := purchaseOrderImageRepository.Insert(ctx, &purchaseOrderImage); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems:  true,
		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}

func (u *purchaseOrderUseCase) Upload(ctx context.Context, request dto_request.PurchaseOrderUploadRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.PurchaseOrderImagePath,
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

func (u *purchaseOrderUseCase) Fetch(ctx context.Context, request dto_request.PurchaseOrderFetchRequest) ([]model.PurchaseOrder, int) {
	queryOption := model.PurchaseOrderQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase:     request.Phrase,
		SupplierId: request.SupplierId,
		Status:     request.Status,
	}

	purchaseOrders, err := u.repositoryManager.PurchaseOrderRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.PurchaseOrderRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadPurchaseOrdersData(ctx, util.SliceValueToSlicePointer(purchaseOrders), purchaseOrdersLoaderParams{
		supplier: true,
	})

	return purchaseOrders, total
}

func (u *purchaseOrderUseCase) Get(ctx context.Context, request dto_request.PurchaseOrderGetRequest) model.PurchaseOrder {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems:  true,
		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}

func (u *purchaseOrderUseCase) Update(ctx context.Context, request dto_request.PurchaseOrderUpdateRequest) model.PurchaseOrder {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS_MUST_BE_PENDING"))
	}

	purchaseOrder.InvoiceNumber = request.InvoiceNumber
	purchaseOrder.Date = request.Date

	panicIfErr(
		u.repositoryManager.PurchaseOrderRepository().Update(ctx, &purchaseOrder),
	)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems: true,

		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}

func (u *purchaseOrderUseCase) Ongoing(ctx context.Context, request dto_request.PurchaseOrderOngoingRequest) model.PurchaseOrder {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS_MUST_BE_PENDING"))
	}

	purchaseOrder.Status = data_type.PurchaseOrderStatusOngoing

	panicIfErr(
		u.repositoryManager.PurchaseOrderRepository().Update(ctx, &purchaseOrder),
	)

	return purchaseOrder
}

func (u *purchaseOrderUseCase) Cancel(ctx context.Context, request dto_request.PurchaseOrderCancelRequest) model.PurchaseOrder {
	var (
		toBeRemovedStockByProductId        map[string]float64            = nil
		toBeDeletedProductStockMutationIds                               = []string{}
		productStockByProductId            map[string]model.ProductStock = nil
	)

	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	if purchaseOrder.Status == data_type.PurchaseOrderStatusCanceled {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.ALREADY_CANCELED"))
	}

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems: true,
	})

	switch purchaseOrder.Status {
	case data_type.PurchaseOrderStatusCompleted:
		u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{})

		if len(purchaseOrder.PurchaseOrderItems) > 0 {
			toBeRemovedStockByProductId = make(map[string]float64)
			productStockByProductId = make(map[string]model.ProductStock)
		}

		for _, purchaseOrderItem := range purchaseOrder.PurchaseOrderItems {
			productStockByProductId[purchaseOrderItem.ProductUnit.ProductId] = *purchaseOrderItem.ProductUnit.ProductStock
			toBeRemovedStockByProductId[purchaseOrderItem.ProductUnit.ProductId] += purchaseOrderItem.Qty * purchaseOrderItem.ScaleToBase

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

	purchaseOrder.Status = data_type.PurchaseOrderStatusCanceled

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			purchaseOrderRepository := u.repositoryManager.PurchaseOrderRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			if err := purchaseOrderRepository.Update(ctx, &purchaseOrder); err != nil {
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

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}

func (u *purchaseOrderUseCase) MarkComplete(ctx context.Context, request dto_request.PurchaseOrderMarkCompleteRequest) model.PurchaseOrder {
	var (
		toBeAddedStockByProductId map[string]float64            = make(map[string]float64)
		productStockMutations                                   = []model.ProductStockMutation{}
		productStockByProductId   map[string]model.ProductStock = make(map[string]model.ProductStock)
	)

	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems:  true,
		purchaseOrderImages: true,

		supplier: true,
	})

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS.MUST_BE_PENDING"))
	}

	// add stock and sync tiktok stock
	for productId, addedStock := range toBeAddedStockByProductId {
		productStock := productStockByProductId[productId]
		productStock.Qty += addedStock
		productStockByProductId[productId] = productStock

		tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, productId)

		if tiktokProduct != nil {
			mustUpdateTiktokProductInventory(ctx, u.repositoryManager, tiktokProduct.TiktokProductId, int(productStock.Qty))
		}

	}

	// change status
	purchaseOrder.Status = data_type.PurchaseOrderStatusCompleted

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			purchaseOrderRepository := u.repositoryManager.PurchaseOrderRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			if err := purchaseOrderRepository.Update(ctx, &purchaseOrder); err != nil {
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

	return purchaseOrder
}

func (u *purchaseOrderUseCase) Delete(ctx context.Context, request dto_request.PurchaseOrderDeleteRequest) {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			purchaseOrderRepository := u.repositoryManager.PurchaseOrderRepository()
			purchaseOrderItemRepository := u.repositoryManager.PurchaseOrderItemRepository()
			purchaseOrderImageRepository := u.repositoryManager.PurchaseOrderImageRepository()

			if err := purchaseOrderItemRepository.DeleteManyByPurchaseOrderId(ctx, purchaseOrder.Id); err != nil {
				return err
			}

			if err := purchaseOrderImageRepository.DeleteManyByPurchaseOrderId(ctx, purchaseOrder.Id); err != nil {
				return err
			}

			if err := purchaseOrderRepository.Delete(ctx, &purchaseOrder); err != nil {
				return err
			}

			return nil
		}),
	)
}

func (u *purchaseOrderUseCase) DeleteImage(ctx context.Context, request dto_request.PurchaseOrderDeleteImageRequest) model.PurchaseOrder {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS.MUST_BE_PENDING"))
	}

	purchaseOrderImage := mustGetPurchaseOrderImage(ctx, u.repositoryManager, request.PurchaseOrderImageId, true)

	if purchaseOrderImage.PurchaseOrderId != purchaseOrder.Id {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER_ITEM.NOT_FOUND"))
	}

	file := mustGetFile(ctx, u.repositoryManager, purchaseOrderImage.FileId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			purchaseOrderImageRepository := u.repositoryManager.PurchaseOrderImageRepository()

			if err := purchaseOrderImageRepository.Delete(ctx, &purchaseOrderImage); err != nil {
				return err
			}

			if err := fileRepository.Delete(ctx, &file); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mainFilesystem.Delete(file.Path)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems:  true,
		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}

func (u *purchaseOrderUseCase) DeleteItem(ctx context.Context, request dto_request.PurchaseOrderDeleteItemRequest) model.PurchaseOrder {
	purchaseOrder := mustGetPurchaseOrder(ctx, u.repositoryManager, request.PurchaseOrderId, true)

	if purchaseOrder.Status != data_type.PurchaseOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER.STATUS.MUST_BE_PENDING"))
	}

	purchaseOrderItem := mustGetPurchaseOrderItem(ctx, u.repositoryManager, request.PurchaseOrderItemId, true)

	if purchaseOrderItem.PurchaseOrderId != purchaseOrder.Id {
		panic(dto_response.NewBadRequestErrorResponse("PURCHASE_ORDER_ITEM.NOT_FOUND"))
	}

	// deduct total from purchase_order
	purchaseOrder.TotalEstimatedPrice -= purchaseOrderItem.Qty * purchaseOrderItem.PricePerUnit

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			purchaseOrderRepository := u.repositoryManager.PurchaseOrderRepository()
			purchaseOrderItemRepository := u.repositoryManager.PurchaseOrderItemRepository()

			if err := purchaseOrderRepository.Update(ctx, &purchaseOrder); err != nil {
				return err
			}

			if err := purchaseOrderItemRepository.Delete(ctx, &purchaseOrderItem); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadPurchaseOrdersData(ctx, []*model.PurchaseOrder{&purchaseOrder}, purchaseOrdersLoaderParams{
		purchaseOrderItems:  true,
		purchaseOrderImages: true,
		supplier:            true,
	})

	return purchaseOrder
}
