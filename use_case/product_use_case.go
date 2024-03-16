package use_case

import (
	"context"
	"fmt"
	"myapp/constant"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/internal/filesystem"
	"myapp/loader"
	"myapp/model"
	"myapp/repository"
	"myapp/util"
	"path"

	gotiktok "github.com/david-yappeter/go-tiktok"
	"golang.org/x/sync/errgroup"
)

type productLoaderParams struct {
	productStock  bool
	tiktokProduct bool
	productUnits  bool
	productImage  bool
}

type ProductUseCase interface {
	//  create
	Create(ctx context.Context, request dto_request.ProductCreateRequest) model.Product
	Upload(ctx context.Context, request dto_request.ProductUploadRequest) string

	//  read
	Fetch(ctx context.Context, request dto_request.ProductFetchRequest) ([]model.Product, int)
	Get(ctx context.Context, request dto_request.ProductGetRequest) model.Product

	//  update
	Update(ctx context.Context, request dto_request.ProductUpdateRequest) model.Product

	//  delete
	Delete(ctx context.Context, request dto_request.ProductDeleteRequest)

	// option
	OptionForProductReceiveForm(ctx context.Context, request dto_request.ProductOptionForProductReceiveFormRequest) ([]model.Product, int)
	OptionForDeliveryOrderForm(ctx context.Context, request dto_request.ProductOptionForDeliveryOrderFormRequest) ([]model.Product, int)
	OptionForCustomerTypeForm(ctx context.Context, request dto_request.ProductOptionForCustomerTypeFormRequest) ([]model.Product, int)
}

type productUseCase struct {
	repositoryManager repository.RepositoryManager

	baseFileUseCase baseFileUseCase
}

func NewProductUseCase(
	repositoryManager repository.RepositoryManager,
	mainFilesystem filesystem.Client,
	tmpFilesystem filesystem.Client,
) ProductUseCase {
	return &productUseCase{
		repositoryManager: repositoryManager,
		baseFileUseCase: newBaseFileUseCase(
			mainFilesystem,
			tmpFilesystem,
		),
	}
}

func (u *productUseCase) mustValidateNameNotDuplicate(ctx context.Context, name string) {
	isExist, err := u.repositoryManager.ProductRepository().IsExistByName(ctx, name)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.NAME.ALREADY_EXIST"))
	}
}

func (u *productUseCase) mustValidateAllowDeleteProduct(ctx context.Context, productId string) {
	isExist, err := u.repositoryManager.ProductStockRepository().IsExistByProductId(ctx, productId)
	panicIfErr(err)

	if isExist {
		panic(dto_response.NewBadRequestErrorResponse("PRODUCT.ALREADY_HAVE_STOCK"))
	}
}

func (u *productUseCase) mustLoadProductDatas(ctx context.Context, products []*model.Product, option productLoaderParams) {
	productStockLoader := loader.NewProductStockLoader(u.repositoryManager.ProductStockRepository())
	tiktokProductLoader := loader.NewTiktokProductLoader(u.repositoryManager.TiktokProductRepository())
	productUnitsLoader := loader.NewProductUnitsLoader(u.repositoryManager.ProductUnitRepository())
	fileLoader := loader.NewFileLoader(u.repositoryManager.FileRepository())

	panicIfErr(
		util.Await(func(group *errgroup.Group) {
			for i := range products {
				if option.productStock {
					group.Go(productStockLoader.ProductFnNotStrict(products[i]))
				}

				if option.tiktokProduct {
					group.Go(tiktokProductLoader.ProductFnNotStrict(products[i]))
				}

				if option.productUnits {
					group.Go(productUnitsLoader.ProductFn(products[i]))
				}

				if option.productImage {
					group.Go(fileLoader.ProductFn(products[i]))
				}
			}
		}),
	)

	for i := range products {
		products[i].ImageFile.SetLink(u.baseFileUseCase.mainFilesystem)
	}
}

