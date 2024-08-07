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
//	@Param		file	formData	file	true	"Body with file"
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
//	@Router		/products/options/product-receive-item-form [post]
//	@Summary	Option for Product Receive Item Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForProductReceiveItemFormRequest	body	dto_request.ProductOptionForProductReceiveItemFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForProductReceiveItemForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForProductReceiveItemForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForProductReceiveItemFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForProductReceiveItemForm(ctx.context(), request)

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
//	@Router		/products/options/delivery-order-item-form [post]
//	@Summary	Option for Delivery Order Item Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForDeliveryOrderItemFormRequest	body	dto_request.ProductOptionForDeliveryOrderItemFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForDeliveryOrderItemForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForDeliveryOrderItemForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForDeliveryOrderItemFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForDeliveryOrderItemForm(ctx.context(), request)

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
//	@Router		/products/options/customer-type-discount-form [post]
//	@Summary	Option for Customer Type Discount Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForCustomerTypeDiscountFormRequest	body	dto_request.ProductOptionForCustomerTypeDiscountFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForCustomerTypeDiscountForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForCustomerTypeDiscountForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForCustomerTypeDiscountFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForCustomerTypeDiscountForm(ctx.context(), request)

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
//	@Router		/products/options/cart-add-item-form [post]
//	@Summary	Option for Cart Add Item Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForCartAddItemFormRequest	body	dto_request.ProductOptionForCartAddItemFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForCartAddItemForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForCartAddItemForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForCartAddItemFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForCartAddItemForm(ctx.context(), request)

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
//	@Router		/products/options/product-discount-form [post]
//	@Summary	Option for Product Discount Form
//	@tags		Products
//	@Accept		json
//	@Param		dto_request.ProductOptionForProductDiscountFormRequest	body	dto_request.ProductOptionForProductDiscountFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductResponse}}
func (a *ProductApi) OptionForProductDiscountForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductOptionForProductDiscountForm),
		func(ctx apiContext) {
			var request dto_request.ProductOptionForProductDiscountFormRequest
			ctx.mustBind(&request)

			products, total := a.productUseCase.OptionForProductDiscountForm(ctx.context(), request)

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

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/product-receive-item-form", api.OptionForProductReceiveItemForm())
	optionRouterGroup.POST("/delivery-order-item-form", api.OptionForDeliveryOrderItemForm())
	optionRouterGroup.POST("/customer-type-discount-form", api.OptionForCustomerTypeDiscountForm())
	optionRouterGroup.POST("/cart-add-item-form", api.OptionForCartAddItemForm())
	optionRouterGroup.POST("/product-discount-form", api.OptionForProductDiscountForm())
}
