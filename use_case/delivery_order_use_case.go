package use_case

import (
	"context"
	"fmt"
	"log"
	"math"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/global"
	"myapp/infrastructure"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"
	"strings"

	"go.mau.fi/whatsmeow/binary/proto"
	"go.mau.fi/whatsmeow/types"
	"golang.org/x/sync/errgroup"
)

type deliveryOrdersLoaderParams struct {
	customer             bool
	deliveryOrderItems   bool
	deliveryOrderImages  bool
	deliveryOrderDrivers bool

	deliveryOrderItemCosts    bool
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
	FetchDriver(ctx context.Context, request dto_request.DeliveryOrderFetchDriverRequest) ([]model.DeliveryOrder, int)
	Get(ctx context.Context, request dto_request.DeliveryOrderGetRequest) model.DeliveryOrder
	ActiveForDriver(ctx context.Context) *model.DeliveryOrder
	LatestDeliveryLocation(ctx context.Context, request dto_request.LatestDeliveryLocationRequest) *model.DeliveryOrderPosition

	// update
	MarkOngoing(ctx context.Context, request dto_request.DeliveryOrderMarkOngoingRequest) model.DeliveryOrder
	Delivering(ctx context.Context, request dto_request.DeliveryOrderDeliveringRequest) model.DeliveryOrder
	Cancel(ctx context.Context, request dto_request.DeliveryOrderCancelRequest) model.DeliveryOrder
	MarkCompleted(ctx context.Context, request dto_request.DeliveryOrderMarkCompletedRequest) model.DeliveryOrder
	DeliveryLocation(ctx context.Context, request dto_request.DeliveryOrderDeliveryLocationRequest)

	// delete
	Delete(ctx context.Context, request dto_request.DeliveryOrderDeleteRequest)
	DeleteImage(ctx context.Context, request dto_request.DeliveryOrderDeleteImageRequest) model.DeliveryOrder
	DeleteItem(ctx context.Context, request dto_request.DeliveryOrderDeleteItemRequest) model.DeliveryOrder
	DeleteDriver(ctx context.Context, request dto_request.DeliveryOrderDeleteDriverRequest) model.DeliveryOrder
}

type deliveryOrderUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase baseFileUseCase

	whatsappManager *infrastructure.WhatsappManager

	mainFilesystem filesystem.Client
	tmpFilesystem  filesystem.Client
}

func NewDeliveryOrderUseCase(
	repositoryManager repository.RepositoryManager,
	whatsappManager *infrastructure.WhatsappManager,

	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) DeliveryOrderUseCase {
	return &deliveryOrderUseCase{
		repositoryManager: repositoryManager,

		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),

		whatsappManager: whatsappManager,

		mainFilesystem: mainFilesystem,
		tmpFilesystem:  tmpFilesystem,
	}
}

func (u *deliveryOrderUseCase) mustGenerateInvoiceNumber(ctx context.Context) string {
	currentDate := util.CurrentDate()
	prefix := fmt.Sprintf(
		"DO-%d%d%d-",
		currentDate.Time().Year(),
		currentDate.Time().Month(),
		currentDate.Time().Day(),
	)
	return generateSequence(ctx, u.repositoryManager, prefix)
}

