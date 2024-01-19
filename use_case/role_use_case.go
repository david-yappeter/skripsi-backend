package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
)

type RoleUseCase interface {
	//
	OptionForUserForm(ctx context.Context, request dto_request.RoleOptionForUserFormRequest) ([]model.Role, int)
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

func (u *roleUseCase) OptionForUserForm(ctx context.Context, request dto_request.RoleOptionForUserFormRequest) ([]model.Role, int) {
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
