package use_case

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	validatorInternal "myapp/internal/gin/validator"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"sync"
)

const (
	extensionTypeWord       = "word"
	extensionTypeExcel      = "excel"
	extensionTypePowerPoint = "powerpoint"
	extensionTypePdf        = "pdf"
	extensionTypeImage      = "image"
	extensionTypeGif        = "gif"
	extensionTypeAudio      = "audio"
	extensionTypeVideo      = "video"
	extensionTypeCompressed = "compressed"
	extensionTypeMedical    = "medical"
)

var (
	webhookMutexes = make(map[string]*sync.Mutex, 0)

	Validator validatorInternal.Validator = validatorInternal.New()

	extensions = map[string][]string{
		extensionTypeWord: {
			".docs",
			".doc",
			".docx",
		},

		extensionTypeExcel: {
			".xlsx",
			".xls",
			".xltx",
			".xlsb",
			".csv",
		},

		extensionTypePowerPoint: {
			".ppt",
			".pptx",
		},

		extensionTypePdf: {
			".pdf",
		},

		extensionTypeImage: {
			".jpeg",
			".jpg",
			".png",
			".jfif",
		},

		extensionTypeGif: {
			".gif",
		},

		extensionTypeAudio: {
			".mp3",
			".mpeg",
		},

		extensionTypeVideo: {
			".mp4",
		},

		extensionTypeCompressed: {
			".zip",
		},

		extensionTypeMedical: {
			".dcm",
			".dicom",
			".dicm",
			".DCM",
			".DICOM",
		},
	}
)

func listSupportedExtension(extensionTypes []string) []string {
	supportedExtensions := []string{}
	for _, extensionType := range extensionTypes {
		supportedExtensions = append(supportedExtensions, extensions[extensionType]...)
	}

	return supportedExtensions
}

type FilesystemCopy struct {
	Filesystem filesystem.Client
	Path       string
}

func (u FilesystemCopy) CopyTo(ctx context.Context, dest FilesystemCopy) error {
	if u.Filesystem == nil || u.Path == "" {
		panic("source filesystem and path must not empty")
	}

	if dest.Filesystem == nil || dest.Path == "" {
		panic("destination filesystem and path must not empty")
	}

	reader, err := u.Filesystem.Open(u.Path)
	if err != nil {
		return err
	}
	defer reader.Close()

	if err := dest.Filesystem.Write(ctx, reader, dest.Path); err != nil {
		return err
	}

	return nil
}

func (u FilesystemCopy) MustCopyTo(ctx context.Context, dest FilesystemCopy) {
	err := u.CopyTo(ctx, dest)
	if err != nil {
		panic(err)
	}
}

func panicIfErr(err error, excludedErrs ...error) {
	if err != nil {
		for _, excludedErr := range excludedErrs {
			if err == excludedErr {
				return
			}
		}
		panic(err)
	}
}

func panicIfRepositoryError(err error, errNoDataValidateMessage string, isValidate bool) {
	if err != nil {
		if err == constant.ErrNoData {
			if isValidate {
				panic(dto_response.NewBadRequestErrorResponse(errNoDataValidateMessage))
			}

			panic(dto_response.NewNotFoundErrorResponse("DATA_NOT_FOUND"))
		}

		panic(err)
	}
}

