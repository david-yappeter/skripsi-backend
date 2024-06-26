package use_case

import (
	"context"
	"io"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	"golang.org/x/sync/errgroup"
)

type productStockLoaderParams struct {
	product bool
}

type ProductStockUseCase interface {
	//  read
	Fetch(ctx context.Context, request dto_request.ProductStockFetchRequest) ([]model.ProductStock, int)
	Get(ctx context.Context, request dto_request.ProductStockGetRequest) model.ProductStock
	DownloadReport(ctx context.Context) (io.ReadCloser, int64, string, string)

	//  update
	Adjustment(ctx context.Context, request dto_request.ProductStockAdjustmentRequest) model.ProductStock
}

type productStockUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewProductStockUseCase(
	repositoryManager repository.RepositoryManager,
) ProductStockUseCase {
	return &productStockUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *productStockUseCase) mustLoadProductStockDatas(ctx context.Context, productStocks []*model.ProductStock, option productStockLoaderParams) {
	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productStocks {
				group.Go(productLoader.ProductStockFn(productStocks[i]))
			}
		}),
	)
}

func (u *productStockUseCase) Fetch(ctx context.Context, request dto_request.ProductStockFetchRequest) ([]model.ProductStock, int) {
	queryOption := model.ProductStockQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		Phrase: request.Phrase,
	}

	productStocks, err := u.repositoryManager.ProductStockRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductStockRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductStockDatas(ctx, util.SliceValueToSlicePointer(productStocks), productStockLoaderParams{
		product: true,
	})

	return productStocks, total
}

func (u *productStockUseCase) Get(ctx context.Context, request dto_request.ProductStockGetRequest) model.ProductStock {
	productStock := mustGetProductStock(ctx, u.repositoryManager, request.ProductStockId, true)

	u.mustLoadProductStockDatas(ctx, []*model.ProductStock{&productStock}, productStockLoaderParams{
		product: true,
	})

	return productStock
}

func (u *productStockUseCase) DownloadReport(
	ctx context.Context,
) (io.ReadCloser, int64, string, string) {
	productRepository := u.repositoryManager.ProductRepository()
	productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

	products, err := productRepository.Fetch(ctx)
	panicIfErr(err)

	baseProductUnitLoader := loader.NewBaseProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range products {
				group.Go(baseProductUnitLoader.ProductFnNotStrict(&products[i]))
				group.Go(productStockLoader.ProductFnNotStrict(&products[i]))
			}
		}),
	)

	unitLoader := loader.NewUnitLoader(u.repositoryManager.UnitRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range products {
				group.Go(unitLoader.ProductUnitFn(products[i].BaseProductUnit))
			}
		}),
	)

	// construct excel report sheets
	reportExcel, err := NewReportStockExcel(
		util.CurrentDateTime(),
	)
	panicIfErr(err)
	defer reportExcel.Close()

	for _, product := range products {
		baseUnit := "-"
		stockLeft := 0.0
		currentSellingPrice := 0.0

		if product.BaseProductUnit != nil {
			baseUnit = product.BaseProductUnit.Unit.Name
		}

		if product.Price != nil {
			currentSellingPrice = *product.Price
		}

		if product.ProductStock != nil {
			stockLeft = product.ProductStock.Qty
		}

		reportExcel.AddSheet1Data(ReportStockExcelSheet1Data{
			ProductId:           product.Id,
			ProductName:         product.Name,
			BaseUnit:            baseUnit,
			CurrentSellingPrice: currentSellingPrice,
			IsActive:            false,
			StockLeft:           stockLeft,
		})
	}

	productStockMutations, err := productStockMutationRepository.FetchHaveQtyLeft(ctx)
	panicIfErr(err)

	productUnitLoader := loader.NewProductUnitLoader(u.repositoryManager.ProductUnitRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productStockMutations {
				group.Go(productUnitLoader.ProductStockMutationFn(&productStockMutations[i]))
			}
		}),
	)

	productLoader := loader.NewProductLoader(u.repositoryManager.ProductRepository())
	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range productStockMutations {
				group.Go(productLoader.ProductUnitFn(productStockMutations[i].ProductUnit))
				group.Go(unitLoader.ProductUnitFn(productStockMutations[i].ProductUnit))
			}
		}),
	)

	for _, productStockMutation := range productStockMutations {
		reportExcel.AddSheet2Data(ReportStockExcelSheet2Data{
			ProductId:     productStockMutation.ProductUnit.ProductId,
			UnitId:        productStockMutation.ProductUnit.UnitId,
			ProductName:   productStockMutation.ProductUnit.Product.Name,
			UnitName:      productStockMutation.ProductUnit.Unit.Name,
			MutationType:  productStockMutation.Type.String(),
			Qty:           productStockMutation.Qty,
			ScaleToBase:   productStockMutation.ScaleToBase,
			BaseQty:       productStockMutation.ScaleToBase * productStockMutation.Qty,
			BaseQtyLeft:   productStockMutation.BaseQtyLeft,
			BaseQtySold:   (productStockMutation.ScaleToBase * productStockMutation.Qty) - productStockMutation.BaseQtyLeft,
			BaseCostPrice: productStockMutation.BaseCostPrice,
			MutatedAt:     productStockMutation.MutatedAt.Time(),
		})
	}

	readCloser, contentLength, err := reportExcel.ToReadSeekCloserWithContentLength()
	panicIfErr(err)

	return readCloser, contentLength, "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet", "product_stock.xlsx"
}

