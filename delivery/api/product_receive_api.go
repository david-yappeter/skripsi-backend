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

type ProductReceiveApi struct {
	api
	productReceiveUseCase use_case.ProductReceiveUseCase
}

//	@Router		/product-receives [post]
//	@Summary	Create
//	@tags		Product Receives
//	@Accept		json
//	@Param		dto_request.ProductReceiveCreateRequest	body	dto_request.ProductReceiveCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveCreate),
		func(ctx apiContext) {
			var request dto_request.ProductReceiveCreateRequest
			ctx.mustBind(&request)

			productReceive := a.productReceiveUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

//	@Router		/product-receives/{id}/items [post]
//	@Summary	CreaAddItemte
//	@tags		Product Receives
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.ProductReceiveAddItemRequest	body	dto_request.ProductReceiveAddItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) AddItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveAddItem),
		func(ctx apiContext) {
			var request dto_request.ProductReceiveAddItemRequest
			ctx.mustBind(&request)

			productReceive := a.productReceiveUseCase.AddItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

//	@Router		/product-receives/filter [post]
//	@Summary	Filter
//	@tags		Product Receives
//	@Accept		json
//	@Param		dto_request.ProductReceiveFetchRequest	body	dto_request.ProductReceiveFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveCreate),
		func(ctx apiContext) {
			var request dto_request.ProductReceiveFetchRequest
			ctx.mustBind(&request)

			productReceives, total := a.productReceiveUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(productReceives, dto_response.NewProductReceiveResponse)

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

//	@Router		/product-receives/{id} [get]
//	@Summary	Update
//	@tags		Product Receives
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveGetRequest
			ctx.mustBind(&request)

			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

//	@Router		/product-receives/{id} [delete]
//	@Summary	Update Password
//	@tags		Product Receives
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *ProductReceiveApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveDeleteRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			a.productReceiveUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterProductReceiveApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductReceiveApi{
		api:                   newApi(useCaseManager),
		productReceiveUseCase: useCaseManager.ProductReceiveUseCase(),
	}

	adminRouterGroup := router.Group("/product-receives")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.DELETE("/:id", api.Delete())

	adminRouterGroup.POST("/:id/items", api.AddItem())
}