func (u *productUseCase) Create(ctx context.Context, request dto_request.ProductCreateRequest) model.Product {
	u.mustValidateNameNotDuplicate(ctx, request.Name)

	product := model.Product{
		Id:          util.NewUuid(),
		ImageFileId: "",
		Name:        request.Name,
		Description: request.Description,
		Price:       nil,
		IsActive:    false,
	}

	// product stock
	productStock := model.ProductStock{
		Id:        util.NewUuid(),
		ProductId: product.Id,
		Qty:       0,
	}

	// upload image file
	imageFile := model.File{
		Id:   util.NewUuid(),
		Type: data_type.FileTypeProductImage,
	}

	product.ImageFileId = imageFile.Id

	imageFile.Path, imageFile.Name = u.baseFileUseCase.mustUploadFileFromTemporaryToMain(
		ctx,
		constant.ProductImagePath,
		product.Id,
		fmt.Sprintf("%s%s", imageFile.Id, path.Ext(request.ImageFilePath)),
		request.ImageFilePath,
		fileUploadTemporaryToMainParams{
			deleteTmpOnSuccess: false,
		},
	)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			fileRepository := u.repositoryManager.FileRepository()
			productRepository := u.repositoryManager.ProductRepository()
			productStockRepository := u.repositoryManager.ProductStockRepository()

			if err := fileRepository.Insert(ctx, &imageFile); err != nil {
				return err
			}

			if err := productRepository.Insert(ctx, &product); err != nil {
				return err
			}

			if err := productStockRepository.Insert(ctx, &productStock); err != nil {
				return err
			}

			return nil
		}),
	)

	return product
}

func (u *productUseCase) Upload(ctx context.Context, request dto_request.ProductUploadRequest) string {
	return u.baseFileUseCase.mustUploadFileToTemporary(
		ctx,
		constant.ProductImagePath,
		request.File.Filename,
		request.File,
		fileUploadTemporaryParams{
			supportedExtensions: listSupportedExtension([]string{
				extensionTypeImage,
			}),
			maxFileSizeInBytes: util.Pointer[int64](2 << 20),
		},
	)
}

func (u *productUseCase) Fetch(ctx context.Context, request dto_request.ProductFetchRequest) ([]model.Product, int) {
	queryOption := model.ProductQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		IsActive: request.IsActive,
		Phrase:   request.Phrase,
	}

	products, err := u.repositoryManager.ProductRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductDatas(ctx, util.SliceValueToSlicePointer(products), productLoaderParams{
		productStock:  true,
		productImage:  true,
		tiktokProduct: true,
	})

	return products, total
}

func (u *productUseCase) Get(ctx context.Context, request dto_request.ProductGetRequest) model.Product {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	u.mustLoadProductDatas(ctx, []*model.Product{&product}, productLoaderParams{
		productStock:  true,
		tiktokProduct: true,
		productImage:  true,
		productUnits:  true,
	})

	return product
}

func (u *productUseCase) Update(ctx context.Context, request dto_request.ProductUpdateRequest) model.Product {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)
	tiktokProduct := shouldGetTiktokProductByProductId(ctx, u.repositoryManager, product.Id)
	isProductUnitsExist, err := u.repositoryManager.ProductUnitRepository().IsExistByProductId(ctx, product.Id)
	panicIfErr(err)

	if request.IsActive && !isProductUnitsExist {
		panic(dto_response.NewBadRequestErrorResponse("ACTIVE_PRODUCT.MUST_HAVE_AN_UNIT"))
	}

	if product.Name != request.Name {
		u.mustValidateNameNotDuplicate(ctx, request.Name)
	}

	if request.IsActive && request.Price == nil {
		panic(dto_response.NewBadRequestErrorResponse("ACTIVE_PRODUCT.MUST_HAVE_PRICE"))
	}

	// update product price in tiktok
	if tiktokProduct != nil && request.Price != nil && (product.Price == nil || *product.Price != *request.Price) {
		tiktokProductDetail := mustGetTiktokProductDetail(ctx, u.repositoryManager, tiktokProduct.TiktokProductId)

		client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

		if tiktokConfig.AccessToken == nil {
			panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
		}

		_, err := client.UpdateProductPrice(
			ctx,
			gotiktok.CommonParam{
				AccessToken: *tiktokConfig.AccessToken,
				ShopCipher:  tiktokConfig.ShopCipher,
				ShopId:      tiktokConfig.ShopId,
			},
			tiktokProduct.TiktokProductId,
			gotiktok.UpdateProductPriceRequest{
				Skus: []gotiktok.UpdateProductPriceRequestSku{
					{
						Id: tiktokProductDetail.Skus[0].Id,
						Price: gotiktok.UpdateProductPriceRequestSkuPrice{
							Amount:   fmt.Sprintf("%f", *request.Price),
							Currency: "IDR",
						},
					},
				},
			},
		)
		panicIfErr(err)
	}

	// deactivate tiktok product
	if tiktokProduct != nil && product.IsActive != request.IsActive && !request.IsActive {
		client, tiktokConfig := mustGetTiktokClient(ctx, u.repositoryManager)

		if tiktokConfig.AccessToken == nil {
			panic("TIKTOK_CONFIG.ACCESS_TOKEN_EMPTY")
		}

		_, err := client.DeactivateProduct(
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
	}

	product.Name = request.Name
	product.Description = request.Description
	product.Price = request.Price
	product.IsActive = request.IsActive

	panicIfErr(
		u.repositoryManager.ProductRepository().Update(ctx, &product),
	)

	u.mustLoadProductDatas(ctx, []*model.Product{&product}, productLoaderParams{
		productStock: true,
		productUnits: true,
	})

	return product
}

