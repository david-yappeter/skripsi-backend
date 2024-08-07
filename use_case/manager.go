package use_case

import (
	"myapp/infrastructure"
	filesystemInternal "myapp/internal/filesystem"
	jwtInternal "myapp/internal/jwt"
	"myapp/repository"
)

type UseCaseManager interface {
	AuthUseCase() AuthUseCase
	BalanceUseCase() BalanceUseCase
	CartUseCase() CartUseCase
	CashierSessionUseCase() CashierSessionUseCase
	CustomerUseCase() CustomerUseCase
	CustomerDebtUseCase() CustomerDebtUseCase
	CustomerTypeUseCase() CustomerTypeUseCase
	CustomerTypeDiscountUseCase() CustomerTypeDiscountUseCase
	DashboardUseCase() DashboardUseCase
	DebtUseCase() DebtUseCase
	DeliveryOrderUseCase() DeliveryOrderUseCase
	DeliveryOrderReviewUseCase() DeliveryOrderReviewUseCase
	PermissionUseCase() PermissionUseCase
	ProductDiscountUseCase() ProductDiscountUseCase
	ProductUseCase() ProductUseCase
	ProductReceiveUseCase() ProductReceiveUseCase
	ProductReturnUseCase() ProductReturnUseCase
	ProductStockAdjustmentUseCase() ProductStockAdjustmentUseCase
	ProductStockUseCase() ProductStockUseCase
	ProductUnitUseCase() ProductUnitUseCase
	PurchaseOrderUseCase() PurchaseOrderUseCase
	RoleUseCase() RoleUseCase
	ShopOrderUseCase() ShopOrderUseCase
	SupplierTypeUseCase() SupplierTypeUseCase
	SupplierUseCase() SupplierUseCase
	TiktokConfigUseCase() TiktokConfigUseCase
	TiktokProductUseCase() TiktokProductUseCase
	TransactionUseCase() TransactionUseCase
	UnitUseCase() UnitUseCase
	UserUseCase() UserUseCase
	WebhookUseCase() WebhookUseCase
	WhatsappUseCase() WhatsappUseCase
}

type useCaseManager struct {
	authUseCase                   AuthUseCase
	balanceUseCase                BalanceUseCase
	cartUseCase                   CartUseCase
	cashierSessionUseCase         CashierSessionUseCase
	customerUseCase               CustomerUseCase
	customerDebtUseCase           CustomerDebtUseCase
	customerTypeUseCase           CustomerTypeUseCase
	customerTypeDiscountUseCase   CustomerTypeDiscountUseCase
	dashboardUseCase              DashboardUseCase
	debtUseCase                   DebtUseCase
	deliveryOrderUseCase          DeliveryOrderUseCase
	deliveryOrderReviewUseCase    DeliveryOrderReviewUseCase
	permissionUseCase             PermissionUseCase
	productDiscountUseCase        ProductDiscountUseCase
	productUseCase                ProductUseCase
	productReceiveUseCase         ProductReceiveUseCase
	productReturnUseCase          ProductReturnUseCase
	productStockAdjustmentUseCase ProductStockAdjustmentUseCase
	productStockUseCase           ProductStockUseCase
	productUnitUseCase            ProductUnitUseCase
	purchaseOrderUseCase          PurchaseOrderUseCase
	roleUseCase                   RoleUseCase
	shopOrderUseCase              ShopOrderUseCase
	supplierTypeUseCase           SupplierTypeUseCase
	supplierUseCase               SupplierUseCase
	tiktokConfigUseCase           TiktokConfigUseCase
	tiktokProductUseCase          TiktokProductUseCase
	transactionUseCase            TransactionUseCase
	unitUseCase                   UnitUseCase
	userUseCase                   UserUseCase
	webhookUseCase                WebhookUseCase
	whatsappUseCase               WhatsappUseCase
}

func (u *useCaseManager) AuthUseCase() AuthUseCase {
	return u.authUseCase
}

func (u *useCaseManager) BalanceUseCase() BalanceUseCase {
	return u.balanceUseCase
}

func (u *useCaseManager) CartUseCase() CartUseCase {
	return u.cartUseCase
}

func (u *useCaseManager) CashierSessionUseCase() CashierSessionUseCase {
	return u.cashierSessionUseCase
}

func (u *useCaseManager) CustomerUseCase() CustomerUseCase {
	return u.customerUseCase
}

func (u *useCaseManager) CustomerDebtUseCase() CustomerDebtUseCase {
	return u.customerDebtUseCase
}

func (u *useCaseManager) CustomerTypeUseCase() CustomerTypeUseCase {
	return u.customerTypeUseCase
}

func (u *useCaseManager) CustomerTypeDiscountUseCase() CustomerTypeDiscountUseCase {
	return u.customerTypeDiscountUseCase
}

func (u *useCaseManager) DashboardUseCase() DashboardUseCase {
	return u.dashboardUseCase
}

func (u *useCaseManager) DebtUseCase() DebtUseCase {
	return u.debtUseCase
}

func (u *useCaseManager) DeliveryOrderUseCase() DeliveryOrderUseCase {
	return u.deliveryOrderUseCase
}

func (u *useCaseManager) DeliveryOrderReviewUseCase() DeliveryOrderReviewUseCase {
	return u.deliveryOrderReviewUseCase
}

func (u *useCaseManager) PermissionUseCase() PermissionUseCase {
	return u.permissionUseCase
}

