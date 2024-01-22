package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type SupplierTypeUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.SupplierTypeCreateRequest) model.SupplierType

	// read
	Fetch(ctx context.Context, request dto_request.SupplierTypeFetchRequest) ([]model.SupplierType, int)
	Get(ctx context.Context, request dto_request.SupplierTypeGetRequest) model.SupplierType

	// update
	Update(ctx context.Context, request dto_request.SupplierTypeUpdateRequest) model.SupplierType

	// delete
	Delete(ctx context.Context, request dto_request.SupplierTypeDeleteRequest)
}

type supplierTypeUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewSupplierTypeUseCase(
	repositoryManager repository.RepositoryManager,
) SupplierTypeUseCase {
	return &supplierTypeUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *supplierTypeUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.SupplierTypeRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("SUPPLIER_TYPE.NAME.ALREADY_EXIST"))
	}
}

func (u *supplierTypeUseCase) mustValidateAllowDeleteSupplierType(ctx context.Context, supplierTypeId string) {
	isExist, err := u.repositoryManager.SupplierRepository().IsExistBySupplierTypeId(ctx, supplierTypeId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("SUPPLIER_TYPE.IN_USED_BY.SUPPLIER"))
	}
}

func (u *supplierTypeUseCase) Create(ctx context.Context, request dto_request.SupplierTypeCreateRequest) model.SupplierType {
	u.mustValidateNameNotDuplicate(ctx, request.Name)

	supplierType := model.SupplierType{
		Id:          util.NewUuid(),
		Name:        request.Name,
		Description: request.Description,
	}

	panicIfErr(
		u.repositoryManager.SupplierTypeRepository().Insert(ctx, &supplierType),
	)

	return supplierType
}

func (u *supplierTypeUseCase) Fetch(ctx context.Context, request dto_request.SupplierTypeFetchRequest) ([]model.SupplierType, int) {
	queryOption := model.SupplierTypeQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	supplierTypes, err := u.repositoryManager.SupplierTypeRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.SupplierTypeRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return supplierTypes, total
}

func (u *supplierTypeUseCase) Get(ctx context.Context, request dto_request.SupplierTypeGetRequest) model.SupplierType {
	supplierType := mustGetSupplierType(ctx, u.repositoryManager, request.SupplierTypeId, true)

	return supplierType
}

func (u *supplierTypeUseCase) Update(ctx context.Context, request dto_request.SupplierTypeUpdateRequest) model.SupplierType {
	supplierType := mustGetSupplierType(ctx, u.repositoryManager, request.SupplierTypeId, true)

	if supplierType.Name != request.Name {
		u.mustValidateNameNotDuplicate(ctx, request.Name)
	}

	supplierType.Name = request.Name
	supplierType.Description = request.Description

	panicIfErr(
		u.repositoryManager.SupplierTypeRepository().Update(ctx, &supplierType),
	)

	return supplierType
}

func (u *supplierTypeUseCase) Delete(ctx context.Context, request dto_request.SupplierTypeDeleteRequest) {
	supplierType := mustGetSupplierType(ctx, u.repositoryManager, request.SupplierTypeId, true)

	u.mustValidateAllowDeleteSupplierType(ctx, request.SupplierTypeId)

	panicIfErr(
		u.repositoryManager.SupplierTypeRepository().Delete(ctx, &supplierType),
	)
}