func (u *deliveryOrderUseCase) mustLoadDeliveryOrdersData(ctx context.Context, deliveryOrders []*model.DeliveryOrder, option deliveryOrdersLoaderParams) {
	customerLoader := loader.NewCustomerLoader(u.repositoryManager.CustomerRepository())
	deliveryOrderItemsLoader := loader.NewDeliveryOrderItemsLoader(u.repositoryManager.DeliveryOrderItemRepository())
	deliveryOrderImagesLoader := loader.NewDeliveryOrderImagesLoader(u.repositoryManager.DeliveryOrderImageRepository())
	deliveryOrderDriversLoader := loader.NewDeliveryOrderDriversLoader(u.repositoryManager.DeliveryOrderDriverRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range deliveryOrders {
				if option.customer {
					group.Go(customerLoader.DeliveryOrderFn(deliveryOrders[i]))
				}

				if option.deliveryOrderImages {
					group.Go(deliveryOrderImagesLoader.DeliveryOrderFn(deliveryOrders[i]))
				}

				if option.deliveryOrderItems {
					group.Go(deliveryOrderItemsLoader.DeliveryOrderFn(deliveryOrders[i]))
				}

				if option.deliveryOrderDrivers {
					group.Go(deliveryOrderDriversLoader.DeliveryOrderFn(deliveryOrders[i]))
				}
			}
		}),
	)

	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())
	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	deliveryOrderItemCostsLoader := loader.NewDeliveryOrderItemCostsLoader(u.repositoryManager.DeliveryOrderItemCostRepository())
	userLoader := loader.NewUserLoader(u.repositoryManager.UserRepository())

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

				if option.deliveryOrderDrivers {
					for j := range deliveryOrders[i].DeliveryOrderDrivers {
						group.Go(userLoader.DeliveryOrderDriverFn(&deliveryOrders[i].DeliveryOrderDrivers[j]))
					}
				}

				if option.deliveryOrderItemCosts {
					for j := range deliveryOrders[i].DeliveryOrderItems {
						group.Go(deliveryOrderItemCostsLoader.DeliveryOrderItemFn(&deliveryOrders[i].DeliveryOrderItems[j]))
					}
				}
			}
		}),
	)

	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())
	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range deliveryOrders {
				for j := range deliveryOrders[i].DeliveryOrderItems {
					group.Go(productLoader.ProductUnitFn(deliveryOrders[i].DeliveryOrderItems[j].ProductUnit))
					group.Go(unitLoader.ProductUnitFn(deliveryOrders[i].DeliveryOrderItems[j].ProductUnit))

					if option.deliveryOrderItems && option.deliveryOrderProductStock {
						group.Go(productStockLoader.ProductUnitFn(deliveryOrders[i].DeliveryOrderItems[j].ProductUnit))
					}
				}
			}
		}),
	)

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range deliveryOrders {
				if option.deliveryOrderItems {
					for j := range deliveryOrders[i].DeliveryOrderItems {
						group.Go(fileLoader.ProductFn(deliveryOrders[i].DeliveryOrderItems[j].ProductUnit.Product))
					}
				}
			}
		}),
	)

	for i := range deliveryOrders {
		for j := range deliveryOrders[i].DeliveryOrderImages {
			deliveryOrders[i].DeliveryOrderImages[j].File.SetLink(u.mainFilesystem)
		}
		for j := range deliveryOrders[i].DeliveryOrderItems {
			deliveryOrders[i].DeliveryOrderItems[j].ProductUnit.Product.ImageFile.SetLink(u.mainFilesystem)
		}
	}
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

	mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	deliveryOrder := model.DeliveryOrder{
		Id:            util.NewUuid(),
		CustomerId:    request.CustomerId,
		UserId:        authUser.Id,
		InvoiceNumber: u.mustGenerateInvoiceNumber(ctx),
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
		deliveryOrder       = mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, false)
		productUnit         = mustGetProductUnitByProductIdAndUnitId(ctx, u.repositoryManager, request.ProductId, request.UnitId, true)
		product             = mustGetProduct(ctx, u.repositoryManager, request.ProductId, false)
		productStock        = shouldGetProductStockByProductId(ctx, u.repositoryManager, product.Id)
		isProductStockExist = productStock != nil

		// discount
		customer        = mustGetCustomer(ctx, u.repositoryManager, deliveryOrder.CustomerId, true)
		discountPerUnit = 0.0

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

	// add product receive item
	deliveryOrderItem := u.shouldGetDeliveryOrderItemByDeliveryOrderIdAndProductUnitId(ctx, deliveryOrder.Id, productUnit.Id)
	isNewDeliveryOrderItem := deliveryOrderItem == nil

	// check for customer discount
	if customer.CustomerTypeId != nil {
		customerTypeDiscount := shouldGetCustomerTypeDiscountByCustomerTypeIdAndProductId(ctx, u.repositoryManager, *customer.CustomerTypeId, product.Id)

		if customerTypeDiscount != nil {
			if customerTypeDiscount.DiscountAmount != nil {
				discountPerUnit = *customerTypeDiscount.DiscountAmount
			} else {
				discountPerUnit = *customerTypeDiscount.DiscountPercentage * *product.Price / 100.0
			}
		}
	}

	// add total with discount
	deliveryOrder.TotalPrice += totalSmallestQty * math.Max(*product.Price-discountPerUnit, 0)

	if isNewDeliveryOrderItem {
		deliveryOrderItem = &model.DeliveryOrderItem{
			Id:              util.NewUuid(),
			DeliveryOrderId: deliveryOrder.Id,
			ProductUnitId:   productUnit.Id,
			ScaleToBase:     productUnit.ScaleToBase,
			Qty:             request.Qty,
			PricePerUnit:    *product.Price,
			DiscountPerUnit: discountPerUnit,
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
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
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

	deliveryOrderImage := model.DeliveryOrderImage{
		Id:              util.NewUuid(),
		DeliveryOrderId: deliveryOrder.Id,
		FileId:          imageFile.Id,
		Description:     request.Description,
	}

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			deliveryOrderImageRepository := u.repositoryManager.DeliveryOrderImageRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := deliveryOrderImageRepository.Insert(ctx, &deliveryOrderImage); err != nil {
				return err
			}

			return nil
		}),
	)

	return deliveryOrder
}

