package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type SupplierUseCase interface {
	// create
	Create(ctx context.Context, request dto_request.SupplierCreateRequest) model.Supplier

	// read
	Fetch(ctx context.Context, request dto_request.SupplierFetchRequest) ([]model.Supplier, int)
	Get(ctx context.Context, request dto_request.SupplierGetRequest) model.Supplier

	// update
	Update(ctx context.Context, request dto_request.SupplierUpdateRequest) model.Supplier

	// delete
	Delete(ctx context.Context, request dto_request.SupplierDeleteRequest)
}

type supplierUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewSupplierUseCase(
	repositoryManager repository.RepositoryManager,
) SupplierUseCase {
	return &supplierUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *supplierUseCase) mustValidateCodeNotDuplicate(ctx context.Context, code string) {
	isExist, err := u.repositoryManager.SupplierRepository().IsExistByCode(ctx, code)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("SUPPLIER_TYPE.CODE.ALREADY_EXIST"))
	}
}

func (u *supplierUseCase) mustValidateAllowDeleteSupplier(ctx context.Context, supplierId string) {

}

func (u *supplierUseCase) mustLoadSupplierDatas(ctx context.Context, suppliers []*model.Supplier) {
	supplierTypeLoader := loader.NewSupplierTypeLoader(u.repositoryManager.SupplierTypeRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range suppliers {
				supplier := suppliers[i]
				group.Go(supplierTypeLoader.SupplierFn(supplier))
			}
		}),
	)
}

func (u *supplierUseCase) Create(ctx context.Context, request dto_request.SupplierCreateRequest) model.Supplier {
	mustGetSupplierType(ctx, u.repositoryManager, request.SupplierTypeId, true)
	u.mustValidateCodeNotDuplicate(ctx, request.Name)

	supplier := model.Supplier{
		Id:             util.NewUuid(),
		SupplierTypeId: request.SupplierTypeId,
		Code:           request.Code,
		Name:           request.Name,
		IsActive:       request.IsActive,
		Address:        request.Address,
		Phone:          request.Phone,
		Email:          request.Email,
		Description:    request.Description,
	}

	panicIfErr(
		u.repositoryManager.SupplierRepository().Insert(ctx, &supplier),
	)

	u.mustLoadSupplierDatas(ctx, []*model.Supplier{&supplier})

	return supplier
}

func (u *supplierUseCase) Fetch(ctx context.Context, request dto_request.SupplierFetchRequest) ([]model.Supplier, int) {
	queryOption := model.SupplierQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		IsActive:        request.IsActive,
		SupplierTypeIds: util.AppendIfNotNil([]string{}, request.SupplierTypeId),
		Phrase:          request.Phrase,
	}

	suppliers, err := u.repositoryManager.SupplierRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.SupplierRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadSupplierDatas(ctx, util.SliceValueToSlicePointer(suppliers))

	return suppliers, total
}

func (u *supplierUseCase) Get(ctx context.Context, request dto_request.SupplierGetRequest) model.Supplier {
	supplier := mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	u.mustLoadSupplierDatas(ctx, []*model.Supplier{&supplier})

	return supplier
}

func (u *supplierUseCase) Update(ctx context.Context, request dto_request.SupplierUpdateRequest) model.Supplier {
	supplier := mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	if supplier.Name != request.Name {
		u.mustValidateCodeNotDuplicate(ctx, request.Name)
	}

	supplier.SupplierTypeId = request.SupplierTypeId
	supplier.Name = request.Name
	supplier.Address = request.Name
	supplier.IsActive = request.IsActive
	supplier.Address = request.Address
	supplier.Phone = request.Phone
	supplier.Email = request.Email
	supplier.Description = request.Description

	panicIfErr(
		u.repositoryManager.SupplierRepository().Update(ctx, &supplier),
	)

	u.mustLoadSupplierDatas(ctx, []*model.Supplier{&supplier})

	return supplier
}

func (u *supplierUseCase) Delete(ctx context.Context, request dto_request.SupplierDeleteRequest) {
	supplier := mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	u.mustValidateAllowDeleteSupplier(ctx, request.SupplierId)

	panicIfErr(
		u.repositoryManager.SupplierRepository().Delete(ctx, &supplier),
	)
}