func mustGetUser(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.User {
	user, err := repositoryManager.UserRepository().Get(ctx, id)
	panicIfRepositoryError(err, "USER.NOT_FOUND", isValidate)
	return *user
}

func mustGetRole(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Role {
	role, err := repositoryManager.RoleRepository().GetById(ctx, id)
	panicIfRepositoryError(err, "ROLE.NOT_FOUND", isValidate)
	return *role
}

func mustGetUserRoleByUserIdAndRoleId(ctx context.Context, repositoryManager repository.RepositoryManager, userId string, roleId string, isValidate bool) model.UserRole {
	userRole, err := repositoryManager.UserRoleRepository().GetByUserIdAndRoleId(ctx, userId, roleId)
	panicIfRepositoryError(err, "USER_ROLE.NOT_FOUND", isValidate)
	return *userRole
}

func mustGetUnit(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Unit {
	unit, err := repositoryManager.UnitRepository().Get(ctx, id)
	panicIfRepositoryError(err, "UNIT.NOT_FOUND", isValidate)
	return *unit
}

func mustGetProduct(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Product {
	product, err := repositoryManager.ProductRepository().Get(ctx, id)
	panicIfRepositoryError(err, "PRODUCT.NOT_FOUND", isValidate)
	return *product
}

func mustGetSupplierType(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.SupplierType {
	supplierType, err := repositoryManager.SupplierTypeRepository().Get(ctx, id)
	panicIfRepositoryError(err, "SUPPLIER_TYPE.NOT_FOUND", isValidate)
	return *supplierType
}

func mustGetSupplier(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Supplier {
	supplier, err := repositoryManager.SupplierRepository().Get(ctx, id)
	panicIfRepositoryError(err, "SUPPLIER.NOT_FOUND", isValidate)
	return *supplier
}

func mustGetCustomer(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Customer {
	customer, err := repositoryManager.CustomerRepository().Get(ctx, id)
	panicIfRepositoryError(err, "CUSTOMER.NOT_FOUND", isValidate)
	return *customer
}

func mustGetProductUnit(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.ProductUnit {
	productUnit, err := repositoryManager.ProductUnitRepository().Get(ctx, id)
	panicIfRepositoryError(err, "PRODUCT_UNIT.NOT_FOUND", isValidate)
	return *productUnit
}

func mustGetProductUnitByProductIdAndUnitId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string, unitId string, isValidate bool) model.ProductUnit {
	productUnit, err := repositoryManager.ProductUnitRepository().GetByProductIdAndUnitId(ctx, productId, unitId)
	panicIfRepositoryError(err, "PRODUCT_UNIT.NOT_FOUND", isValidate)
	return *productUnit
}

func mustGetProductDiscount(ctx context.Context, repositoryManager repository.RepositoryManager, productDiscountId string, isValidate bool) model.ProductDiscount {
	productDiscount, err := repositoryManager.ProductDiscountRepository().Get(ctx, productDiscountId)
	panicIfRepositoryError(err, "PRODUCT_DISCOUNT.NOT_FOUND", isValidate)
	return *productDiscount
}

func mustGetFile(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.File {
	file, err := repositoryManager.FileRepository().Get(ctx, id)
	panicIfRepositoryError(err, "FILE.NOT_FOUND", isValidate)
	return *file
}

func mustGetBalance(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Balance {
	balance, err := repositoryManager.BalanceRepository().Get(ctx, id)
	panicIfRepositoryError(err, "BALANCE.NOT_FOUND", isValidate)
	return *balance
}

func mustGetTransaction(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.Transaction {
	transaction, err := repositoryManager.TransactionRepository().Get(ctx, id)
	panicIfRepositoryError(err, "TRANSACTION.NOT_FOUND", isValidate)
	return *transaction
}

func mustGetProductReceive(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.ProductReceive {
	productReceive, err := repositoryManager.ProductReceiveRepository().Get(ctx, id)
	panicIfRepositoryError(err, "PRODUCT_RECEIVE.NOT_FOUND", isValidate)
	return *productReceive
}

func mustGetProductReceiveItem(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.ProductReceiveItem {
	productReceiveItem, err := repositoryManager.ProductReceiveItemRepository().Get(ctx, id)
	panicIfRepositoryError(err, "PRODUCT_RECEIVE_ITEM.NOT_FOUND", isValidate)
	return *productReceiveItem
}

func mustGetProductReceiveImage(ctx context.Context, repositoryManager repository.RepositoryManager, productReceiveImageId string, isValidate bool) model.ProductReceiveImage {
	productReceiveImage, err := repositoryManager.ProductReceiveImageRepository().Get(ctx, productReceiveImageId)
	panicIfRepositoryError(err, "PRODUCT_RECEIVE_IMAGE.NOT_FOUND", isValidate)
	return *productReceiveImage
}

func mustGetPurchaseOrder(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.PurchaseOrder {
	purchaseOrder, err := repositoryManager.PurchaseOrderRepository().Get(ctx, id)
	panicIfRepositoryError(err, "PURCHASE_ORDER.NOT_FOUND", isValidate)
	return *purchaseOrder
}

func mustGetPurchaseOrderItem(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.PurchaseOrderItem {
	purchaseOrderItem, err := repositoryManager.PurchaseOrderItemRepository().Get(ctx, id)
	panicIfRepositoryError(err, "PURCHASE_ORDER_ITEM.NOT_FOUND", isValidate)
	return *purchaseOrderItem
}

func mustGetPurchaseOrderImage(ctx context.Context, repositoryManager repository.RepositoryManager, purchaseOrderImageId string, isValidate bool) model.PurchaseOrderImage {
	purchaseOrderImage, err := repositoryManager.PurchaseOrderImageRepository().Get(ctx, purchaseOrderImageId)
	panicIfRepositoryError(err, "PRODUCT_RECEIVE_IMAGE.NOT_FOUND", isValidate)
	return *purchaseOrderImage
}

func mustGetProductReceiveImageByProductReceiveIdAndFileId(ctx context.Context, repositoryManager repository.RepositoryManager, productReceiveId string, fileId string, isValidate bool) model.ProductReceiveImage {
	productReceiveImage, err := repositoryManager.ProductReceiveImageRepository().GetByProductReceiveIdAndFileId(ctx, productReceiveId, fileId)
	panicIfRepositoryError(err, "PRODUCT_RECEIVE_IMAGE.NOT_FOUND", isValidate)
	return *productReceiveImage
}

func mustGetProductReceiveItemByProductReceiveIdAndProductUnitId(ctx context.Context, repositoryManager repository.RepositoryManager, productReceiveId string, productUnitId string, isValidate bool) model.ProductReceiveItem {
	productReceiveItem, err := repositoryManager.ProductReceiveItemRepository().GetByProductReceiveIdAndProductUnitId(ctx, productReceiveId, productUnitId)
	panicIfRepositoryError(err, "PRODUCT_RECEIVE_ITEM.NOT_FOUND", isValidate)
	return *productReceiveItem
}

func mustGetDeliveryOrder(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.DeliveryOrder {
	deliveryOrder, err := repositoryManager.DeliveryOrderRepository().Get(ctx, id)
	panicIfRepositoryError(err, "DELIVERY_ORDER.NOT_FOUND", isValidate)
	return *deliveryOrder
}

func mustGetDeliveryOrderPosition(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.DeliveryOrderPosition {
	deliveryOrderPosition, err := repositoryManager.DeliveryOrderPositionRepository().Get(ctx, id)
	panicIfRepositoryError(err, "DELIVERY_ORDER_POSITION.NOT_FOUND", isValidate)
	return *deliveryOrderPosition
}

func mustGetDeliveryOrderItem(ctx context.Context, repositoryManager repository.RepositoryManager, id string, isValidate bool) model.DeliveryOrderItem {
	deliveryOrderItem, err := repositoryManager.DeliveryOrderItemRepository().Get(ctx, id)
	panicIfRepositoryError(err, "DELIVERY_ORDER_ITEM.NOT_FOUND", isValidate)
	return *deliveryOrderItem
}

func mustGetDeliveryOrderImageByDeliveryOrderIdAndFileId(ctx context.Context, repositoryManager repository.RepositoryManager, deliveryOrderId string, fileId string, isValidate bool) model.DeliveryOrderImage {
	deliveryOrderImage, err := repositoryManager.DeliveryOrderImageRepository().GetByDeliveryOrderIdAndFileId(ctx, deliveryOrderId, fileId)
	panicIfRepositoryError(err, "DELIVERY_ORDER_IMAGE.NOT_FOUND", isValidate)
	return *deliveryOrderImage
}

func mustGetDeliveryOrderImageByIdAndDeliveryOrderId(ctx context.Context, repositoryManager repository.RepositoryManager, id string, deliveryOrderId string, isValidate bool) model.DeliveryOrderImage {
	deliveryOrderImage, err := repositoryManager.DeliveryOrderImageRepository().GetByIdAndDeliveryOrderId(ctx, id, deliveryOrderId)
	panicIfRepositoryError(err, "DELIVERY_ORDER_IMAGE.NOT_FOUND", isValidate)
	return *deliveryOrderImage
}

func mustGetDeliveryOrderItemByDeliveryOrderIdAndProductUnitId(ctx context.Context, repositoryManager repository.RepositoryManager, deliveryOrderId string, productUnitId string, isValidate bool) model.DeliveryOrderItem {
	deliveryOrderItem, err := repositoryManager.DeliveryOrderItemRepository().GetByDeliveryOrderIdAndProductUnitId(ctx, deliveryOrderId, productUnitId)
	panicIfRepositoryError(err, "DELIVERY_ORDER_ITEM.NOT_FOUND", isValidate)
	return *deliveryOrderItem
}

func mustGetDeliveryOrderDriverByDeliveryOrderIdAndUserId(ctx context.Context, repositoryManager repository.RepositoryManager, deliveryOrderId string, driverId string, isValidate bool) model.DeliveryOrderDriver {
	deliveryOrderDriver, err := repositoryManager.DeliveryOrderDriverRepository().GetByDeliveryOrderIdAndDriverUserId(ctx, deliveryOrderId, driverId)
	panicIfRepositoryError(err, "DELIVERY_ORDER_ITEM.NOT_FOUND", isValidate)
	return *deliveryOrderDriver
}

func mustGetCashierSession(ctx context.Context, repositoryManager repository.RepositoryManager, cashierSessionId string, isValidate bool) model.CashierSession {
	cashierSession, err := repositoryManager.CashierSessionRepository().Get(ctx, cashierSessionId)
	panicIfRepositoryError(err, "CASHIER_SESSION.NOT_FOUND", isValidate)
	return *cashierSession
}

func mustGetCart(ctx context.Context, repositoryManager repository.RepositoryManager, cartId string, isValidate bool) model.Cart {
	cart, err := repositoryManager.CartRepository().Get(ctx, cartId)
	panicIfRepositoryError(err, "CART.NOT_FOUND", isValidate)
	return *cart
}

func mustGetCartItem(ctx context.Context, repositoryManager repository.RepositoryManager, cartItemId string, isValidate bool) model.CartItem {
	cartItem, err := repositoryManager.CartItemRepository().Get(ctx, cartItemId)
	panicIfRepositoryError(err, "CART_ITEM.NOT_FOUND", isValidate)
	return *cartItem
}

func mustGetProductStock(ctx context.Context, repositoryManager repository.RepositoryManager, productStockId string, isValidate bool) model.ProductStock {
	productStock, err := repositoryManager.ProductStockRepository().Get(ctx, productStockId)
	panicIfRepositoryError(err, "PRODUCT_STOCK.NOT_FOUND", isValidate)
	return *productStock
}

func mustGetProductStockByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string, isValidate bool) model.ProductStock {
	productStock, err := repositoryManager.ProductStockRepository().GetByProductId(ctx, productId)
	panicIfRepositoryError(err, "PRODUCT_STOCK.NOT_FOUND", isValidate)
	return *productStock
}

func mustGetCustomerDebt(ctx context.Context, repositoryManager repository.RepositoryManager, customerDebtId string, isValidate bool) model.CustomerDebt {
	customerDebt, err := repositoryManager.CustomerDebtRepository().Get(ctx, customerDebtId)
	panicIfRepositoryError(err, "CUSTOMER_DEBT.NOT_FOUND", isValidate)
	return *customerDebt
}

func mustGetTiktokProduct(ctx context.Context, repositoryManager repository.RepositoryManager, tiktokProductId string, isValidate bool) model.TiktokProduct {
	tiktokProduct, err := repositoryManager.TiktokProductRepository().Get(ctx, tiktokProductId)
	panicIfRepositoryError(err, "TIKTOK_PRODUCT.NOT_FOUND", isValidate)
	return *tiktokProduct
}

func mustGetTiktokProductByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string, isValidate bool) model.TiktokProduct {
	tiktokProduct := shouldGetTiktokProductByProductId(ctx, repositoryManager, productId)
	if tiktokProduct == nil {
		panicIfRepositoryError(constant.ErrNoData, "TIKTOK_PRODUCT.NOT_FOUND", isValidate)
	}
	return *tiktokProduct
}

func mustGetShopOrder(ctx context.Context, repositoryManager repository.RepositoryManager, shopOrderId string, isValidate bool) model.ShopOrder {
	shopOrder, err := repositoryManager.ShopOrderRepository().Get(ctx, shopOrderId)
	panicIfRepositoryError(err, "SHOP_ORDER.NOT_FOUND", isValidate)
	return *shopOrder
}

func mustGetDeliveryOrderReview(ctx context.Context, repositoryManager repository.RepositoryManager, deliveryOrderReviewId string, isValidate bool) model.DeliveryOrderReview {
	deliveryOrderReview, err := repositoryManager.DeliveryOrderReviewRepository().Get(ctx, deliveryOrderReviewId)
	panicIfRepositoryError(err, "DELIVERY_ORDER_REVIEW.NOT_FOUND", isValidate)
	return *deliveryOrderReview
}

func mustGetCustomerType(ctx context.Context, repositoryManager repository.RepositoryManager, customerTypeId string, isValidate bool) model.CustomerType {
	customerType, err := repositoryManager.CustomerTypeRepository().Get(ctx, customerTypeId)
	panicIfRepositoryError(err, "CUSTOMER_TYPE.NOT_FOUND", isValidate)
	return *customerType
}

func mustGetCustomerTypeDiscount(ctx context.Context, repositoryManager repository.RepositoryManager, customerTypeDiscountId string, isValidate bool) model.CustomerTypeDiscount {
	customerTypeDiscount, err := repositoryManager.CustomerTypeDiscountRepository().Get(ctx, customerTypeDiscountId)
	panicIfRepositoryError(err, "CUSTOMER_TYPE_DISCOUNT.NOT_FOUND", isValidate)
	return *customerTypeDiscount
}

func mustGetDebt(ctx context.Context, repositoryManager repository.RepositoryManager, debtId string, isValidate bool) model.Debt {
	debt, err := repositoryManager.DebtRepository().Get(ctx, debtId)
	panicIfRepositoryError(err, "DEBT.NOT_FOUND", isValidate)
	return *debt
}

func mustGetProductDiscountByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string, isValidate bool) model.ProductDiscount {
	productDiscount := shouldGetProductDiscountByProductId(ctx, repositoryManager, productId)

	if productDiscount == nil {
		panicIfRepositoryError(constant.ErrNoData, "PRODUCT_DISCOUNT.NOT_FOUND", isValidate)
	}

	return *productDiscount
}

func mustDeliveryOrderPositionByDeliveryOrderId(ctx context.Context, repositoryManager repository.RepositoryManager, deliveryOrderId string, isValidate bool) model.DeliveryOrderPosition {
	deliveryOrderPosition := shouldGetDeliveryOrderPositionByDeliveryOrderId(ctx, repositoryManager, deliveryOrderId)

	if deliveryOrderPosition == nil {
		panicIfRepositoryError(constant.ErrNoData, "DELIVERY_ORDER_POSITION.NOT_FOUND", isValidate)
	}

	return *deliveryOrderPosition
}

func mustGetCustomerTypeDiscountByCustomerTypeIdAndCustomerTypeDiscountId(ctx context.Context, repositoryManager repository.RepositoryManager, customerTypeId string, customerTypeDiscountId string, isValidate bool) model.CustomerTypeDiscount {
	customerTypeDiscount := shouldGetCustomerTypeDiscountByCustomerTypeIdAndCustomerTypeDiscountId(ctx, repositoryManager, customerTypeId, customerTypeDiscountId)

	if customerTypeDiscount == nil {
		panicIfRepositoryError(constant.ErrNoData, "CUSTOMER_TYPE_DISCOUNT.NOT_FOUND", isValidate)
	}

	return *customerTypeDiscount
}

func mustGetCustomerTypeDiscountByCustomerTypeIdAndProductId(ctx context.Context, repositoryManager repository.RepositoryManager, customerTypeId string, productId string, isValidate bool) model.CustomerTypeDiscount {
	customerTypeDiscount := shouldGetCustomerTypeDiscountByCustomerTypeIdAndProductId(ctx, repositoryManager, customerTypeId, productId)

	if customerTypeDiscount == nil {
		panicIfRepositoryError(constant.ErrNoData, "CUSTOMER_TYPE_DISCOUNT.NOT_FOUND", isValidate)
	}

	return *customerTypeDiscount
}

func shouldGetTiktokProductByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string) *model.TiktokProduct {
	tiktokProduct, err := repositoryManager.TiktokProductRepository().GetByProductId(ctx, productId)
	panicIfErr(err, constant.ErrNoData)
	return tiktokProduct
}

func shouldGetProductStockByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string) *model.ProductStock {
	productStock, err := repositoryManager.ProductStockRepository().GetByProductId(ctx, productId)
	panicIfErr(err, constant.ErrNoData)
	return productStock
}

func shouldGetProductDiscountByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string) *model.ProductDiscount {
	productDiscount, err := repositoryManager.ProductDiscountRepository().GetByProductId(ctx, productId)
	panicIfErr(err, constant.ErrNoData)
	return productDiscount
}

func shouldGetBaseProductUnitByProductId(ctx context.Context, repositoryManager repository.RepositoryManager, productId string) *model.ProductUnit {
	productUnit, err := repositoryManager.ProductUnitRepository().GetBaseProductUnitByProductId(ctx, productId)
	panicIfErr(err, constant.ErrNoData)
	return productUnit
}

func shouldGetDeliveryOrderPositionByDeliveryOrderId(ctx context.Context, repositoryManager repository.RepositoryManager, deliveryOrderId string) *model.DeliveryOrderPosition {
	deliveryOrderPosition, err := repositoryManager.DeliveryOrderPositionRepository().GetByDeliveryOrderId(ctx, deliveryOrderId)
	panicIfErr(err, constant.ErrNoData)
	return deliveryOrderPosition
}

func shouldGetCustomerTypeDiscountByCustomerTypeIdAndCustomerTypeDiscountId(ctx context.Context, repositoryManager repository.RepositoryManager, customerTypeId string, customerTypeDiscountId string) *model.CustomerTypeDiscount {
	customerTypeDiscount, err := repositoryManager.CustomerTypeDiscountRepository().GetByIdAndCustomerTypeId(ctx, customerTypeDiscountId, customerTypeId)
	panicIfErr(err, constant.ErrNoData)
	return customerTypeDiscount
}

func shouldGetCustomerTypeDiscountByCustomerTypeIdAndProductId(ctx context.Context, repositoryManager repository.RepositoryManager, customerTypeId string, productId string) *model.CustomerTypeDiscount {
	customerTypeDiscount, err := repositoryManager.CustomerTypeDiscountRepository().GetByCustomerTypeIdAndProductId(ctx, customerTypeId, productId)
	panicIfErr(err, constant.ErrNoData)
	return customerTypeDiscount
}

func execWebhookMutex(fn func()) {
	m, exist := webhookMutexes["single"]
	if !exist {
		m = new(sync.Mutex)
		webhookMutexes["single"] = m
	}

	m.Lock()
	defer m.Unlock()

	fn()
}

func generateSequence(ctx context.Context, repositoryManager repository.RepositoryManager, uniqueIdentifier string) string {
	sequence, err := repositoryManager.SequenceRepository().GetLatestByUniqueIdentifier(ctx, uniqueIdentifier)
	panicIfErr(err, constant.ErrNoData)

	if sequence == nil {
		sequence = &model.Sequence{
			Id:               util.NewUuid(),
			UniqueIdentifier: uniqueIdentifier,
			Sequence:         1,
		}
	} else {
		sequence.Id = util.NewUuid()
		sequence.Sequence++
	}

	panicIfErr(
		repositoryManager.SequenceRepository().Insert(ctx, sequence),
	)

	return fmt.Sprintf("%s%04d", sequence.UniqueIdentifier, sequence.Sequence)
}