func (u *useCaseManager) ProductDiscountUseCase() ProductDiscountUseCase {
	return u.productDiscountUseCase
}

func (u *useCaseManager) ProductUseCase() ProductUseCase {
	return u.productUseCase
}

func (u *useCaseManager) ProductReceiveUseCase() ProductReceiveUseCase {
	return u.productReceiveUseCase
}

func (u *useCaseManager) ProductReturnUseCase() ProductReturnUseCase {
	return u.productReturnUseCase
}

func (u *useCaseManager) ProductStockAdjustmentUseCase() ProductStockAdjustmentUseCase {
	return u.productStockAdjustmentUseCase
}

func (u *useCaseManager) ProductStockUseCase() ProductStockUseCase {
	return u.productStockUseCase
}

func (u *useCaseManager) ProductUnitUseCase() ProductUnitUseCase {
	return u.productUnitUseCase
}

func (u *useCaseManager) PurchaseOrderUseCase() PurchaseOrderUseCase {
	return u.purchaseOrderUseCase
}

func (u *useCaseManager) RoleUseCase() RoleUseCase {
	return u.roleUseCase
}

func (u *useCaseManager) ShopOrderUseCase() ShopOrderUseCase {
	return u.shopOrderUseCase
}

func (u *useCaseManager) SupplierTypeUseCase() SupplierTypeUseCase {
	return u.supplierTypeUseCase
}

func (u *useCaseManager) SupplierUseCase() SupplierUseCase {
	return u.supplierUseCase
}

func (u *useCaseManager) TiktokConfigUseCase() TiktokConfigUseCase {
	return u.tiktokConfigUseCase
}

func (u *useCaseManager) TiktokProductUseCase() TiktokProductUseCase {
	return u.tiktokProductUseCase
}

func (u *useCaseManager) TransactionUseCase() TransactionUseCase {
	return u.transactionUseCase
}

func (u *useCaseManager) UnitUseCase() UnitUseCase {
	return u.unitUseCase
}

func (u *useCaseManager) UserUseCase() UserUseCase {
	return u.userUseCase
}

func (u *useCaseManager) WebhookUseCase() WebhookUseCase {
	return u.webhookUseCase
}

func (u *useCaseManager) WhatsappUseCase() WhatsappUseCase {
	return u.whatsappUseCase
}

func NewUseCaseManager(
	repositoryManager repository.RepositoryManager,
	filesystemManager filesystemInternal.FilesystemManager,
	jwt jwtInternal.Jwt,
	loggerStack infrastructure.LoggerStack,
	whatsappManager *infrastructure.WhatsappManager,
) UseCaseManager {
	return &useCaseManager{
		authUseCase: NewAuthUseCase(
			repositoryManager,
			jwt,
		),
		balanceUseCase: NewBalanceUseCase(
			repositoryManager,
		),
		cartUseCase: NewCartUseCase(
			repositoryManager,
			filesystemManager.Main(),
		),
		cashierSessionUseCase: NewCashierSessionUseCase(
			repositoryManager,
		),
		customerUseCase: NewCustomerUseCase(
			repositoryManager,
		),
		customerDebtUseCase: NewCustomerDebtUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		customerTypeUseCase: NewCustomerTypeUseCase(
			repositoryManager,
		),
		customerTypeDiscountUseCase: NewCustomerTypeDiscountUseCase(
			repositoryManager,
		),
		dashboardUseCase: NewDashboardUseCase(
			repositoryManager,
		),
		debtUseCase: NewDebtUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		deliveryOrderUseCase: NewDeliveryOrderUseCase(
			repositoryManager,
			whatsappManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		deliveryOrderReviewUseCase: NewDeliveryOrderReviewUseCase(
			repositoryManager,
		),
		permissionUseCase: NewPermissionUseCase(
			repositoryManager,
		),
		productDiscountUseCase: NewProductDiscountUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		productUseCase: NewProductUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		productReceiveUseCase: NewProductReceiveUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		productReturnUseCase: NewProductReturnUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		productStockAdjustmentUseCase: NewProductStockAdjustmentUseCase(
			repositoryManager,
		),
		productStockUseCase: NewProductStockUseCase(
			repositoryManager,
		),
		productUnitUseCase: NewProductUnitUseCase(
			repositoryManager,
		),
		purchaseOrderUseCase: NewPurchaseOrderUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		roleUseCase: NewRoleUseCase(
			repositoryManager,
		),
		shopOrderUseCase: NewShopOrderUseCase(
			repositoryManager,
		),
		supplierTypeUseCase: NewSupplierTypeUseCase(
			repositoryManager,
		),
		supplierUseCase: NewSupplierUseCase(
			repositoryManager,
		),
		tiktokConfigUseCase: NewTiktokConfigUseCase(
			repositoryManager,
		),
		tiktokProductUseCase: NewTiktokProductUseCase(
			repositoryManager,
		),
		transactionUseCase: NewTransactionUseCase(
			repositoryManager,
		),
		unitUseCase: NewUnitUseCase(
			repositoryManager,
		),
		userUseCase: NewUserUseCase(
			repositoryManager,
		),
		webhookUseCase: NewWebhookUseCase(
			repositoryManager,
		),
		whatsappUseCase: NewWhatsappUseCase(
			repositoryManager,
			whatsappManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
	}
}
