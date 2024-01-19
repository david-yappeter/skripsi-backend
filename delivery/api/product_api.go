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
//	@Param		id								path	string							true	"Id"
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

func RegisterProductApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductApi{
		api:            newApi(useCaseManager),
		productUseCase: useCaseManager.ProductUseCase(),
	}

	routerGroup := router.Group("/products")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}