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

type deliveryOrdersLoaderParams struct {
	deliveryOrderItems  bool
	deliveryOrderImages bool

	deliveryOrderProductStock bool
}

type DeliveryOrderUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.DeliveryOrderCreateRequest) model.DeliveryOrder
	AddItem(ctx context.Context, request dto_request.DeliveryOrderAddItemRequest) model.DeliveryOrder
	AddImage(ctx context.Context, request dto_request.DeliveryOrderAddImageRequest) model.DeliveryOrder
	AddDriver(ctx context.Context, request dto_request.DeliveryOrderAddDriverRequest) model.DeliveryOrder
	Upload(ctx context.Context, request dto_request.DeliveryOrderUploadRequest) string

	// read
	Fetch(ctx context.Context, request dto_request.DeliveryOrderFetchRequest) ([]model.DeliveryOrder, int)
	Get(ctx context.Context, request dto_request.DeliveryOrderGetRequest) model.DeliveryOrder

	// update
	MarkOngoing(ctx context.Context, request dto_request.DeliveryOrderMarkOngoingRequest) model.DeliveryOrder
	Cancel(ctx context.Context, request dto_request.DeliveryOrderCancelRequest) model.DeliveryOrder
	MarkCompleted(ctx context.Context, request dto_request.DeliveryOrderMarkCompletedRequest) model.DeliveryOrder

	// delete
	Delete(ctx context.Context, request dto_request.DeliveryOrderDeleteRequest)
	DeleteImage(ctx context.Context, request dto_request.DeliveryOrderDeleteImageRequest) model.DeliveryOrder
	DeleteItem(ctx context.Context, request dto_request.DeliveryOrderDeleteItemRequest) model.DeliveryOrder
	DeleteDriver(ctx context.Context, request dto_request.DeliveryOrderDeleteDriverRequest) model.DeliveryOrder
}

type deliveryOrderUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase baseFileUseCase

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func NewDeliveryOrderUseCase(
	repositoryManager repository.RepositoryManager,

	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) DeliveryOrderUseCase {
	return &deliveryOrderUseCase{
		repositoryManager: repositoryManager,

		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

func (u *deliveryOrderUseCase) mustLoadDeliveryOrdersData(ctx context.Context, deliveryOrders []*model.DeliveryOrder, option deliveryOrdersLoaderParams) {
	deliveryOrderItemsLoader := loader.NewDeliveryOrderItemsLoader(u.repositoryManager.DeliveryOrderItemRepository())
	deliveryOrderImagesLoader := loader.NewDeliveryOrderImagesLoader(u.repositoryManager.DeliveryOrderImageRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range deliveryOrders {
				if option.deliveryOrderImages {
					group.Go(deliveryOrderImagesLoader.DeliveryOrderFn(deliveryOrders[i]))
				}

				if option.deliveryOrderItems {
					group.Go(deliveryOrderItemsLoader.DeliveryOrderFn(deliveryOrders[i]))
				}
			}
		}),
	)

	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())
	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range deliveryOrders {
				if option.deliveryOrderImages {
					for j := range deliveryOrders[i].DeliveryOrderImages {
						group.Go(fileLoader.DeliveryOrderImageFn(&deliveryOrders[i].DeliveryOrderImages[j]))
					}
				}

				if option.deliveryOrderItems {
					for j := range deliveryOrders[i].DeliveryOrderItems {
						group.Go(productUnitLoader.DeliveryOrderItemFn(&deliveryOrders[i].DeliveryOrderItems[j]))
					}
				}
			}
		}),
	)

	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range deliveryOrders {
				if option.deliveryOrderItems && option.deliveryOrderProductStock {
					for j := range deliveryOrders[i].DeliveryOrderItems {
						group.Go(productStockLoader.ProductUnitFn(deliveryOrders[i].DeliveryOrderItems[j].ProductUnit))
					}
				}
			}
		}),
	)
}