func (u *deliveryOrderUseCase) AddDriver(ctx context.Context, request dto_request.DeliveryOrderAddDriverRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, false)
	mustGetUser(ctx, u.repositoryManager, request.DriverUserId, false)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	deliveryOrderDriver, err := u.repositoryManager.DeliveryOrderDriverRepository().GetByDeliveryOrderIdAndDriverUserId(ctx, deliveryOrder.Id, request.DriverUserId)
	panicIfErr(err, constant.ErrNoData)

	if deliveryOrderDriver != nil {
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
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
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
		Status:     request.Status,
		CustomerId: request.CustomerId,
		Phrase:     request.Phrase,
	}

	deliveryOrders, err := u.repositoryManager.DeliveryOrderRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.DeliveryOrderRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadDeliveryOrdersData(ctx, util.SliceValueToSlicePointer(deliveryOrders), deliveryOrdersLoaderParams{
		customer: true,
	})

	return deliveryOrders, total
}

func (u *deliveryOrderUseCase) FetchDriver(ctx context.Context, request dto_request.DeliveryOrderFetchDriverRequest) ([]model.DeliveryOrder, int) {
	authUser := model.MustGetUserCtx(ctx)

	queryOption := model.DeliveryOrderQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{},
		),
		ExcludeStatuses: []data_type.DeliveryOrderStatus{
			data_type.DeliveryOrderStatusPending,
		},
		Status:       request.Status,
		DriverUserId: &authUser.Id,
	}

	deliveryOrders, err := u.repositoryManager.DeliveryOrderRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.DeliveryOrderRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadDeliveryOrdersData(ctx, util.SliceValueToSlicePointer(deliveryOrders), deliveryOrdersLoaderParams{
		customer: true,
	})

	return deliveryOrders, total
}

func (u *deliveryOrderUseCase) Get(ctx context.Context, request dto_request.DeliveryOrderGetRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) ActiveForDriver(ctx context.Context) *model.DeliveryOrder {
	authUser := model.MustGetUserCtx(ctx)

	deliveryOrder, err := u.repositoryManager.DeliveryOrderRepository().GetByDriverUserIdAndStatusDelivering(ctx, authUser.Id)
	panicIfErr(err, constant.ErrNoData)

	if deliveryOrder != nil {
		u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{deliveryOrder}, deliveryOrdersLoaderParams{
			customer:             true,
			deliveryOrderItems:   true,
			deliveryOrderImages:  true,
			deliveryOrderDrivers: true,
		})
	}

	return deliveryOrder
}

func (u *deliveryOrderUseCase) LatestDeliveryLocation(ctx context.Context, request dto_request.LatestDeliveryLocationRequest) *model.DeliveryOrderPosition {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusDelivering {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_DELIVERING"))
	}

	deliveryOrderPosition, err := u.repositoryManager.DeliveryOrderPositionRepository().GetByDeliveryOrderId(ctx, deliveryOrder.Id)
	panicIfErr(err, constant.ErrNoData)

	return deliveryOrderPosition
}

