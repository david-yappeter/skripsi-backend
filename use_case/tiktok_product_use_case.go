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
	"strconv"

	gotiktok "github.com/david-yappeter/go-tiktok"
	"golang.org/x/sync/errgroup"
)

type TiktokProductUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.TiktokProductCreateRequest) model.TiktokProduct
	UploadImage(ctx context.Context, request dto_request.TiktokProductUploadImageRequest) (string, string)

	// read
	FetchBrands(ctx context.Context, request dto_request.TiktokProductFetchBrandsRequest) (brandList []model.TiktokBrand, nextPageToken string, totalCount int)
	FetchCategories(ctx context.Context) []model.TiktokCategory
	GetCategoryRules(ctx context.Context, request dto_request.TiktokProductGetCategoryRulesRequest) model.TiktokCategoryRule
	GetCategoryAttributes(ctx context.Context, request dto_request.TiktokProductGetCategoryAttributesRequest) []model.TiktokAttribute
	Get(ctx context.Context, request dto_request.TiktokProductGetRequest) model.TiktokPlatformProduct
	RecommendedCategory(ctx context.Context, request dto_request.TiktokProductRecommendedCategoryRequest) model.TiktokCategory

	// update
	Activate(ctx context.Context, request dto_request.TiktokProductActivateRequest)
	Deactivate(ctx context.Context, request dto_request.TiktokProductDeactivateRequest)
	Update(ctx context.Context, request dto_request.TiktokProductUpdateRequest) model.TiktokProduct
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

	// size chart
	var sizeChart *gotiktok.CreateProductRequestSizeChart = nil
	if request.SizeChartUri != nil {
		sizeChart = &gotiktok.CreateProductRequestSizeChart{
			Image: &gotiktok.CreateProductRequestSizeChartImage{
				Uri: *request.SizeChartUri,
			},
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
			SizeChart: sizeChart,
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

	return resp.Url, resp.Uri
}

func (u *tiktokProductUseCase) FetchBrands(ctx context.Context, request dto_request.TiktokProductFetchBrandsRequest) ([]model.TiktokBrand, string, int) {
	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.GetBrands(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.CursorPaginationParam{
			NextPageToken: request.NextPageToken,
			PageSize:      50,
			SortField:     nil,
			SortOrder:     nil,
		},
		gotiktok.GetBrandsRequest{
			CategoryId:   request.CategoryId,
			IsAuthorized: nil,
			BrandName:    request.Phrase,
		},
	)
	panicIfErr(err)

	tiktokBrandList := []model.TiktokBrand{}
	removedBrandCount := 0

	for _, brand := range resp.Brands {
		if brand.BrandStatus != "AVAILABLE" {
			removedBrandCount++
			continue
		}

		tiktokBrandList = append(tiktokBrandList, model.TiktokBrand{
			Id:   brand.Id,
			Name: brand.Name,
		})
	}

	return tiktokBrandList, resp.NextPageToken, resp.TotalCount
}

func (u *tiktokProductUseCase) FetchCategories(ctx context.Context) []model.TiktokCategory {
	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.GetCategories(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.GetCategoriesRequest{
			Locale: util.StringP("id-ID"),
		},
	)
	panicIfErr(err)

	categoryMap := map[string]*model.TiktokCategory{}
	topCategoryList := []*model.TiktokCategory{}

	for _, category := range resp.Categories {
		if category.PermissionStatuses[0] != "AVAILABLE" {
			continue
		}

		var parentId *string = nil
		if category.ParentId != "0" {
			parentId = &category.ParentId
		}

		tiktokCategory := model.TiktokCategory{
			Id:                 category.Id,
			ParentId:           parentId,
			Name:               category.LocalName,
			IsLeaf:             category.IsLeaf,
			ChildrenCategories: []*model.TiktokCategory{},
			ParentCategory:     nil,
		}

		if parentId != nil && categoryMap[*parentId] != nil {
			tiktokCategory.ParentCategory = categoryMap[*parentId]

			categoryMap[*parentId].ChildrenCategories = append(categoryMap[*parentId].ChildrenCategories, &tiktokCategory)
		}

		categoryMap[tiktokCategory.Id] = &tiktokCategory
		if category.ParentId == "0" {
			topCategoryList = append(topCategoryList, &tiktokCategory)
		}
	}

	return util.SlicePointerToSliceValue(topCategoryList)
}

func (u *tiktokProductUseCase) GetCategoryRules(ctx context.Context, request dto_request.TiktokProductGetCategoryRulesRequest) model.TiktokCategoryRule {
	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.GetCategoryRules(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		request.CategoryId,
	)
	panicIfErr(err)

	return model.TiktokCategoryRule{
		PackageDimensionIsRequired: resp.PackageDimension.IsRequired,
		SizeChartIsSupported:       resp.SizeChart.IsSupported,
		SizeChartIsRequired:        resp.SizeChart.IsRequired,
	}
}

func (u *tiktokProductUseCase) GetCategoryAttributes(ctx context.Context, request dto_request.TiktokProductGetCategoryAttributesRequest) []model.TiktokAttribute {
	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.GetAttributes(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		request.CategoryId,
		gotiktok.GetAttributesRequest{
			Locale: util.StringP("id-ID"),
		},
	)
	panicIfErr(err)

	tiktokAttributes := []model.TiktokAttribute{}
	for _, attribute := range resp.Attributes {
		tiktokAttributeValues := []model.TiktokAttributeValue{}

		for _, value := range attribute.Values {
			tiktokAttributeValues = append(tiktokAttributeValues, model.TiktokAttributeValue{
				Id:   value.Id,
				Name: value.Name,
			})
		}

		tiktokAttributes = append(tiktokAttributes, model.TiktokAttribute{
			Id:                  attribute.Id,
			Name:                attribute.Name,
			IsCustomizable:      attribute.IsCustomizable,
			IsMultipleSelection: attribute.IsMultipleSelection,
			Values:              tiktokAttributeValues,
		})
	}

	return tiktokAttributes
}

func (u *tiktokProductUseCase) Get(ctx context.Context, request dto_request.TiktokProductGetRequest) model.TiktokPlatformProduct {
	tiktokProduct, err := u.repositoryManager.TiktokProductRepository().GetByProductId(ctx, request.ProductId)
	panicIfErr(err)

	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	resp, err := client.GetProductDetail(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		tiktokProduct.TiktokProductId,
	)
	panicIfErr(err)

	var (
		tiktokBrand                                      *model.TiktokBrand
		topCategory                                      *model.TiktokCategory
		dimensionHeight, dimensionLength, dimensionWidth *int
		dimensionUnit                                    data_type.TiktokProductDimensionUnit
		weightValue                                      *float64
		weightUnit                                       data_type.TiktokProductPackageWeight
		platformImages                                   []model.TiktokPlatformImage
		platformAttributes                               []model.TiktokPlatformAttribute
	)

	var previousCategory *model.TiktokCategory = nil
	for _, category := range resp.CategoryChains {

		tiktokCategory := model.TiktokCategory{
			Id:                 category.Id,
			Name:               category.LocalName,
			IsLeaf:             category.IsLeaf,
			ParentCategory:     nil,
			ChildrenCategories: []*model.TiktokCategory{},
		}

		if previousCategory != nil {
			previousCategory.ChildrenCategories = append(previousCategory.ChildrenCategories, &tiktokCategory)
		}
		previousCategory = &tiktokCategory

		if topCategory == nil {
			topCategory = previousCategory
		}
	}

	// brand
	if resp.Brand != nil {
		tiktokBrand = &model.TiktokBrand{
			Id:   resp.Brand.Id,
			Name: resp.Brand.Name,
		}
	}

	// dimension
	if resp.PackageDimensions.Height != "0" {
		v, _ := strconv.Atoi(resp.PackageDimensions.Height)
		dimensionHeight = util.IntP(v)
	}

	if resp.PackageDimensions.Width != "0" {
		v, _ := strconv.Atoi(resp.PackageDimensions.Width)
		dimensionWidth = util.IntP(v)
	}

	if resp.PackageDimensions.Length != "0" {
		v, _ := strconv.Atoi(resp.PackageDimensions.Length)
		dimensionLength = util.IntP(v)
	}

	dimensionUnit.Determine(resp.PackageDimensions.Unit)

	// weight
	weightUnit.Determine(resp.PackageWeight.Unit)

	func() {
		v, _ := strconv.ParseFloat(resp.PackageWeight.Value, 64)
		weightValue = util.Float64P(v)
	}()

	// images
	for _, image := range resp.MainImages {
		platformImages = append(platformImages, model.TiktokPlatformImage{
			Height:   image.Height,
			Width:    image.Width,
			ThumbUrl: image.ThumbUrls[0],
			Uri:      image.Uri,
			Url:      image.Urls[0],
		})
	}

	// attributes
	for _, attribute := range resp.ProductAttributes {
		values := []model.TiktokAttributeValue{}
		for _, val := range attribute.Values {
			values = append(values, model.TiktokAttributeValue{
				Id:   val.Id,
				Name: val.Name,
			})
		}

		platformAttributes = append(platformAttributes, model.TiktokPlatformAttribute{
			Id:     attribute.Id,
			Name:   attribute.Name,
			Values: values,
		})
	}

	tiktokPlatformProduct := model.TiktokPlatformProduct{
		Id:              resp.Id,
		Status:          resp.Status,
		Title:           resp.Title,
		Description:     resp.Description,
		Category:        *topCategory,
		Brand:           tiktokBrand,
		DimensionHeight: dimensionHeight,
		DimensionWidth:  dimensionWidth,
		DimensionLength: dimensionLength,
		DimensionUnit:   &dimensionUnit,
		WeightValue:     weightValue,
		WeightUnit:      &weightUnit,
		Images:          platformImages,
		Attributes:      platformAttributes,
	}

	return tiktokPlatformProduct
}

func (u *tiktokProductUseCase) RecommendedCategory(ctx context.Context, request dto_request.TiktokProductRecommendedCategoryRequest) model.TiktokCategory {
	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	var productImages []gotiktok.GetRecommendCategoryRequestImage = nil

	if len(request.ImagesUri) > 0 {
		productImages = []gotiktok.GetRecommendCategoryRequestImage{}
		for _, imageUri := range request.ImagesUri {
			productImages = append(productImages, gotiktok.GetRecommendCategoryRequestImage{
				Uri: imageUri,
			})
		}
	}

	resp, err := client.GetRecommendCategory(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.GetRecommendCategoryRequest{
			ProductTitle: request.ProductTitle,
			Description:  request.Description,
			Images:       productImages,
		},
	)
	panicIfErr(err)

	var topCategory *model.TiktokCategory = nil

	var previousCategory *model.TiktokCategory = nil
	for _, category := range resp.Categories {

		tiktokCategory := model.TiktokCategory{
			Id:                 category.Id,
			Name:               category.Name,
			IsLeaf:             category.IsLeaf,
			ParentCategory:     nil,
			ChildrenCategories: []*model.TiktokCategory{},
		}

		if previousCategory != nil {
			previousCategory.ChildrenCategories = append(previousCategory.ChildrenCategories, &tiktokCategory)
		}
		previousCategory = &tiktokCategory

		if topCategory == nil {
			topCategory = previousCategory
		}
	}

	return *topCategory
}

func (u *tiktokProductUseCase) Activate(ctx context.Context, request dto_request.TiktokProductActivateRequest) {
	tiktokProduct, err := u.repositoryManager.TiktokProductRepository().GetByProductId(ctx, request.ProductId)
	panicIfErr(err)

	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	_, err = client.ActivateProduct(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.ActivateProductRequest{
			ProductIds: []string{tiktokProduct.TiktokProductId},
		},
	)
	panicIfErr(err)

	tiktokProduct.Status = data_type.TiktokProductStatusActive
	panicIfErr(
		u.repositoryManager.TiktokProductRepository().Update(ctx, tiktokProduct),
	)
}

func (u *tiktokProductUseCase) Deactivate(ctx context.Context, request dto_request.TiktokProductDeactivateRequest) {
	tiktokProduct, err := u.repositoryManager.TiktokProductRepository().GetByProductId(ctx, request.ProductId)
	panicIfErr(err)

	client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

	if tiktokConfig.AccessToken == nil {
		panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
	}

	_, err = client.DeactivateProduct(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		gotiktok.DeactivateProductRequest{
			ProductIds: []string{tiktokProduct.TiktokProductId},
		},
	)
	panicIfErr(err)

	tiktokProduct.Status = data_type.TiktokProductStatusInActive
	panicIfErr(
		u.repositoryManager.TiktokProductRepository().Update(ctx, tiktokProduct),
	)
}

func (u *tiktokProductUseCase) Update(ctx context.Context, request dto_request.TiktokProductUpdateRequest) model.TiktokProduct {
	tiktokProduct, err := u.repositoryManager.TiktokProductRepository().GetByProductId(ctx, request.ProductId)
	panicIfErr(err)

	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	if !product.IsActive {
		panic("TIKTOK_PRODUCT.PRODUCT_MUST_BE_ACTIVE")
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
	uriImages := []gotiktok.UpdateProductRequestMainImage{}
	for _, uri := range request.ImagesUri {
		uriImages = append(uriImages, gotiktok.UpdateProductRequestMainImage{
			Uri: uri,
		})
	}

	// package dimension
	var packageDimension *gotiktok.UpdateProductRequestPackageDimension = nil
	if request.DimensionUnit != nil {
		packageDimension = &gotiktok.UpdateProductRequestPackageDimension{
			Height: fmt.Sprintf("%+v", *request.DimensionHeight),
			Length: fmt.Sprintf("%+v", *request.DimensionLength),
			Width:  fmt.Sprintf("%+v", *request.DimensionWidth),
			Unit:   request.DimensionUnit.String(),
		}
	}

	// size chart
	var sizeChart *gotiktok.UpdateProductRequestSizeChart = nil
	if request.SizeChartUri != nil {
		sizeChart = &gotiktok.UpdateProductRequestSizeChart{
			Image: &gotiktok.UpdateProductRequestSizeChartImage{
				Uri: *request.SizeChartUri,
			},
		}
	}

	_, err = client.UpdateProduct(
		ctx,
		gotiktok.CommonParam{
			AccessToken: *tiktokConfig.AccessToken,
			ShopCipher:  tiktokConfig.ShopCipher,
			ShopId:      tiktokConfig.ShopId,
		},
		tiktokProduct.TiktokProductId,
		gotiktok.UpdateProductRequest{
			Description: request.Description,
			CategoryId:  request.CategoryId,
			BrandId:     request.BrandId,
			MainImages:  uriImages,
			Skus: []gotiktok.UpdateProductRequestSku{
				{
					Inventory: []gotiktok.UpdateProductRequestSkuInventory{
						{
							WarehouseId: tiktokConfig.WarehouseId,
							Quantity:    int(product.ProductStock.Qty),
						},
					},
					SellerSku: &product.Id,
					Price: gotiktok.UpdateProductRequestSkuPrice{
						Amount:   fmt.Sprintf("%+v", *product.Price),
						Currency: "IDR",
					},
				},
			},
			Title:          request.Title,
			IsCodAllowed:   false,
			Certifications: nil,
			PackageWeight: gotiktok.UpdateProductRequestPackageWeight{
				Unit:  request.WeightUnit.String(),
				Value: fmt.Sprintf("%+v", request.Weight),
			},
			ProductAttributes: request.Attributes,
			SizeChart:         sizeChart,
			PackageDimensions: packageDimension,
			ExternalProductId: nil,
			DeliveryOptionIds: nil,
			Video:             nil,
		},
	)
	panicIfErr(err)

	return *tiktokProduct
}
