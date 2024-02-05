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
	DeliveryOrderUseCase() DeliveryOrderUseCase
	PermissionUseCase() PermissionUseCase
	ProductUseCase() ProductUseCase
	ProductReceiveUseCase() ProductReceiveUseCase
	ProductStockUseCase() ProductStockUseCase
	ProductUnitUseCase() ProductUnitUseCase
	RoleUseCase() RoleUseCase
	SupplierTypeUseCase() SupplierTypeUseCase
	SupplierUseCase() SupplierUseCase
	UnitUseCase() UnitUseCase
	UserUseCase() UserUseCase
}

type useCaseManager struct {
	authUseCase           AuthUseCase
	balanceUseCase        BalanceUseCase
	cartUseCase           CartUseCase
	cashierSessionUseCase CashierSessionUseCase
	customerUseCase       CustomerUseCase
	deliveryOrderUseCase  DeliveryOrderUseCase
	permissionUseCase     PermissionUseCase
	productUseCase        ProductUseCase
	productReceiveUseCase ProductReceiveUseCase
	productStockUseCase   ProductStockUseCase
	productUnitUseCase    ProductUnitUseCase
	roleUseCase           RoleUseCase
	supplierTypeUseCase   SupplierTypeUseCase
	supplierUseCase       SupplierUseCase
	unitUseCase           UnitUseCase
	userUseCase           UserUseCase
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
func (u *useCaseManager) DeliveryOrderUseCase() DeliveryOrderUseCase {
	return u.deliveryOrderUseCase
}

func (u *useCaseManager) PermissionUseCase() PermissionUseCase {
	return u.permissionUseCase
}

func (u *useCaseManager) ProductUseCase() ProductUseCase {
	return u.productUseCase
}

func (u *useCaseManager) ProductReceiveUseCase() ProductReceiveUseCase {
	return u.productReceiveUseCase
}

func (u *useCaseManager) ProductStockUseCase() ProductStockUseCase {
	return u.productStockUseCase
}

func (u *useCaseManager) ProductUnitUseCase() ProductUnitUseCase {
	return u.productUnitUseCase
}

func (u *useCaseManager) RoleUseCase() RoleUseCase {
	return u.roleUseCase
}

func (u *useCaseManager) SupplierTypeUseCase() SupplierTypeUseCase {
	return u.supplierTypeUseCase
}

func (u *useCaseManager) SupplierUseCase() SupplierUseCase {
	return u.supplierUseCase
}

func (u *useCaseManager) UnitUseCase() UnitUseCase {
	return u.unitUseCase
}

func (u *useCaseManager) UserUseCase() UserUseCase {
	return u.userUseCase
}

func NewUseCaseManager(
	repositoryManager repository.RepositoryManager,
	filesystemManager filesystemInternal.FilesystemManager,
	jwt jwtInternal.Jwt,
	loggerStack infrastructure.LoggerStack,
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
		),
		cashierSessionUseCase: NewCashierSessionUseCase(
			repositoryManager,
		),
		customerUseCase: NewCustomerUseCase(
			repositoryManager,
		),
		deliveryOrderUseCase: NewDeliveryOrderUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		permissionUseCase: NewPermissionUseCase(
			repositoryManager,
		),
		productUseCase: NewProductUseCase(
			repositoryManager,
		),
		productReceiveUseCase: NewProductReceiveUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		productStockUseCase: NewProductStockUseCase(
			repositoryManager,
		),
		productUnitUseCase: NewProductUnitUseCase(
			repositoryManager,
			filesystemManager.Main(),
			filesystemManager.Tmp(),
		),
		roleUseCase: NewRoleUseCase(
			repositoryManager,
		),
		supplierTypeUseCase: NewSupplierTypeUseCase(
			repositoryManager,
		),
		supplierUseCase: NewSupplierUseCase(
			repositoryManager,
		),
		unitUseCase: NewUnitUseCase(
			repositoryManager,
		),
		userUseCase: NewUserUseCase(
			repositoryManager,
		),
	}
}
