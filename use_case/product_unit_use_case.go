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

type productUnitLoaderParams struct {
	product bool
	unit    bool
	toUnit  bool
}

type ProductUnitUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.ProductUnitCreateRequest) model.ProductUnit

	//  read
	Get(ctx context.Context, request dto_request.ProductUnitGetRequest) model.ProductUnit

	//  delete
	Delete(ctx context.Context, request dto_request.ProductUnitDeleteRequest)

	// option
	OptionForProductReceiveItemForm(ctx context.Context, request dto_request.ProductUnitOptionForProductReceiveItemFormRequest) ([]model.ProductUnit, int)
	OptionForDeliveryOrderItemForm(ctx context.Context, request dto_request.ProductUnitOptionForDeliveryOrderItemFormRequest) ([]model.ProductUnit, int)
}

type productUnitUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewProductUnitUseCase(
	repositoryManager repository.RepositoryManager,
) ProductUnitUseCase {
	return &productUnitUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *productUnitUseCase) mustLoadProductUnitsData(ctx context.Context, productUnits []*model.ProductUnit, option productUnitLoaderParams) {
	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())

	panicIfErr(util.Await(func(group *errgroup.Group) {
		for i := range productUnits {
			if option.product {
				group.Go(productLoader.ProductUnitFn(productUnits[i]))
			}

			if option.unit {
				group.Go(unitLoader.ProductUnitFn(productUnits[i]))
			}

			if option.toUnit {
				group.Go(unitLoader.ProductUnitToUnitIdFn(productUnits[i]))
			}
		}
	}))
}

func (u *productUnitUseCase) mustValidateProductUnitNotDuplicate(ctx context.Context, productId string, unitId string) {
	isExist, err := u.repositoryManager.ProductUnitRepository().IsExistByProductIdAndUnitId(ctx, productId, unitId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_UNIT.ALREADY_EXIST"))
	}
}

func (u *productUnitUseCase) mustValidateAllowDeleteProductUnit(ctx context.Context, productUnit model.ProductUnit) {
	isExist, err := u.repositoryManager.ProductUnitRepository().IsExistByProductIdAndToUnitId(ctx, productUnit.ProductId, productUnit.UnitId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_UNIT.IN_USED_BY_OTHER_UNIT"))
	}
}

func (u *productUnitUseCase) Create(ctx context.Context, request dto_request.ProductUnitCreateRequest) model.ProductUnit {
	mustGetUnit(ctx, u.repositoryManager, request.UnitId, true)
	mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	u.mustValidateProductUnitNotDuplicate(ctx, request.ProductId, request.UnitId)

	productUnit := model.ProductUnit{
		Id:          util.NewUuid(),
		ToUnitId:    request.ToUnitId,
		UnitId:      request.UnitId,
		ProductId:   request.ProductId,
		Scale:       request.Scale,
		ScaleToBase: request.Scale,
	}

	if request.ToUnitId != nil {
		toProductUnit := mustGetProductUnitByProductIdAndUnitId(ctx, u.repositoryManager, request.ProductId, *request.ToUnitId, true)

		productUnit.ScaleToBase *= toProductUnit.ScaleToBase
	}

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productUnitRepository := u.repositoryManager.ProductUnitRepository()

			err := productUnitRepository.Insert(ctx, &productUnit)
			if err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadProductUnitsData(ctx, []*model.ProductUnit{&productUnit}, productUnitLoaderParams{
		product: true,
		unit:    true,
		toUnit:  true,
	})

	return productUnit
}

func (u *productUnitUseCase) Get(ctx context.Context, request dto_request.ProductUnitGetRequest) model.ProductUnit {
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	u.mustLoadProductUnitsData(ctx, []*model.ProductUnit{&productUnit}, productUnitLoaderParams{
		product: true,
		unit:    true,
		toUnit:  true,
	})

	return productUnit
}

func (u *productUnitUseCase) Delete(ctx context.Context, request dto_request.ProductUnitDeleteRequest) {
	productUnit := mustGetProductUnit(ctx, u.repositoryManager, request.ProductUnitId, true)

	u.mustValidateAllowDeleteProductUnit(ctx, productUnit)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productUnitRepository := u.repositoryManager.ProductUnitRepository()

			err := productUnitRepository.Delete(ctx, &productUnit)
			if err != nil {
				return err
			}

			return nil
		}),
	)

	u.mustLoadProductUnitsData(ctx, []*model.ProductUnit{&productUnit}, productUnitLoaderParams{
		product: true,
		unit:    true,
		toUnit:  true,
	})
}

func (u *productUnitUseCase) OptionForProductReceiveItemForm(ctx context.Context, request dto_request.ProductUnitOptionForProductReceiveItemFormRequest) ([]model.ProductUnit, int) {
	queryOption := model.ProductUnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		ProductId: &request.ProductId,
		Phrase:    request.Phrase,
	}

	productUnits, err := u.repositoryManager.ProductUnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductUnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductUnitsData(ctx, util.SliceValueToSlicePointer(productUnits), productUnitLoaderParams{
		product: true,
		unit:    true,
		toUnit:  true,
	})

	return productUnits, total
}

func (u *productUnitUseCase) OptionForDeliveryOrderItemForm(ctx context.Context, request dto_request.ProductUnitOptionForDeliveryOrderItemFormRequest) ([]model.ProductUnit, int) {
	mustGetDeliveryOrder(ctx, u.repositoryManager, request.DeliveryOrderId, true)

	deliveryOrderItems, err := u.repositoryManager.DeliveryOrderItemRepository().FetchByDeliveryOrderIds(ctx, []string{request.DeliveryOrderId})
	panicIfErr(err)

	excludeProductUnitIds := []string{}
	for _, deliveryOrderItem := range deliveryOrderItems {
		excludeProductUnitIds = append(excludeProductUnitIds, deliveryOrderItem.ProductUnitId)
	}

	queryOption := model.ProductUnitQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		ExcludeIds: excludeProductUnitIds,
		Phrase:     request.Phrase,
	}

	productUnits, err := u.repositoryManager.ProductUnitRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductUnitRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductUnitsData(ctx, util.SliceValueToSlicePointer(productUnits), productUnitLoaderParams{
		product: true,
		unit:    true,
		toUnit:  true,
	})

	return productUnits, total
}
