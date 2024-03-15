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

type ProductApi struct {
	api
	productUseCase use_case.ProductUseCase
}

// API:
//
//	@Router		/products [post]
//	@Summary	Create
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductCreateRequest	body	dto_request.ProductCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product=dto_response.ProductResponse}}
func (a *ProductApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductCreate),
		func(ctx apiContext) {
			var request dto_request.ProductCreateRequest
			ctx.mustBind(&request)

			product := a.productUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product": dto_response.NewProductResponse(product),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/upload [post]
//	@Summary	Upload
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductUploadRequest	body	dto_request.ProductUploadRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *ProductApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUpload),
		func(ctx apiContext) {
			var request dto_request.ProductUploadRequest
			ctx.mustBind(&request)

			path := a.productUseCase.Upload(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"path": path,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/filter [post]
//	@Summary	Filter
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductFetchRequest	body	dto_request.ProductFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductCreate),
		func(ctx apiContext) {
			var request dto_request.ProductFetchRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(products, dto_response.NewProductResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: nodes,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/{id} [get]
//	@Summary	Get
//	@tags		Products
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product=dto_response.ProductResponse}}
func (a *ProductApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductGetRequest
			ctx.mustBind(&request)

			request.ProductId = id

			product := a.productUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product": dto_response.NewProductResponse(product),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/{id} [put]
//	@Summary	Update
//	@tags		Products
//	@Accept		json
//	@Param		id									path	string								true	"Id"
//	@Param		dto_request.ProductUpdateRequest	body	dto_request.ProductUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product=dto_response.ProductResponse}}
func (a *ProductApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductUpdateRequest
			ctx.mustBind(&request)

			request.ProductId = id

			product := a.productUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product": dto_response.NewProductResponse(product),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/{id} [delete]
//	@Summary	Delete
//	@tags		Products
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *ProductApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductDeleteRequest
			ctx.mustBind(&request)
			request.ProductId = id

			a.productUseCase.Delete(ctx.context(), request)

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
//	@Router		/products/options/product-receive-form [post]
//	@Summary	Option for Product Receive Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForProductReceiveFormRequest	body	dto_request.ProductOptionForProductReceiveFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForProductReceiveForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForProductReceiveForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForProductReceiveFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForProductReceiveForm(ctx.context(), request)

			nodes := util.ConvertArray(products, dto_response.NewProductResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: nodes,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/options/delivery-order-form [post]
//	@Summary	Option for Delivery Order Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForDeliveryOrderFormRequest	body	dto_request.ProductOptionForDeliveryOrderFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForDeliveryOrderForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForDeliveryOrderForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForDeliveryOrderFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForDeliveryOrderForm(ctx.context(), request)

			nodes := util.ConvertArray(products, dto_response.NewProductResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: nodes,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/products/options/customer-type-form [post]
//	@Summary	Option for Customer Type Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForCustomerTypeFormRequest	body	dto_request.ProductOptionForCustomerTypeFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForCustomerTypeForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForCustomerTypeForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForCustomerTypeFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForCustomerTypeForm(ctx.context(), request)

			nodes := util.ConvertArray(products, dto_response.NewProductResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: nodes,
					},
				},
			)
		},
	)
}

func RegisterProductApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductApi{
		api:            newApi(useCaseManager),
		productUseCase: useCaseManager.ProductUseCase(),
	}

	routerGroup := router.Group("/products")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/upload", api.Upload())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/product-receive-form", api.OptionForProductReceiveForm())
	optionRouterGroup.POST("/delivery-order-form", api.OptionForDeliveryOrderForm())
	optionRouterGroup.POST("/customer-type-form", api.OptionForCustomerTypeForm())
}
