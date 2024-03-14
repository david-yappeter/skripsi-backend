package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type CustomerTypeUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.CustomerTypeCreateRequest) model.CustomerType

	//  read
	Fetch(ctx context.Context, request dto_request.CustomerTypeFetchRequest) ([]model.CustomerType, int)
	Get(ctx context.Context, request dto_request.CustomerTypeGetRequest) model.CustomerType

	//  update
	Update(ctx context.Context, request dto_request.CustomerTypeUpdateRequest) model.CustomerType

	//  delete
	Delete(ctx context.Context, request dto_request.CustomerTypeDeleteRequest)
}

type customerTypeUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewCustomerTypeUseCase(
	repositoryManager repository.RepositoryManager,
) CustomerTypeUseCase {
	return &customerTypeUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *customerTypeUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.CustomerTypeRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("CUSTOMER_TYPE.NAME.ALREADY_EXIST"))
	}
}

func (u *customerTypeUseCase) mustValidateAllowDeleteCustomerType(ctx context.Context, customerTypeId string) {
	isExist, err := u.repositoryManager.CustomerRepository().IsExistByCustomerTypeId(ctx, customerTypeId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("CUSTOMER_TYPE.IN_USED_BY_CUSTOMERS"))
	}
}

func (u *customerTypeUseCase) Create(ctx context.Context, request dto_request.CustomerTypeCreateRequest) model.CustomerType {
	u.mustValidateNameNotDuplicate(ctx, request.Name)

	customerType := model.CustomerType{
		Id:          util.NewUuid(),
		Name:        request.Name,
		Description: request.Description,
	}

	panicIfErr(
		u.repositoryManager.CustomerTypeRepository().Insert(ctx, &customerType),
	)

	return customerType
}

func (u *customerTypeUseCase) Fetch(ctx context.Context, request dto_request.CustomerTypeFetchRequest) ([]model.CustomerType, int) {
	queryOption := model.CustomerTypeQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	customerTypes, err := u.repositoryManager.CustomerTypeRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CustomerTypeRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return customerTypes, total
}

func (u *customerTypeUseCase) Get(ctx context.Context, request dto_request.CustomerTypeGetRequest) model.CustomerType {
	customerType := mustGetCustomerType(ctx, u.repositoryManager, request.CustomerTypeId, true)

	return customerType
}

func (u *customerTypeUseCase) Update(ctx context.Context, request dto_request.CustomerTypeUpdateRequest) model.CustomerType {
	customerType := mustGetCustomerType(ctx, u.repositoryManager, request.CustomerTypeId, true)

	if customerType.Name != request.Name {
		u.mustValidateNameNotDuplicate(ctx, request.Name)
	}

	customerType.Name = request.Name
	customerType.Description = request.Description

	panicIfErr(
		u.repositoryManager.CustomerTypeRepository().Update(ctx, &customerType),
	)

	return customerType
}

func (u *customerTypeUseCase) Delete(ctx context.Context, request dto_request.CustomerTypeDeleteRequest) {
	customerType := mustGetCustomerType(ctx, u.repositoryManager, request.CustomerTypeId, true)

	u.mustValidateAllowDeleteCustomerType(ctx, request.CustomerTypeId)

	panicIfErr(
		u.repositoryManager.CustomerTypeRepository().Delete(ctx, &customerType),
	)
}
