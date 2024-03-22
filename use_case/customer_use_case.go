package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type customerLoaderParams struct {
	customerType bool
}

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

func (u *customerUseCase) mustLoadCustomersData(ctx context.Context, customers []*model.Customer, option customerLoaderParams) {
	customerTypeLoader := loader.NewCustomerTypeLoader(u.repositoryManager.CustomerTypeRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range customers {
			if option.customerType {
				group.Go(customerTypeLoader.CustomerFn(customers[i]))
			}
		}
	}))
}

func (u *customerUseCase) mustValidateAllowDeleteCustomer(ctx context.Context, customerId string) {

}

func (u *customerUseCase) Create(ctx context.Context, request dto_request.CustomerCreateRequest) model.Customer {
	if request.CustomerTypeId != nil {
		mustGetCustomerType(ctx, u.repositoryManager, *request.CustomerTypeId, true)
	}

	customer := model.Customer{
		Id:             util.NewUuid(),
		CustomerTypeId: request.CustomerTypeId,
		Name:           request.Name,
		Email:          request.Email,
		Address:        request.Address,
		Latitude:       request.Latitude,
		Longitude:      request.Longitude,
		Phone:          request.Phone,
		IsActive:       request.IsActive,
	}

	panicIfErr(
		u.repositoryManager.CustomerRepository().Insert(ctx, &customer),
	)

	u.mustLoadCustomersData(ctx, []*model.Customer{&customer}, customerLoaderParams{
		customerType: true,
	})

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

	u.mustLoadCustomersData(ctx, util.SliceValueToSlicePointer(customers), customerLoaderParams{
		customerType: true,
	})

	return customers, total
}

func (u *customerUseCase) Get(ctx context.Context, request dto_request.CustomerGetRequest) model.Customer {
	customer := mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	u.mustLoadCustomersData(ctx, []*model.Customer{&customer}, customerLoaderParams{
		customerType: true,
	})

	return customer
}

func (u *customerUseCase) Update(ctx context.Context, request dto_request.CustomerUpdateRequest) model.Customer {
	customer := mustGetCustomer(ctx, u.repositoryManager, request.CustomerId, true)

	if customer.CustomerTypeId != nil {
		mustGetCustomerType(ctx, u.repositoryManager, *request.CustomerTypeId, true)
	}

	customer.CustomerTypeId = request.CustomerTypeId
	customer.Name = request.Name
	customer.Email = request.Email
	customer.Address = request.Address
	customer.Latitude = request.Latitude
	customer.Longitude = request.Longitude
	customer.Phone = request.Phone
	customer.IsActive = request.IsActive

	panicIfErr(
		u.repositoryManager.CustomerRepository().Update(ctx, &customer),
	)

	u.mustLoadCustomersData(ctx, []*model.Customer{&customer}, customerLoaderParams{
		customerType: true,
	})

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

	u.mustLoadCustomersData(ctx, util.SliceValueToSlicePointer(customers), customerLoaderParams{
		customerType: true,
	})

	return customers, total
}
