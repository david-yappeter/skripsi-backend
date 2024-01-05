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
	Create(ctx context.Context, request dto_request.AdminUnitCreateRequest) model.Unit

	// admin read
	Fetch(ctx context.Context, request dto_request.AdminUnitFetchRequest) ([]model.Unit, int)
	Get(ctx context.Context, request dto_request.AdminUnitGetRequest) model.Unit

	// admin update
	Update(ctx context.Context, request dto_request.AdminUnitUpdateRequest) model.Unit

	// admin delete
	Delete(ctx context.Context, request dto_request.AdminUnitDeleteRequest)
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

func (u *unitUseCase) Create(ctx context.Context, request dto_request.AdminUnitCreateRequest) model.Unit {
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

func (u *unitUseCase) Fetch(ctx context.Context, request dto_request.AdminUnitFetchRequest) ([]model.Unit, int) {
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

func (u *unitUseCase) Get(ctx context.Context, request dto_request.AdminUnitGetRequest) model.Unit {
	unit := mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)

	return unit
}

func (u *unitUseCase) Update(ctx context.Context, request dto_request.AdminUnitUpdateRequest) model.Unit {
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

func (u *unitUseCase) Delete(ctx context.Context, request dto_request.AdminUnitDeleteRequest) {
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