func (u *deliveryOrderUseCase) Cancel(ctx context.Context, request dto_request.DeliveryOrderCancelRequest) model.DeliveryOrder {
	var (
		currentDateTime                              = util.CurrentDateTime()
		toBeCanceledCustomerDebt *model.CustomerDebt = nil
		productStockMutations                        = []model.ProductStockMutation{}
		err                      error
	)

	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	switch deliveryOrder.Status {
	case data_type.DeliveryOrderStatusCompleted:
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.ALREADY_COMPLETED"))

	case data_type.DeliveryOrderStatusCanceled:
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.ALREADY_CANCELED"))
	}

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

	switch deliveryOrder.Status {
	case data_type.DeliveryOrderStatusOngoing,
		data_type.DeliveryOrderStatusDelivering:
		// add canceled stock back
		u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
			deliveryOrderProductStock: true,
			deliveryOrderItems:        true,
			deliveryOrderItemCosts:    true,
		})

		for _, deliveryOrderItem := range deliveryOrder.DeliveryOrderItems {
			for _, deliveryOrderItemCost := range deliveryOrderItem.DeliveryOrderItemCosts {
				productStockMutations = append(productStockMutations, model.ProductStockMutation{
					Id:            util.NewUuid(),
					ProductUnitId: deliveryOrderItem.ProductUnitId,
					Type:          data_type.ProductStockMutationTypeDeliveryOrderItemCostCancel,
					IdentifierId:  deliveryOrderItemCost.Id,
					Qty:           deliveryOrderItemCost.Qty,
					ScaleToBase:   1,
					BaseQtyLeft:   deliveryOrderItemCost.Qty,
					BaseCostPrice: deliveryOrderItemCost.BaseCostPrice,
					MutatedAt:     currentDateTime,
				})
			}
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
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			if err := deliveryOrderRepository.Update(ctx, &deliveryOrder); err != nil {
				return err
			}

			if err := productStockMutationRepository.InsertMany(ctx, productStockMutations); err != nil {
				return err
			}

			if toBeCanceledCustomerDebt != nil {
				if err := customerDebtRepository.Update(ctx, toBeCanceledCustomerDebt); err != nil {
					return err
				}
			}

			return nil
		}),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

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
				deductQtyLeft := deliveryOrderItem.Qty

				for deductQtyLeft > 0 {
					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx, deliveryOrderItem.ProductUnit.ProductId)
					if err != nil {
						return err
					}

					if deductQtyLeft > productStockMutation.BaseQtyLeft {
						deliveryOrderItemCosts = append(deliveryOrderItemCosts, model.DeliveryOrderItemCost{
							Id:                  util.NewUuid(),
							DeliveryOrderItemId: deliveryOrderItem.Id,
							Qty:                 deductQtyLeft,
							BaseCostPrice:       productStockMutation.BaseCostPrice,
							TotalCostPrice:      productStockMutation.BaseCostPrice * productStockMutation.BaseQtyLeft * productStockMutation.ScaleToBase,
						})

						deductQtyLeft -= productStockMutation.BaseQtyLeft
						productStockMutation.BaseQtyLeft = 0
					} else {
						deliveryOrderItemCosts = append(deliveryOrderItemCosts, model.DeliveryOrderItemCost{
							Id:                  util.NewUuid(),
							DeliveryOrderItemId: deliveryOrderItem.Id,
							Qty:                 deductQtyLeft,
							BaseCostPrice:       productStockMutation.BaseCostPrice,
							TotalCostPrice:      productStockMutation.BaseCostPrice * deductQtyLeft * productStockMutation.ScaleToBase,
						})

						productStockMutation.BaseQtyLeft -= deductQtyLeft
						deductQtyLeft = 0
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
		customer:             true,
		deliveryOrderImages:  true,
		deliveryOrderItems:   true,
		deliveryOrderDrivers: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) Delivering(ctx context.Context, request dto_request.DeliveryOrderDeliveringRequest) model.DeliveryOrder {
	authUser := model.MustGetUserCtx(ctx)
	isExist, err := u.repositoryManager.DeliveryOrderRepository().IsExistByDriverUserIdAndStatusDelivering(ctx, authUser.Id)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.CURRENT_USER_ALREADY_HAVE_ACTIVE_DELIVERY"))
	}

	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusOngoing {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_ON_GOING"))
	}

	// change status
	deliveryOrder.Status = data_type.DeliveryOrderStatusDelivering

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			deliveryOrderRepository := u.repositoryManager.DeliveryOrderRepository()

			if err := deliveryOrderRepository.Update(ctx, &deliveryOrder); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		customer:             true,
		deliveryOrderImages:  true,
		deliveryOrderItems:   true,
		deliveryOrderDrivers: true,
	})

	go func() {
		if u.whatsappManager == nil {
			return
		}

		customerJID, err := types.ParseJID(fmt.Sprintf("%s@s.whatsapp.net", strings.Trim(deliveryOrder.Customer.Phone, "+")))
		if err != nil {
			log.Println(err)
			return
		}

		err = (*u.whatsappManager).SendMessage(context.Background(), customerJID, &proto.Message{
			Conversation: util.Pointer(fmt.Sprintf(
				`🚚 Pemberitahuan Pengiriman Pesanan

Halo %s,

Kami senang memberitahu Anda bahwa pesanan Anda telah diproses dan siap untuk dikirim! 🎉

Berikut adalah rincian pengiriman pesanan Anda:

📦	Nomor Pesanan: %s
📍	Alamat Pengiriman: %s
🚚	Link Live Tracking: %s

📢 Penting: Mohon pastikan ada seseorang di alamat yang tercantum untuk menerima pesanan Anda pada waktu yang ditentukan.

Jika Anda memiliki pertanyaan atau membutuhkan bantuan tambahan, jangan ragu untuk menghubungi kami di nomor ini.

Terima kasih atas kepercayaan Anda kepada kami!

Salam hangat,
%s
`,
				deliveryOrder.Customer.Name,
				deliveryOrder.InvoiceNumber,
				deliveryOrder.Customer.Address,
				fmt.Sprintf("%s/delivery-orders/testing-api/%s", global.GetConfig().Uri, deliveryOrder.Id),
				"Toko Setia Abadi",
			)),
		})

		if err != nil {
			log.Println(err)
			return
		}
	}()

	return deliveryOrder
}

