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

type deliveryOrdersLoaderParams struct {
	deliveryOrderItems  bool
	deliveryOrderImages bool
}

type DeliveryOrderUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.DeliveryOrderCreateRequest) model.DeliveryOrder
	AddItem(ctx context.Context, request dto_request.DeliveryOrderAddItemRequest) model.DeliveryOrder
	AddImage(ctx context.Context, request dto_request.DeliveryOrderAddImageRequest) model.DeliveryOrder
	Upload(ctx context.Context, request dto_request.DeliveryOrderUploadRequest) string

	// read
	Fetch(ctx context.Context, request dto_request.DeliveryOrderFetchRequest) ([]model.DeliveryOrder, int)
	Get(ctx context.Context, request dto_request.DeliveryOrderGetRequest) model.DeliveryOrder

	// update
	MarkCompleted(ctx context.Context, request dto_request.DeliveryOrderMarkCompletedRequest) model.DeliveryOrder

	// delete
	Delete(ctx context.Context, request dto_request.DeliveryOrderDeleteRequest)
	DeleteImage(ctx context.Context, request dto_request.DeliveryOrderDeleteImageRequest) model.DeliveryOrder
	DeleteItem(ctx context.Context, request dto_request.DeliveryOrderDeleteItemRequest) model.DeliveryOrder
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

	// add stock
	productStock.Qty += totalSmallestQty

	// add total to product receive
	deliveryOrder.TotalPrice += totalSmallestQty * *product.Price

	// add product receive item
	deliveryOrderItem := model.DeliveryOrderItem{
		Id:              util.NewUuid(),
		DeliveryOrderId: deliveryOrder.Id,
		ProductUnitId:   productUnit.Id,
		UserId:          authUser.Id,
		Qty:             request.Qty,
		PricePerUnit:    *product.Price,
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error {
			deliveryOrderRepository := repository.NewDeliveryOrderRepository(tx, loggerStack)
			deliveryOrderItemRepository := repository.NewDeliveryOrderItemRepository(tx, loggerStack)
			productStockRepository := repository.NewProductStockRepository(tx, loggerStack)

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

			if err := deliveryOrderItemRepository.Insert(ctx, &deliveryOrderItem); err != nil {
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

func (u *deliveryOrderUseCase) Upload(ctx context.Context, request dto_request.DeliveryOrderUploadRequest) string {
	deliveryOrder := mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, false)

	if deliveryOrder.Status != data_type.DeliveryOrderStatusPending {
		panic(dto_response.NewBadRequestErrorResponse("DELIVERY_ORDER.STATUS.MUST_BE_PENDING"))
	}

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

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(tx infrastructure.DBTX, loggerStack infrastructure.LoggerStack) error {
			deliveryOrderRepository := repository.NewDeliveryOrderRepository(tx, loggerStack)
			deliveryOrderItemRepository := repository.NewDeliveryOrderItemRepository(tx, loggerStack)
			deliveryOrderImageRepository := repository.NewDeliveryOrderImageRepository(tx, loggerStack)

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

	productUnit := mustGetFile(ctx, u.repositoryManager, request.ProductUnitId, true)
	deliveryOrderItem := mustGetDeliveryOrderItemByDeliveryOrderIdAndProductUnitId(ctx, u.repositoryManager, request.DeliveryOrderId, request.ProductUnitId, true)

	panicIfErr(
		u.repositoryManager.DeliveryOrderItemRepository().Delete(ctx, &deliveryOrderItem),
	)

	panicIfErr(u.mainFilesystem.Delete(productUnit.Path))

	u.mustLoadDeliveryOrdersData(ctx, []*model.DeliveryOrder{&deliveryOrder}, deliveryOrdersLoaderParams{
		deliveryOrderItems:  true,
		deliveryOrderImages: true,
	})

	return deliveryOrder
}
