package use_case

import (
	"context"
	"fmt"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"

	gotiktok "github.com/david-yappeter/go-tiktok"
	"golang.org/x/sync/errgroup"
)

type TiktokProductUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.TiktokProductCreateRequest) model.TiktokProduct

	UploadImage(ctx context.Context, request dto_request.TiktokProductUploadImageRequest) (string, string)
}

type tiktokProductUseCase struct {
	repositoryManager repository.RepositoryManager
}

func NewTiktokProductUseCase(
	repositoryManager repository.RepositoryManager,
) TiktokProductUseCase {
	return &tiktokProductUseCase{
		repositoryManager: repositoryManager,
	}
}

func (u *tiktokProductUseCase) Create(ctx context.Context, request dto_request.TiktokProductCreateRequest) model.TiktokProduct {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	if !product.IsActive {
		panic("TIKTOK_PRODUCT.PRODUCT_MUST_BE_ACTIVE")
	}

	// check tiktok product already created
	isExist, err := u.repositoryManager.TiktokProductRepository().IsExistByProductId(ctx, request.ProductId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("TIKTOK_PRODUCT.ALREADY_EXIST"))
	}

	// check product have stock
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			group.Go(productStockLoader.ProductFn(&product))
		}),
	)

	if product.ProductStock == nil {
		panic("TIKTOK_PRODUCT.PRODUCT_MUST_HAVE_STOCK")
	}

	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	// images
	uriImages := []gotiktok.CreateProductRequestMainImage{}
	for _, uri := range request.ImagesUri {
		uriImages = append(uriImages, gotiktok.CreateProductRequestMainImage{
			Uri: uri,
		})
	}

	// package dimension
	var packageDimension *gotiktok.PackageDimensions = nil
	if request.DimensionUnit != nil {
		packageDimension = &gotiktok.PackageDimensions{
			Height: fmt.Sprintf("%+v", *request.DimensionHeight),
			Length: fmt.Sprintf("%+v", *request.DimensionLength),
			Width:  fmt.Sprintf("%+v", *request.DimensionWidth),
			Unit:   request.DimensionUnit.String(),
		}
	}

	resp, err := client.CreateProduct(ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.CreateProductRequest{
			SaveMode:    nil,
			Description: request.Description,
			CategoryId:  request.CategoryId,
			BrandId:     request.BrandId,
			MainImages:  uriImages,
			Skus: []gotiktok.CreateProductRequestSku{
				{
					Inventory: []gotiktok.CreateProductRequestSkuInventory{
						{
							WarehouseId: tiktokConfig.WarehouseId,
							Quantity:    int(product.ProductStock.Qty),
						},
					},
					SellerSku: &product.Id,
					Price: gotiktok.CreateProductRequestSkuPrice{
						Amount:   fmt.Sprintf("%+v", *product.Price),
						Currency: "IDR",
					},
				},
			},
			Title:             request.Title,
			IsCodAllowed:      false,
			PackageDimensions: packageDimension,
			ProductAttributes: request.Attributes,
			PackageWeight: gotiktok.PackageWeight{
				Unit:  request.WeightUnit.String(),
				Value: fmt.Sprintf("%+v", request.Weight),
			},
			Video:     nil,
			SizeChart: nil,
		},
	)
	panicIfErr(err)

	tiktokProduct := model.TiktokProduct{
		TiktokProductId: resp.ProductId,
		ProductId:       request.ProductId,
		Status:          data_type.TiktokProductStatusActive,
	}

	panicIfErr(
		u.repositoryManager.TiktokProductRepository().Insert(ctx, &tiktokProduct),
	)

	return tiktokProduct
}

func (u *tiktokProductUseCase) UploadImage(ctx context.Context, request dto_request.TiktokProductUploadImageRequest) (string, string) {
	file, err := request.File.Open()
	panicIfErr(err)

	defer file.Close()

	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.UploadImage(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.UploadProductImageRequest{
			Data:    file,
			UseCase: util.StringP("MAIN_IMAGE"),
		},
	)

	panicIfErr(err)

	fmt.Printf("RESP: %+v\n\n", resp)

	return resp.Url, resp.Uri
}
