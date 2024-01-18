package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
)

type RoleUseCase interface {
	// admin
	AdminOptionForUserForm(ctx context.Context, request dto_request.AdminRoleOptionForUserFormRequest) ([]model.Role, int)
}

type roleUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewRoleUseCase(
	repositoryManager repository.RepositoryManager,
) RoleUseCase {
	return &roleUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *roleUseCase) AdminOptionForUserForm(ctx context.Context, request dto_request.AdminRoleOptionForUserFormRequest) ([]model.Role, int) {
	queryOption := model.RoleQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(request.Page, request.Limit, model.Sorts(request.Sorts)),
		Phrase:      request.Phrase,
	}
	roles, err := u.repositoryManager.RoleRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.RoleRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return roles, total
}
