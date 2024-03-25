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

type ProductUnitApi struct {
	api
	productUnitUseCase use_case.ProductUnitUseCase
}

// API:
//
//	@Router		/product-units [post]
//	@Summary	Create
//	@tags		Product Units
//	@Accept		json
//	@Param		dto_request.ProductUnitCreateRequest	body	dto_request.ProductUnitCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_unit=dto_response.ProductUnitResponse}}
func (a *ProductUnitApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUnitCreate),
		func(ctx apiContext) {
			var request dto_request.ProductUnitCreateRequest
			ctx.mustBind(&request)

			productUnit := a.productUnitUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_unit": dto_response.NewProductUnitResponse(productUnit),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-units/{id} [delete]
//	@Summary	Delete
//	@tags		Product Units
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductUnitDeleteRequest	body	dto_request.ProductUnitDeleteRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *ProductUnitApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUnitDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductUnitDeleteRequest
			ctx.mustBind(&request)
			request.ProductUnitId = id

			a.productUnitUseCase.Delete(ctx.context(), request)

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
//	@Router		/product-units/options/product-receive-item-form [post]
//	@Summary	Option for Product Receive Item Form
//	@tags		Product Units
//	@Accept		json
//	@Param		dto_request.ProductUnitOptionForProductReceiveItemFormRequest	body	dto_request.ProductUnitOptionForProductReceiveItemFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductUnitResponse}}
func (a *ProductUnitApi) OptionForProductReceiveItemForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUnitOptionForProductReceiveItemForm),
		func(ctx apiContext) {
			var request dto_request.ProductUnitOptionForProductReceiveItemFormRequest
			ctx.mustBind(&request)

			productUnits, total := a.productUnitUseCase.OptionForProductReceiveItemForm(ctx.context(), request)

			nodes := util.ConvertArray(productUnits, dto_response.NewProductUnitResponse)

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
//	@Router		/product-units/options/delivery-order-form [post]
//	@Summary	Option for Delivery Order Form
//	@tags		Product Units
//	@Accept		json
//	@Param		dto_request.ProductUnitOptionForDeliveryOrderFormRequest	body	dto_request.ProductUnitOptionForDeliveryOrderFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductUnitResponse}}
func (a *ProductUnitApi) OptionForDeliveryOrderForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUnitOptionForDeliveryOrderForm),
		func(ctx apiContext) {
			var request dto_request.ProductUnitOptionForDeliveryOrderFormRequest
			ctx.mustBind(&request)

			productUnits, total := a.productUnitUseCase.OptionForDeliveryOrderForm(ctx.context(), request)

			nodes := util.ConvertArray(productUnits, dto_response.NewProductUnitResponse)

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

func RegisterProductUnitApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductUnitApi{
		api:                newApi(useCaseManager),
		productUnitUseCase: useCaseManager.ProductUnitUseCase(),
	}

	routerGroup := router.Group("/product-units")
	routerGroup.POST("", api.Create())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/product-receive-item-form", api.OptionForProductReceiveItemForm())
	optionRouterGroup.POST("/delivery-order-form", api.OptionForDeliveryOrderForm())
}
