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
	UserUseCase() UserUseCase
}

type useCaseManager struct {
	authUseCase       AuthUseCase
	permissionUseCase PermissionUseCase
	userUseCase       UserUseCase
}

func (u *useCaseManager) AuthUseCase() AuthUseCase {
	return u.authUseCase
}

func (u *useCaseManager) PermissionUseCase() PermissionUseCase {
	return u.permissionUseCase
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
		userUseCase: NewUserUseCase(
			repositoryManager,
		),
	}
}
