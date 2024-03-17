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
//	@Router		/carts/in-active [get]
//	@Summary	Get current user In-Active Cart
//	@tags		Carts
//	@Accept		json
//	@Param		dto_request.CartFetchInActiveRequest	body	dto_request.CartFetchInActiveRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{carts=[]dto_response.CartResponse}}
func (a *CartApi) FetchInActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartGetActive),
		func(ctx apiContext) {
			var request dto_request.CartFetchInActiveRequest
			ctx.mustBind(&request)

			carts := a.cartUseCase.FetchInActive(ctx.context(), request)

			nodes := util.ConvertArray(carts, dto_response.NewCartResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"carts": nodes,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/carts/{id}/set-active [patch]
//	@Summary	Set InActive Cart to Active
//	@tags		Carts
//	@Accept		json
//	@Param		dto_request.CartSetActiveRequest	body	dto_request.CartSetActiveRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cart=dto_response.CartResponse}}
func (a *CartApi) SetActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartSetActive),
		func(ctx apiContext) {
			var request dto_request.CartSetActiveRequest
			ctx.mustBind(&request)

			cart := a.cartUseCase.SetActive(ctx.context(), request)

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
//	@Router		/carts/set-in-active [patch]
//	@Summary	Set Active Cart to In Active
//	@tags		Carts
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cart=dto_response.CartResponse}}
func (a *CartApi) SetInActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartSetInActive),
		func(ctx apiContext) {
			cart := a.cartUseCase.SetInActive(ctx.context())

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
//	@Router		/carts/items [post]
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
//	@Param		product_unit_id						path	string								true	"Product Unit Id"
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

// API:
//
//	@Router		/carts/{id} [delete]
//	@Summary	Delete Item
//	@tags		Carts
//	@Accept		json
//	@Param		dto_request.CartDeleteRequest	body	dto_request.CartDeleteRequest	true	"Body Request"
//	@Param		id								path	string							true	"Cart Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *CartApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCartDelete),
		func(ctx apiContext) {
			var request dto_request.CartDeleteRequest
			ctx.mustBind(&request)

			a.cartUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
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
	routerGroup.GET("/in-active", api.FetchInActive())
	routerGroup.PATCH("/:id/set-active", api.SetActive())
	routerGroup.PATCH("/set-in-active", api.SetInActive())
	routerGroup.POST("/items", api.AddItem())
	routerGroup.PATCH("/items/:product_unit_id", api.UpdateItem())
	routerGroup.DELETE("/items/:product_unit_id", api.DeleteItem())
	routerGroup.DELETE("/:id", api.Delete())
}
