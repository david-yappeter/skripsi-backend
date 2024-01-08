package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type UnitUseCase interface {
	// admin create
	AdminCreate(ctx context.Context, request dto_request.AdminUnitCreateRequest) model.Unit

	// admin read
	AdminFetch(ctx context.Context, request dto_request.AdminUnitFetchRequest) ([]model.Unit, int)
	AdminGet(ctx context.Context, request dto_request.AdminUnitGetRequest) model.Unit

	// admin update
	AdminUpdate(ctx context.Context, request dto_request.AdminUnitUpdateRequest) model.Unit

	// admin delete
	AdminDelete(ctx context.Context, request dto_request.AdminUnitDeleteRequest)
}

type unitUseCase struct {
	repositoryManager repository.RepositoryManager
}

func (u *unitUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.UnitRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("Unit name already exist"))
	}
}

func (u *unitUseCase) mustValidateAllowDeleteUnit(ctx context.Context, unitId string) {

}

func (u *unitUseCase) AdminCreate(ctx context.Context, request dto_request.AdminUnitCreateRequest) model.Unit {
	u.mustValidateNameNotDuplicate(ctx, request.Name)

	unit := model.Unit{
		Id:          util.NewUuid(),
		Name:        request.Name,
		Description: request.Description,
	}

	panicIfErr(
		u.repositoryManager.UnitRepository().Insert(ctx, &unit),
	)

	return unit
}

func (u *unitUseCase) AdminFetch(ctx context.Context, request dto_request.AdminUnitFetchRequest) ([]model.Unit, int) {
	queryOption := model.UnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	units, err := u.repositoryManager.UnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return units, total
}

func (u *unitUseCase) AdminGet(ctx context.Context, request dto_request.AdminUnitGetRequest) model.Unit {
	unit := mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)

	return unit
}

func (u *unitUseCase) AdminUpdate(ctx context.Context, request dto_request.AdminUnitUpdateRequest) model.Unit {
	unit := mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)

	if unit.Name != request.Name {
		u.mustValidateNameNotDuplicate(ctx, request.Name)
	}

	unit.Name = request.Name
	unit.Description = request.Description

	panicIfErr(
		u.repositoryManager.UnitRepository().Update(ctx, &unit),
	)

	return unit
}

func (u *unitUseCase) AdminDelete(ctx context.Context, request dto_request.AdminUnitDeleteRequest) {
	unit := mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)

	u.mustValidateAllowDeleteUnit(ctx, request.UnitId)

	panicIfErr(
		u.repositoryManager.UnitRepository().Delete(ctx, &unit),
	)
}

func NewUnitUseCase(
	repositoryManager repository.RepositoryManager,
) UnitUseCase {
	return &unitUseCase{
		repositoryManager: repositoryManager,
	}
}
