package use_case

import (
	"context"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
)

type SupplierUseCase interface {
	// admin create
	AdminCreate(ctx context.Context, request dto_request.AdminSupplierCreateRequest) model.Supplier

	// admin read
	AdminFetch(ctx context.Context, request dto_request.AdminSupplierFetchRequest) ([]model.Supplier, int)
	AdminGet(ctx context.Context, request dto_request.AdminSupplierGetRequest) model.Supplier

	// admin update
	AdminUpdate(ctx context.Context, request dto_request.AdminSupplierUpdateRequest) model.Supplier

	// admin delete
	AdminDelete(ctx context.Context, request dto_request.AdminSupplierDeleteRequest)
}

type supplierUseCase struct {
	repositoryManager repository.RepositoryManager
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

func (u *supplierUseCase) AdminCreate(ctx context.Context, request dto_request.AdminSupplierCreateRequest) model.Supplier {
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

	return supplier
}

func (u *supplierUseCase) AdminFetch(ctx context.Context, request dto_request.AdminSupplierFetchRequest) ([]model.Supplier, int) {
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

	return suppliers, total
}

func (u *supplierUseCase) AdminGet(ctx context.Context, request dto_request.AdminSupplierGetRequest) model.Supplier {
	supplier := mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	return supplier
}

func (u *supplierUseCase) AdminUpdate(ctx context.Context, request dto_request.AdminSupplierUpdateRequest) model.Supplier {
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

	return supplier
}

func (u *supplierUseCase) AdminDelete(ctx context.Context, request dto_request.AdminSupplierDeleteRequest) {
	supplier := mustGetSupplier(ctx, u.repositoryManager, request.SupplierId, true)

	u.mustValidateAllowDeleteSupplier(ctx, request.SupplierId)

	panicIfErr(
		u.repositoryManager.SupplierRepository().Delete(ctx, &supplier),
	)
}

func NewSupplierUseCase(
	repositoryManager repository.RepositoryManager,
) SupplierUseCase {
	return &supplierUseCase{
		repositoryManager: repositoryManager,
	}
}
