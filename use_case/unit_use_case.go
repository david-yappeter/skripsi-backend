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
	//  create
	Create(ctx context.Context, request dto_request.UnitCreateRequest) model.Unit

	//  read
	Fetch(ctx context.Context, request dto_request.UnitFetchRequest) ([]model.Unit, int)
	Get(ctx context.Context, request dto_request.UnitGetRequest) model.Unit

	//  update
	Update(ctx context.Context, request dto_request.UnitUpdateRequest) model.Unit

	//  delete
	Delete(ctx context.Context, request dto_request.UnitDeleteRequest)

	// option
	OptionForProductUnitForm(ctx context.Context, request dto_request.UnitOptionForProductUnitFormRequest) ([]model.Unit, int)
	OptionForProductUnitToUnitForm(ctx context.Context, request dto_request.UnitOptionForProductUnitToUnitFormRequest) ([]model.Unit, int)
}

type unitUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewUnitUseCase(
	repositoryManager repository.RepositoryManager,
) UnitUseCase {
	return &unitUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *unitUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.UnitRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("UNIT.NAME.ALREADY_EXIST"))
	}
}

func (u *unitUseCase) mustValidateAllowDeleteUnit(ctx context.Context, unitId string) {

}

func (u *unitUseCase) Create(ctx context.Context, request dto_request.UnitCreateRequest) model.Unit {
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

func (u *unitUseCase) Fetch(ctx context.Context, request dto_request.UnitFetchRequest) ([]model.Unit, int) {
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

func (u *unitUseCase) Get(ctx context.Context, request dto_request.UnitGetRequest) model.Unit {
	unit := mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)

	return unit
}

func (u *unitUseCase) Update(ctx context.Context, request dto_request.UnitUpdateRequest) model.Unit {
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

func (u *unitUseCase) Delete(ctx context.Context, request dto_request.UnitDeleteRequest) {
	unit := mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)

	u.mustValidateAllowDeleteUnit(ctx, request.UnitId)

	panicIfErr(
		u.repositoryManager.UnitRepository().Delete(ctx, &unit),
	)
}

func (u *unitUseCase) OptionForProductUnitForm(ctx context.Context, request dto_request.UnitOptionForProductUnitFormRequest) ([]model.Unit, int) {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	queryOption := model.UnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{
				{Field: "name", Direction: "asc"},
			},
		),
		ProductIdNotExist: &product.Id,
		Phrase:            request.Phrase,
	}

	units, err := u.repositoryManager.UnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return units, total
}

func (u *unitUseCase) OptionForProductUnitToUnitForm(ctx context.Context, request dto_request.UnitOptionForProductUnitToUnitFormRequest) ([]model.Unit, int) {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	queryOption := model.UnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts{
				{Field: "name", Direction: "asc"},
			},
		),
		ProductIdExist: &product.Id,
		Phrase:         request.Phrase,
	}

	units, err := u.repositoryManager.UnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.UnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return units, total
}