func (u *deliveryOrderUseCase) MarkCompleted(ctx context.Context, request dto_request.DeliveryOrderMarkCompletedRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusDelivering {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_DELIVERING"))
	}

	deliveryOrder.Status = data_type.DeliveryOrderStatusCompleted

	panicIfErr(
		u.repositoryManager.DeliveryOrderRepository().Update(ctx, &deliveryOrder),
	)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

	go func() {
		if u.whatsappManager == nil {
			return
		}

		customerJID, err := types.ParseJID(fmt.Sprintf("%s@s.whatsapp.net", strings.Trim(deliveryOrder.Customer.Phone, "+")))
		if err != nil {
			log.Println(err)
			return
		}

		err = (*u.whatsappManager).SendMessage(context.Background(), customerJID, &proto.Message{
			Conversation: util.Pointer(fmt.Sprintf(
				`🚚 Pengiriman Selesai - Berikan Ulasan Anda!

Halo %s,

Kami senang memberitahu Anda bahwa pesanan Anda telah sukses dikirim! 🎉 Kami berharap pesanan tersebut tiba dengan baik dan memenuhi harapan Anda.

Jika Anda memiliki waktu, kami sangat menghargai ulasan dan masukan Anda tentang pengalaman berbelanja bersama kami. Ini akan membantu kami terus meningkatkan layanan kami kepada pelanggan.

🌟 Berikan Ulasan Anda: %s

Namun, jika Anda tidak memiliki waktu saat ini atau memiliki pertanyaan lebih lanjut, jangan ragu untuk menghubungi kami di nomor ini.

Terima kasih atas dukungan dan kepercayaan Anda kepada kami!

Salam hangat,
%s
`,
				deliveryOrder.Customer.Name,
				fmt.Sprintf("%s/delivery-orders/testing-api/%s", global.GetConfig().Uri, deliveryOrder.Id),
				"Toko Setia Abadi",
			)),
		})

		if err != nil {
			log.Println(err)
			return
		}
	}()

	return deliveryOrder
}