func (u *deliveryOrderUseCase) shouldGetDeliveryOrderItemByDeliveryOrderIdAndProductUnitId(ctx context.Context, deliveryOrderId string, productUnitId string) *model.DeliveryOrderItem {
	deliveryOrderItem, err := u.repositoryManager.DeliveryOrderItemRepository().GetByDeliveryOrderIdAndProductUnitId(ctx, deliveryOrderId, productUnitId)
	panicIfErr(err, constant.ErrNoData)

	return deliveryOrderItem
}

func (u *deliveryOrderUseCase) Create(ctx context.Context, request dto_request.DeliveryOrderCreateRequest) model.DeliveryOrder {
	var (
		authUser = model.MustGetUserCtx(ctx)
	)

	deliveryOrder := model.DeliveryOrder{
		Id:            util.NewUuid(),
		CustomerId:    request.CustomerId,
		UserId:        authUser.Id,
		InvoiceNumber: "",
		Date:          request.Date,
		Status:        data_type.DeliveryOrderStatusPending,
		TotalPrice:    0,
	}

	panicIfErr(
		u.repositoryManager.DeliveryOrderRepository().Insert(ctx, &deliveryOrder),
	)

	return deliveryOrder
}

func (u *deliveryOrderUseCase) AddItem(ctx context.Context, request dto_request.DeliveryOrderAddItemRequest) model.DeliveryOrder {
	var (
		authUser            = model.MustGetUserCtx(ctx)
		deliveryOrder       = mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, false)
		productUnit         = mustGetProductUnitByProductIdAndUnitId(ctx, u.repositoryManager, request.ProductId, request.UnitId, true)
		product             = mustGetProduct(ctx, u.repositoryManager, request.ProductId, false)
		productStock        = shouldGetProductStockByProductId(ctx, u.repositoryManager, product.Id)
		isProductStockExist = productStock != nil

		totalSmallestQty = request.Qty * productUnit.ScaleToBase
	)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	if !product.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.NOT_FOUND"))
	}

	if !isProductStockExist {
		productStock = &model.ProductStock{
			Id:        util.NewUuid(),
			ProductId: product.Id,
			Qty:       0,
		}
	}

	if productStock.Qty < totalSmallestQty {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.PRODUCT_OUT_OF_STOCK"))
	}

	// deduct product stock
	productStock.Qty -= totalSmallestQty

	// add total to product receive
	deliveryOrder.TotalPrice += totalSmallestQty * *product.Price

	// add product receive item
	deliveryOrderItem := u.shouldGetDeliveryOrderItemByDeliveryOrderIdAndProductUnitId(ctx, deliveryOrder.Id, productUnit.Id)
	isNewDeliveryOrderItem := deliveryOrderItem == nil

	if isNewDeliveryOrderItem {
		deliveryOrderItem = &model.DeliveryOrderItem{
			Id:              util.NewUuid(),
			DeliveryOrderId: deliveryOrder.Id,
			ProductUnitId:   productUnit.Id,
			UserId:          authUser.Id,
			Qty:             request.Qty,
			PricePerUnit:    *product.Price,
		}
	} else {
		deliveryOrderItem.Qty += request.Qty
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			deliveryOrderRepository := u.repositoryManager.DeliveryOrderRepository()
			deliveryOrderItemRepository := u.repositoryManager.DeliveryOrderItemRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()

			if isProductStockExist {
				if err := productStockRepository.Update(ctx, productStock); err != nil {
					return err
				}
			} else {
				if err := productStockRepository.Insert(ctx, productStock); err != nil {
					return err
				}
			}

			if err := deliveryOrderRepository.Update(ctx, &deliveryOrder); err != nil {
				return err
			}

			if isNewDeliveryOrderItem {
				if err := deliveryOrderItemRepository.Insert(ctx, deliveryOrderItem); err != nil {
					return err
				}
			} else {
				if err := deliveryOrderItemRepository.Update(ctx, deliveryOrderItem); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) AddImage(ctx context.Context, request dto_request.DeliveryOrderAddImageRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, false)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	imageFile := model.File{
		Id:   util.NewUuid(),
		Type: data_type.FileTypeDeliveryOrderImage,
	}

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.DeliveryOrderImagePath,
		deliveryOrder.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.FilePath)),
		request.FilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: true,
		},
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) AddDriver(ctx context.Context, request dto_request.DeliveryOrderAddDriverRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, false)
	mustGetUser(ctx, u.repositoryManager, request.DriverUserId, false)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	deliveryOrderDriver, err := u.repositoryManager.DeliveryOrderDriverRepository().Get(ctx, request.DriverUserId)
	panicIfErr(err, constant.ErrNoData)

	if deliveryOrderDriver == nil {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.DRIVER_ALREADY_ADDED"))
	}

	deliveryOrderDriver = &model.DeliveryOrderDriver{
		Id:              util.NewUuid(),
		DeliveryOrderId: request.DeliveryOrderId,
		DriverUserId:    request.DriverUserId,
	}

	panicIfErr(
		u.repositoryManager.DeliveryOrderDriverRepository().Insert(ctx, deliveryOrderDriver),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) Upload(ctx context.Context, request dto_request.DeliveryOrderUploadRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.DeliveryOrderImagePath,
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

func (u *deliveryOrderUseCase) Fetch(ctx context.Context, request dto_request.DeliveryOrderFetchRequest) ([]model.DeliveryOrder, int) {
	queryOption := model.DeliveryOrderQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	deliveryOrders, err := u.repositoryManager.DeliveryOrderRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.DeliveryOrderRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return deliveryOrders, total
}

func (u *deliveryOrderUseCase) Get(ctx context.Context, request dto_request.DeliveryOrderGetRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	return deliveryOrder
}

func (u *deliveryOrderUseCase) Cancel(ctx context.Context, request dto_request.DeliveryOrderCancelRequest) model.DeliveryOrder {
	var (
		toBeAddedStockByProductId map[string]float64            = make(map[string]float64)
		productStockByProductId   map[string]model.ProductStock = make(map[string]model.ProductStock)
		toBeCanceledCustomerDebt  *model.CustomerDebt           = nil
		err                       error
	)

	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	switch deliveryOrder.Status {
	case data_type.DeliveryOrderStatusCompleted:
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.ALREADY_COMPLETED"))

	case data_type.DeliveryOrderStatusCanceled:
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.ALREADY_CANCELED"))
	}

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	switch deliveryOrder.Status {
	case data_type.DeliveryOrderStatusOngoing:
		// add canceled stock back if already 'OnGoing' status
		u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
			deliveryOrderProductStock: true,
		})

		for _, deliveryOrderItem := range deliveryOrder.DeliveryOrderItems {
			productStockByProductId[deliveryOrderItem.ProductUnit.ProductId] = *deliveryOrderItem.ProductUnit.ProductStock
			toBeAddedStockByProductId[deliveryOrderItem.ProductUnit.ProductId] += deliveryOrderItem.Qty * deliveryOrderItem.ProductUnit.ScaleToBase
		}

		for productId, addedStock := range toBeAddedStockByProductId {
			productStock := productStockByProductId[productId]
			productStock.Qty += addedStock
			productStockByProductId[productId] = productStock
		}

		// cancel customer debt if already 'onGoing' status
		toBeCanceledCustomerDebt, err = u.repositoryManager.CustomerDebtRepository().GetByDebtSourceAndDebtSourceId(ctx, data_type.CustomerDebtDebtSourceDeliveryOrder, deliveryOrder.Id)
		panicIfErr(err)

		toBeCanceledCustomerDebt.Status = data_type.CustomerDebtStatusCanceled
	}

	// change status
	deliveryOrder.Status = data_type.DeliveryOrderStatusCanceled

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			customerDebtRepository := u.repositoryManager.CustomerDebtRepository()
			deliveryOrderRepository := u.repositoryManager.DeliveryOrderRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()

			if err := deliveryOrderRepository.Update(ctx, &deliveryOrder); err != nil {
				return err
			}

			for _, productStock := range productStockByProductId {
				if err := productStockRepository.Update(ctx, &productStock); err != nil {
					return err
				}
			}

			if toBeCanceledCustomerDebt != nil {
				if err := customerDebtRepository.Update(ctx, toBeCanceledCustomerDebt); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	return deliveryOrder
}

