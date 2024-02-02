package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CartApi struct {
	api
	cartUseCase use_case.CartUseCase
}

// API:
//
//	@Router		/carts/active [get]
//	@Summary	Get current user Active Cart
//	@tags		Carts
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cart=dto_response.CartResponse}}
func (a *CartApi) GetActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartGetActive),
		func(ctx apiContext) {
			cart := a.cartUseCase.GetCurrent(ctx.context())

			var resp *dto_response.CartResponse
			if cart != nil {
				resp = dto_response.NewCartResponseP(*cart)
			}

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cart": resp,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/carts/items/{product_unit_id} [post]
//	@Summary	Add Item
//	@tags		Carts
//	@Accept		json
//	@Param		dto_request.CartAddItemRequest	body	dto_request.CartAddItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cart=dto_response.CartResponse}}
func (a *CartApi) AddItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartAddItem),
		func(ctx apiContext) {
			var request dto_request.CartAddItemRequest
			ctx.mustBind(&request)

			cart := a.cartUseCase.AddItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cart": dto_response.NewCartResponse(cart),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/carts/items/{product_unit_id} [patch]
//	@Summary	Update Item
//	@tags		Carts
//	@Accept		json
//	@Param		dto_request.CartUpdateItemRequest	body	dto_request.CartUpdateItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cart=dto_response.CartResponse}}
func (a *CartApi) UpdateItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartAddItem),
		func(ctx apiContext) {
			var request dto_request.CartUpdateItemRequest
			ctx.mustBind(&request)

			cart := a.cartUseCase.UpdateItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cart": dto_response.NewCartResponse(cart),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/carts/items/{product_unit_id} [delete]
//	@Summary	Delete Item
//	@tags		Carts
//	@Accept		json
//	@Param		dto_request.CartDeleteItemRequest	body	dto_request.CartDeleteItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cart=dto_response.CartResponse}}
func (a *CartApi) DeleteItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartAddItem),
		func(ctx apiContext) {
			var request dto_request.CartDeleteItemRequest
			ctx.mustBind(&request)

			cart := a.cartUseCase.DeleteItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cart": dto_response.NewCartResponse(cart),
					},
				},
			)
		},
	)
}

func RegisterCartApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CartApi{
		api:         newApi(useCaseManager),
		cartUseCase: useCaseManager.CartUseCase(),
	}

	routerGroup := router.Group("/carts")
	routerGroup.GET("/active", api.GetActive())
	routerGroup.POST("/items/:product_unit_id", api.AddItem())
	routerGroup.PATCH("/items/:product_unit_id", api.UpdateItem())
	routerGroup.DELETE("/items/:product_unit_id", api.DeleteItem())
}