func (u *productUseCase) Delete(ctx context.Context, request dto_request.ProductDeleteRequest) {
	product := mustGetProduct(ctx, u.repositoryManager, request.ProductId, true)

	u.mustValidateAllowDeleteProduct(ctx, request.ProductId)

	imageFile := mustGetFile(ctx, u.repositoryManager, product.ImageFileId, true)

	panicIfErr(
		u.repositoryManager.Transaction(ctx, func(ctx context.Context) error {
			productRepository := u.repositoryManager.ProductRepository()
			fileRepository := u.repositoryManager.FileRepository()

			if err := productRepository.Insert(ctx, &product); err != nil {
				return err
			}

			if err := fileRepository.Delete(ctx, &imageFile); err != nil {
				return err
			}

			return nil
		}),
	)

	u.baseFileUseCase.mainFilesystem.Delete(imageFile.Path)
}

func (u *productUseCase) OptionForProductReceiveForm(ctx context.Context, request dto_request.ProductOptionForProductReceiveFormRequest) ([]model.Product, int) {
	queryOption := model.ProductQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		IsActive: util.BoolP(true),
		Phrase:   request.Phrase,
	}

	products, err := u.repositoryManager.ProductRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return products, total
}

func (u *productUseCase) OptionForDeliveryOrderForm(ctx context.Context, request dto_request.ProductOptionForDeliveryOrderFormRequest) ([]model.Product, int) {
	queryOption := model.ProductQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		IsActive: util.BoolP(true),
		Phrase:   request.Phrase,
	}

	products, err := u.repositoryManager.ProductRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductRepository().Count(ctx, queryOption)
	panicIfErr(err)

	u.mustLoadProductDatas(ctx, util.SliceValueToSlicePointer(products), productLoaderParams{
		productStock: true,
	})

	return products, total
}

func (u *productUseCase) OptionForCustomerTypeForm(ctx context.Context, request dto_request.ProductOptionForCustomerTypeFormRequest) ([]model.Product, int) {
	customerTypeDiscounts, err := u.repositoryManager.CustomerTypeDiscountRepository().FetchByCustomerTypeIds(ctx, []string{request.CustomerTypeId})
	panicIfErr(err)

	excludedProductIds := []string{}
	for _, customerTypeDiscount := range customerTypeDiscounts {
		excludedProductIds = append(excludedProductIds, customerTypeDiscount.ProductId)
	}

	queryOption := model.ProductQueryOption{
		QueryOption: model.NewQueryOptionWithPagination(
			request.Page,
			request.Limit,
			model.Sorts(request.Sorts),
		),
		ExcludeIds: excludedProductIds,
		IsActive:   util.BoolP(true),
		Phrase:     request.Phrase,
	}

	products, err := u.repositoryManager.ProductRepository().Fetch(ctx, queryOption)
	panicIfErr(err)

	total, err := u.repositoryManager.ProductRepository().Count(ctx, queryOption)
	panicIfErr(err)

	return products, total
}