func (u *deliveryOrderUseCase) MarkOngoing(ctx context.Context, request dto_request.DeliveryOrderMarkOngoingRequest) model.DeliveryOrder {
	var (
		deliveryOrderItemCosts = []model.DeliveryOrderItemCost{}
	)

	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:        true,
		deliveryOrderProductStock: true,
	})

	// initialize customer debt
	customerDebt := model.CustomerDebt{
		Id:              util.NewUuid(),
		CustomerId:      deliveryOrder.CustomerId,
		DebtSource:      data_type.CustomerDebtDebtSourceDeliveryOrder,
		DebtSourceId:    deliveryOrder.Id,
		DueDate:         data_type.NewNullDate(nil),
		Status:          data_type.CustomerDebtStatusUnpaid,
		Amount:          deliveryOrder.TotalPrice,
		RemainingAmount: deliveryOrder.TotalPrice,
	}

	// change status
	deliveryOrder.Status = data_type.DeliveryOrderStatusOngoing

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			customerDebtRepository := u.repositoryManager.CustomerDebtRepository()
			deliveryOrderRepository := u.repositoryManager.DeliveryOrderRepository()
			deliveryOrderItemCostRepository := u.repositoryManager.DeliveryOrderItemCostRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			for _, deliveryOrderItem := range deliveryOrder.DeliveryOrderItems {
				productUnit := mustGetProductUnit(ctx, u.repositoryManager, deliveryOrderItem.ProductUnitId, true)

				deductQtyLeft := deliveryOrderItem.Qty

				for deductQtyLeft > 0 {
					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductUnitIdAndBaseQtyLeftNotZero(ctx, deliveryOrderItem.ProductUnitId)
					if err != nil {
						return err
					}

					if deductQtyLeft > productStockMutation.BaseQtyLeft {
						deliveryOrderItemCosts = append(deliveryOrderItemCosts, model.DeliveryOrderItemCost{
							Id:                  util.NewUuid(),
							DeliveryOrderItemId: deliveryOrderItem.Id,
							Qty:                 deductQtyLeft,
							BaseCostPrice:       productStockMutation.BaseCostPrice,
							TotalCostPrice:      productStockMutation.BaseCostPrice * productStockMutation.BaseQtyLeft * productUnit.ScaleToBase,
						})

						deductQtyLeft -= productStockMutation.BaseQtyLeft
						productStockMutation.BaseQtyLeft = 0
					} else {
						deliveryOrderItemCosts = append(deliveryOrderItemCosts, model.DeliveryOrderItemCost{
							Id:                  util.NewUuid(),
							DeliveryOrderItemId: deliveryOrderItem.Id,
							Qty:                 deductQtyLeft,
							BaseCostPrice:       productStockMutation.BaseCostPrice,
							TotalCostPrice:      productStockMutation.BaseCostPrice * deductQtyLeft * productUnit.ScaleToBase,
						})

						productStockMutation.BaseQtyLeft -= deductQtyLeft
					}

					if err := productStockMutationRepository.Update(ctx, productStockMutation); err != nil {
						return err
					}
				}
			}

			if err := deliveryOrderRepository.Update(ctx, &deliveryOrder); err != nil {
				return err
			}

			if err := deliveryOrderItemCostRepository.InsertMany(ctx, deliveryOrderItemCosts); err != nil {
				return err
			}

			if err := customerDebtRepository.Insert(ctx, &customerDebt); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) MarkCompleted(ctx context.Context, request dto_request.DeliveryOrderMarkCompletedRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	deliveryOrder.Status = data_type.DeliveryOrderStatusCompleted

	panicIfErr(
		u.repositoryManager.DeliveryOrderRepository().Update(ctx, &deliveryOrder),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) Delete(ctx context.Context, request dto_request.DeliveryOrderDeleteRequest) {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			deliveryOrderRepository := u.repositoryManager.DeliveryOrderRepository()
			deliveryOrderItemRepository := u.repositoryManager.DeliveryOrderItemRepository()
			deliveryOrderImageRepository := u.repositoryManager.DeliveryOrderImageRepository()

			if err := deliveryOrderItemRepository.DeleteManyByDeliveryOrderId(ctx, deliveryOrder.Id); err != nil {
				return err
			}

			if err := deliveryOrderImageRepository.DeleteManyByDeliveryOrderId(ctx, deliveryOrder.Id); err != nil {
				return err
			}

			if err := deliveryOrderRepository.Delete(ctx, &deliveryOrder); err != nil {
				return err
			}

			return nil
		}),
	)
}

