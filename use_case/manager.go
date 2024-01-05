package use_case

import (
	"myapp/infrastructure"
	filesystemInternal "myapp/internal/filesystem"
	jwtInternal "myapp/internal/jwt"
	"myapp/repository"
)

type UseCaseManager interface {
	AuthUseCase() AuthUseCase
	PermissionUseCase() PermissionUseCase
	ProductUseCase() ProductUseCase
	UnitUseCase() UnitUseCase
	UserUseCase() UserUseCase
}

type useCaseManager struct {
	authUseCase       AuthUseCase
	permissionUseCase PermissionUseCase
	productUseCase    ProductUseCase
	unitUseCase       UnitUseCase
	userUseCase       UserUseCase
}

func (u *useCaseManager) AuthUseCase() AuthUseCase {
	return u.authUseCase
}

func (u *useCaseManager) PermissionUseCase() PermissionUseCase {
	return u.permissionUseCase
}

func (u *useCaseManager) ProductUseCase() ProductUseCase {
	return u.productUseCase
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
		permissionUseCase: NewPermissionUseCase(
			repositoryManager,
		),
		productUseCase: NewProductUseCase(
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