func (u *deliveryOrderUseCase) DeliveryLocation(ctx context.Context, request dto_request.DeliveryOrderDeliveryLocationRequest) {
	authUser := model.MustGetUserCtx(ctx)

	deliveryOrder, err := u.repositoryManager.DeliveryOrderRepository().GetByDriverUserIdAndStatusDelivering(ctx, authUser.Id)
	panicIfErr(err, constant.ErrNoData)

	if deliveryOrder == nil {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.NO_ACTIVE_DELIVERING"))
	}

	deliveryOrderPosition := shouldGetDeliveryOrderPositionByDeliveryOrderId(ctx, u.repositoryManager, deliveryOrder.Id)
	isNewDeliveryOrderPosition := deliveryOrderPosition == nil

	if isNewDeliveryOrderPosition {
		deliveryOrderPosition = &model.DeliveryOrderPosition{
			Id:              util.NewUuid(),
			DeliveryOrderId: deliveryOrder.Id,
			DriverUserId:    authUser.Id,
			Latitude:        request.Latitude,
			Longitude:       request.Longitude,
			Bearing:         request.Bearing,
		}
	} else {
		if authUser.Id != deliveryOrderPosition.DriverUserId {
			panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.INVALID_USER_FOR_LIVE_TRACKING"))
		}

		deliveryOrderPosition.Latitude = request.Latitude
		deliveryOrderPosition.Longitude = request.Longitude
		deliveryOrderPosition.Bearing = request.Bearing
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			deliveryOrderPositionRepository := u.repositoryManager.DeliveryOrderPositionRepository()

			if isNewDeliveryOrderPosition {
				if err := deliveryOrderPositionRepository.Insert(ctx, deliveryOrderPosition); err != nil {
					return err
				}
			} else {
				if err := deliveryOrderPositionRepository.Update(ctx, deliveryOrderPosition); err != nil {
					return err
				}
			}

			return nil
		}),
	)
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
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			deliveryOrderImageRepository := u.repositoryManager.DeliveryOrderImageRepository()

			if err := deliveryOrderImageRepository.Delete(ctx, &deliveryOrderImage); err != nil {
				return err
			}

			if err := fileRepository.Delete(ctx, &file); err != nil {
				return err
			}

			return nil
		}),
	)

	u.mainFilesystem.Delete(file.Path)

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

	return deliveryOrder
}

func (u *deliveryOrderUseCase) DeleteItem(ctx context.Context, request dto_request.DeliveryOrderDeleteItemRequest) model.DeliveryOrder {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

	// check delivery order item
	deliveryOrderItem := mustGetDeliveryOrderItem(ctx, u.repositoryManager, request.DeliveryOrderItemId, true)
	if deliveryOrderItem.DeliveryOrderId != deliveryOrder.Id {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER_ITEM.NOT_FOUND"))
	}

	// get product unit and product
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, deliveryOrderItem.ProductUnitId, true)
	product := mustGetProduct(ctx, u.repositoryManager, productUnit.ProductId, true)

	// calculate qty
	totalSmallestQty := deliveryOrderItem.Qty * productUnit.ScaleToBase

	productStock := shouldGetProductStockByProductId(ctx, u.repositoryManager, product.Id)

	// deduct delivery order total
	deliveryOrder.TotalPrice -= totalSmallestQty * math.Max(deliveryOrderItem.PricePerUnit-deliveryOrderItem.DiscountPerUnit, 0)

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
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
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
		customer:             true,
		deliveryOrderItems:   true,
		deliveryOrderImages:  true,
		deliveryOrderDrivers: true,
	})

	return deliveryOrder
}

/*
	Notes:
	- Delivery Order can only be cancel if status 'PENDING' or 'ONGOING' or 'DELIVERING'
	- Customer Debt will only be created when status changed from 'PENDING' to 'ONGOING'
	- If Delivery Order canceled, the Customer Debt will be canceled
*/
