package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type CustomerUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.CustomerCreateRequest) model.Customer

	//  read
	Fetch(ctx context.Context, request dto_request.CustomerFetchRequest) ([]model.Customer, int)
	Get(ctx context.Context, request dto_request.CustomerGetRequest) model.Customer

	//  update
	Update(ctx context.Context, request dto_request.CustomerUpdateRequest) model.Customer

	//  delete
	Delete(ctx context.Context, request dto_request.CustomerDeleteRequest)

	// option
	OptionForDeliveryOrderForm(ctx context.Context, request dto_request.CustomerOptionForDeliveryOrderFormRequest) ([]model.Customer, int)
}

type customerUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewCustomerUseCase(
	repositoryManager repository.RepositoryManager,
) CustomerUseCase {
	return &customerUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *customerUseCase) mustValidateAllowDeleteCustomer(ctx context.Context, customerId string) {

}

func (u *customerUseCase) Create(ctx context.Context, request dto_request.CustomerCreateRequest) model.Customer {
	customer := model.Customer{
		Id:       util.NewUuid(),
		Name:     request.Name,
		Email:    request.Email,
		Address:  request.Address,
		Phone:    request.Phone,
		IsActive: request.IsActive,
	}

	panicIfErr(
		u.repositoryManager.CustomerRepository().Insert(ctx, &customer),
	)

	return customer
}

func (u *customerUseCase) Fetch(ctx context.Context, request dto_request.CustomerFetchRequest) ([]model.Customer, int) {
	queryOption := model.CustomerQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		IsActive: request.IsActive,
		Phrase:   request.Phrase,
	}

	customers, err := u.repositoryManager.CustomerRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CustomerRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return customers, total
}

func (u *customerUseCase) Get(ctx context.Context, request dto_request.CustomerGetRequest) model.Customer {
	customer := mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	return customer
}

func (u *customerUseCase) Update(ctx context.Context, request dto_request.CustomerUpdateRequest) model.Customer {
	customer := mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	customer.Name = request.Name
	customer.Email = request.Email
	customer.Address = request.Address
	customer.Phone = request.Phone
	customer.IsActive = request.IsActive

	panicIfErr(
		u.repositoryManager.CustomerRepository().Update(ctx, &customer),
	)

	return customer
}

func (u *customerUseCase) Delete(ctx context.Context, request dto_request.CustomerDeleteRequest) {
	customer := mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	u.mustValidateAllowDeleteCustomer(ctx, request.CustomerId)

	panicIfErr(
		u.repositoryManager.CustomerRepository().Delete(ctx, &customer),
	)
}

func (u *customerUseCase) OptionForDeliveryOrderForm(ctx context.Context, request dto_request.CustomerOptionForDeliveryOrderFormRequest) ([]model.Customer, int) {
	queryOption := model.CustomerQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		IsActive: util.BoolP(true),
		Phrase:   request.Phrase,
	}

	customers, err := u.repositoryManager.CustomerRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.CustomerRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return customers, total
}