func (u *deliveryOrderUseCase) DeleteImage(ctx context.Context, request dto_request.DeliveryOrderDeleteImageRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	file := mustGetFile(ctx, u.repositoryManager, request.FileId, true)
	deliveryOrderImage := mustGetDeliveryOrderImageByDeliveryOrderIdAndFileId(ctx, u.repositoryManager, request.DeliveryOrderId, request.FileId, true)

	panicIfErr(
		u.repositoryManager.DeliveryOrderImageRepository().Delete(ctx, &deliveryOrderImage),
	)

	panicIfErr(u.mainFilesystem.Delete(file.Path))

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) DeleteItem(ctx context.Context, request dto_request.DeliveryOrderDeleteItemRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)
	deliveryOrderItem := mustGetDeliveryOrderItemByDeliveryOrderIdAndProductUnitId(ctx, u.repositoryManager, request.DeliveryOrderId, request.ProductUnitId, true)
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, deliveryOrderItem.ProductUnitId, true)
	product := mustGetProduct(ctx, u.repositoryManager, productUnit.ProductId, true)

	totalSmallestQty := deliveryOrderItem.Qty * productUnit.ScaleToBase

	productStock := shouldGetProductStockByProductId(ctx, u.repositoryManager, product.Id)

	// deduct delivery order total
	deliveryOrder.TotalPrice -= totalSmallestQty * *product.Price

	// add product stock back
	productStock.Qty += totalSmallestQty

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			deliveryOrderRepository := u.repositoryManager.DeliveryOrderRepository()
			deliveryOrderItemRepository := u.repositoryManager.DeliveryOrderItemRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()

			if err := deliveryOrderRepository.Update(ctx, &deliveryOrder); err != nil {
				return err
			}

			if err := deliveryOrderItemRepository.Delete(ctx, &deliveryOrderItem); err != nil {
				return err
			}

			if err := productStockRepository.Update(ctx, productStock); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) DeleteDriver(ctx context.Context, request dto_request.DeliveryOrderDeleteDriverRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	mustGetUser(ctx, u.repositoryManager, request.DriverUserId, true)
	deliveryOrderDriver := mustGetDeliveryOrderDriverByDeliveryOrderIdAndUserId(ctx, u.repositoryManager, request.DeliveryOrderId, request.DriverUserId, true)

	panicIfErr(
		u.repositoryManager.DeliveryOrderDriverRepository().Delete(ctx, &deliveryOrderDriver),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}

/*
	Notes:
	- Delivery Order can only be cancel if status 'PENDING' or 'ONGOING'
	- Customer Debt will only be created when status changed from 'PENDING' to 'ONGOING'
	- If Delivery Order canceled, the Customer Debt will be canceled
*/