func (u *productStockUseCase) Adjustment(ctx context.Context, request dto_request.ProductStockAdjustmentRequest) model.ProductStock {
	currentUser := model.MustGetUserCtx(ctx)
	currentDateTime := util.CurrentDateTime()
	productStock := mustGetProductStock(ctx, u.repositoryManager, request.ProductStockId, true)
	previousProductStock := productStock

	product := mustGetProduct(ctx, u.repositoryManager, productStock.ProductId, true)

	if !product.IsActive {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_STOCK.PRODUCT_MUST_BE_ACTIVE"))
	}

	// cost_price is required when adding stock
	if productStock.Qty < request.Qty && request.CostPrice == nil {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT_STOCK.COST_PRICE_IS_REQUIRED_WHEN_ADDING_STOCK"))
	}

	// adjust base cost price
	if request.Qty > productStock.Qty {
		productStock.BaseCostPrice = productStock.RecalculateBaseCostPrice(request.Qty-productStock.Qty, *request.CostPrice)
	}

	var toBeAddedProductStockMutation *model.ProductStockMutation

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productStockAdjustmentRepository := u.repositoryManager.ProductStockAdjustmentRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()
			productStockMutationRepository := u.repositoryManager.ProductStockMutationRepository()

			if productStock.Qty < request.Qty {
				baseProductUnit, err := u.repositoryManager.ProductUnitRepository().GetBaseProductUnitByProductId(ctx, productStock.ProductId)
				if err != nil {
					return err
				}

				toBeAddedProductStockMutation = &model.ProductStockMutation{
					Id:            util.NewUuid(),
					ProductUnitId: baseProductUnit.Id,
					Type:          data_type.ProductStockMutationTypeProductStockAdjustment,
					IdentifierId:  productStock.Id,
					Qty:           request.Qty - productStock.Qty,

					ScaleToBase:   1,
					BaseQtyLeft:   request.Qty - productStock.Qty,
					BaseCostPrice: *request.CostPrice,
					MutatedAt:     currentDateTime,
				}
			} else if productStock.Qty > request.Qty {
				productStock.Qty = request.Qty
				deductQty := request.Qty - productStock.Qty

				for deductQty > 0 {
					productStockMutation, err := u.repositoryManager.ProductStockMutationRepository().GetFIFOByProductIdAndBaseQtyLeftNotZero(ctx, product.Id)
					if err != nil {
						return err
					}

					if deductQty > productStockMutation.BaseQtyLeft {
						deductQty -= productStockMutation.BaseQtyLeft
						productStockMutation.Qty = 0
					} else {
						productStockMutation.Qty -= deductQty
						deductQty = 0
					}

					if err := productStockMutationRepository.Update(ctx, productStockMutation); err != nil {
						return err
					}
				}
			}

			// change qty
			productStock.Qty = request.Qty

			if request.Qty == 0 {
				productStock.BaseCostPrice = 0
			}

			// update tiktok product stock
			tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, product.Id)

			if tiktokProduct != nil {
				mustUpdateTiktokProductInventory(ctx, u.repositoryManager, tiktokProduct.TiktokProductId, int(productStock.Qty))
			}

			if err := productStockRepository.Update(ctx, &productStock); err != nil {
				return err
			}

			if toBeAddedProductStockMutation != nil {
				if err := productStockMutationRepository.Insert(ctx, toBeAddedProductStockMutation); err != nil {
					return err
				}
			}

			if err := productStockAdjustmentRepository.Insert(ctx, &model.ProductStockAdjustment{
				Id:             util.NewUuid(),
				UserId:         currentUser.Id,
				ProductStockId: productStock.Id,
				PreviousQty:    previousProductStock.Qty,
				UpdatedQty:     productStock.Qty,
			}); err != nil {
				return err
			}

			return nil
		}),
	)

	return productStock
}
