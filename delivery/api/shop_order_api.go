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

type ShopOrderApi struct {
	api
	shopOrderUseCase use_case.ShopOrderUseCase
}

// API:
//
//	@Router		/shop-orders/filter [post]
//	@Summary	Fetch
//	@tags		Shop Orders
//	@Accept		json
//	@Param		dto_request.ShopOrderFetchRequest	body	dto_request.ShopOrderFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ShopOrderResponse}}
func (a *ShopOrderApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionShopOrderFetch),
		func(ctx apiContext) {
			var request dto_request.ShopOrderFetchRequest
			ctx.mustBind(&request)

			shopOrders, total := a.shopOrderUseCase.Fetch(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: util.ConvertArray(shopOrders, dto_response.NewShopOrderResponse),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/shop-orders/{id} [get]
//	@Summary	Get
//	@tags		Shop Orders
//	@Accept		json
//	@Param		id	path	string	true	"Shop Order Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{shop_order=dto_response.ShopOrderResponse}}
func (a *ShopOrderApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionShopOrderGet),
		func(ctx apiContext) {
			var request dto_request.ShopOrderGetRequest
			ctx.mustBind(&request)

			shopOrder := a.shopOrderUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"shop_order": dto_response.NewShopOrderResponse(shopOrder),
					},
				},
			)
		},
	)
}

func RegisterShopOrderApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ShopOrderApi{
		api:              newApi(useCaseManager),
		shopOrderUseCase: useCaseManager.ShopOrderUseCase(),
	}

	routerGroup := router.Group("/shop-orders")

	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
}
