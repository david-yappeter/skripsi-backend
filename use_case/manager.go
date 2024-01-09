package use_case

import (
	"myapp/infrastructure"
	filesystemInternal "myapp/internal/filesystem"
	jwtInternal "myapp/internal/jwt"
	"myapp/repository"
)

type UseCaseManager interface {
	AuthUseCase() AuthUseCase
	CustomerUseCase() CustomerUseCase
	PermissionUseCase() PermissionUseCase
	ProductUseCase() ProductUseCase
	SupplierTypeUseCase() SupplierTypeUseCase
	SupplierUseCase() SupplierUseCase
	UnitUseCase() UnitUseCase
	UserUseCase() UserUseCase
}

type useCaseManager struct {
	authUseCase         AuthUseCase
	customerUseCase     CustomerUseCase
	permissionUseCase   PermissionUseCase
	productUseCase      ProductUseCase
	supplierTypeUseCase SupplierTypeUseCase
	supplierUseCase     SupplierUseCase
	unitUseCase         UnitUseCase
	userUseCase         UserUseCase
}

func (u *useCaseManager) AuthUseCase() AuthUseCase {
	return u.authUseCase
}

func (u *useCaseManager) CustomerUseCase() CustomerUseCase {
	return u.customerUseCase
}

func (u *useCaseManager) PermissionUseCase() PermissionUseCase {
	return u.permissionUseCase
}

func (u *useCaseManager) ProductUseCase() ProductUseCase {
	return u.productUseCase
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
		customerUseCase: NewCustomerUseCase(
			repositoryManager,
		),
		permissionUseCase: NewPermissionUseCase(
			repositoryManager,
		),
		productUseCase: NewProductUseCase(
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
