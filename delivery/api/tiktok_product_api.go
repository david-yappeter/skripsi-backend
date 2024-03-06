package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"myapp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TiktokProductApi struct {
	api
	tiktokProductUseCase use_case.TiktokProductUseCase
}

// API:
//
//	@Router		/tiktok-products [post]
//	@Summary	Create
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		dto_request.TiktokProductCreateRequest	body	dto_request.TiktokProductCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *TiktokProductApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductCreate),
		func(ctx apiContext) {
			var request dto_request.TiktokProductCreateRequest
			ctx.mustBind(&request)

			a.tiktokProductUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/upload-image [post]
//	@Summary	Upload Image
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		dto_request.TiktokProductUploadImageRequest	body	dto_request.TiktokProductUploadImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{url=string,uri=string}}
func (a *TiktokProductApi) UploadImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductUploadImage),
		func(ctx apiContext) {
			var request dto_request.TiktokProductUploadImageRequest
			ctx.mustBind(&request)

			url, uri := a.tiktokProductUseCase.UploadImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"url": url,
						"uri": uri,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/brands [post]
//	@Summary	Fetch Brands field for Form
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		dto_request.TiktokProductFetchBrandsRequest	body	dto_request.TiktokProductFetchBrandsRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{brands=[]dto_response.TiktokBrandResponse,next_page_token=string,total_count=int}}
func (a *TiktokProductApi) FetchBrands() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductFetchBrands),
		func(ctx apiContext) {
			var request dto_request.TiktokProductFetchBrandsRequest
			ctx.mustBind(&request)

			brands, nextPageToken, totalCount := a.tiktokProductUseCase.FetchBrands(ctx.context(), request)

			nodes := util.ConvertArray(brands, dto_response.NewTiktokBrandResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"brands":          nodes,
						"next_page_token": nextPageToken,
						"total_count":     totalCount,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/categories [post]
//	@Summary	Fetch Categories field for Form
//	@tags		Tiktok Products
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{categories=[]dto_response.TiktokCategoryResponse}}
func (a *TiktokProductApi) FetchCategories() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductFetchCategories),
		func(ctx apiContext) {
			categories := a.tiktokProductUseCase.FetchCategories(ctx.context())

			nodes := util.ConvertArray(categories, dto_response.NewTiktokCategoryResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"categories": nodes,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/categories/{category_id}/rules [post]
//	@Summary	Get Category Rule
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		category_id	path	string	true	"Tiktok Category Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{category_rule=dto_response.TiktokCategoryRuleResponse}}
func (a *TiktokProductApi) GetCategoryRules() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductGetCategoryRules),
		func(ctx apiContext) {
			categoryId := ctx.getParam("category_id")
			var request dto_request.TiktokProductGetCategoryRulesRequest
			ctx.mustBind(&request)
			request.CategoryId = categoryId

			categoryRule := a.tiktokProductUseCase.GetCategoryRules(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"category_rule": dto_response.NewTiktokCategoryRuleResponse(categoryRule),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/categories/{category_id}/attributes [post]
//	@Summary	Get Category Attributes
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		category_id	path	string	true	"Tiktok Category Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{attributes=[]dto_response.TiktokAttributeResponse}}
func (a *TiktokProductApi) GetCategoryAttributes() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductGetCategoryAttributes),
		func(ctx apiContext) {
			categoryId := ctx.getParam("category_id")
			var request dto_request.TiktokProductGetCategoryAttributesRequest
			request.CategoryId = categoryId

			attributes := a.tiktokProductUseCase.GetCategoryAttributes(ctx.context(), request)

			nodes := util.ConvertArray(attributes, dto_response.NewTiktokAttributeResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"attributes": nodes,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/{id} [get]
//	@Summary	Get
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		id	path	string	true	"Product Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{tiktok_product=dto_response.TiktokPlatformProductResponse}}
func (a *TiktokProductApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductGet),
		func(ctx apiContext) {
			productId := ctx.getUuidParam("id")
			var request dto_request.TiktokProductGetRequest
			request.ProductId = productId

			product := a.tiktokProductUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"tiktok_product": dto_response.NewTiktokPlatformProductResponse(product),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/{id} [put]
//	@Summary	Update
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		id	path	string	true	"Product Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *TiktokProductApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductUpdate),
		func(ctx apiContext) {
			productId := ctx.getUuidParam("id")
			var request dto_request.TiktokProductUpdateRequest
			ctx.mustBind(&request)
			request.ProductId = productId

			a.tiktokProductUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/recommended-category [post]
//	@Summary	Recommended Category
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		dto_request.TiktokProductRecommendedCategoryRequest	body	dto_request.TiktokProductRecommendedCategoryRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{categories=[]dto_response.TiktokCategoryResponse}}
func (a *TiktokProductApi) RecommendedCategory() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductRecommendedCategory),
		func(ctx apiContext) {
			var request dto_request.TiktokProductRecommendedCategoryRequest
			ctx.mustBind(&request)

			category := a.tiktokProductUseCase.RecommendedCategory(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"category": dto_response.NewTiktokCategoryResponse(category),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/{id}/activate [post]
//	@Summary	Activate Tiktok Product
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		id											path	string										true	"Product Id"
//	@Param		dto_request.TiktokProductActivateRequest	body	dto_request.TiktokProductActivateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *TiktokProductApi) Activate() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductActivate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.TiktokProductActivateRequest
			ctx.mustBind(&request)
			request.ProductId = id

			a.tiktokProductUseCase.Activate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// API:
//
//	@Router		/tiktok-products/{id}/deactivate [post]
//	@Summary	Deactivate Tiktok Product
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		id											path	string										true	"Product Id"
//	@Param		dto_request.TiktokProductDeactivateRequest	body	dto_request.TiktokProductDeactivateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *TiktokProductApi) Deactivate() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductDeactivate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.TiktokProductDeactivateRequest
			ctx.mustBind(&request)
			request.ProductId = id

			a.tiktokProductUseCase.Deactivate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterTiktokProductApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := TiktokProductApi{
		api:                  newApi(useCaseManager),
		tiktokProductUseCase: useCaseManager.TiktokProductUseCase(),
	}

	routerGroup := router.Group("/tiktok-products")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/upload-image", api.UploadImage())
	routerGroup.POST("/brands", api.FetchBrands())
	routerGroup.POST("/categories", api.FetchCategories())
	routerGroup.POST("/categories/:category_id/rules", api.GetCategoryRules())
	routerGroup.POST("/categories/:category_id/attributes", api.GetCategoryAttributes())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.POST("/recommended-category", api.RecommendedCategory())
	routerGroup.PATCH("/:id/activate", api.Activate())
	routerGroup.PATCH("/:id/deactivate", api.Deactivate())

}
