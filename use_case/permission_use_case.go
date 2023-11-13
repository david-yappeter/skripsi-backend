package use_case

import (
	"context"
	"myapp/data_type"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
)

type PermissionUseCase interface {
	// read
	Authorize(ctx context.Context, permissionEnum data_type.Permission, clientIp string)
}

type permissionUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewPermissionUseCase(
	repositoryManager repository.RepositoryManager,
) PermissionUseCase {
	return &permissionUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *permissionUseCase) Authorize(
	ctx context.Context,
	permissionEnum data_type.Permission,
	clientIp string,
) {
	authenticatedUser := model.MustGetUserCtx(ctx)

	permission, err := u.repositoryManager.PermissionRepository().GetByName(ctx, permissionEnum)
	if err != nil {
		panic(err)
	}

	isUserHasPermission, err := u.repositoryManager.RolePermissionRepository().IsExist(ctx, authenticatedUser.Id, permission.Id)
	if err != nil {
		panic(err)
	}

	if !isUserHasPermission {
		panic(dto_response.NewForbiddenErrorResponse("You do not have the permission for this action"))
	}
}
