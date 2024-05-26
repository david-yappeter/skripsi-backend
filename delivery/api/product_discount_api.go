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

type ProductDiscountApi struct {
	api
	productDiscountUseCase use_case.ProductDiscountUseCase
}

// API:
//
//	@Router		/product-discounts [post]
//	@Summary	Create
//	@tags		Product Discounts
//	@Accept		json
//	@Param		dto_request.ProductDiscountCreateRequest	body	dto_request.ProductDiscountCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_discount=dto_response.ProductDiscountResponse}}
func (a *ProductDiscountApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductDiscountCreate),
		func(ctx apiContext) {
			var request dto_request.ProductDiscountCreateRequest
			ctx.mustBind(&request)

			productDiscount := a.productDiscountUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_discount": dto_response.NewProductDiscountResponse(productDiscount),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-discounts/filter [post]
//	@Summary	Fetch
//	@tags		Product Discounts
//	@Accept		json
//	@Param		dto_request.ProductDiscountFetchRequest	body	dto_request.ProductDiscountFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductDiscountResponse}}
func (a *ProductDiscountApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductDiscountFetch),
		func(ctx apiContext) {
			var request dto_request.ProductDiscountFetchRequest
			ctx.mustBind(&request)

			productDiscounts, total := a.productDiscountUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(productDiscounts, dto_response.NewProductDiscountResponse)

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
//	@Router		/product-discounts/{id} [get]
//	@Summary	Get
//	@tags		Product Discounts
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductDiscountGetRequest	body	dto_request.ProductDiscountGetRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_discount=dto_response.ProductDiscountResponse}}
func (a *ProductDiscountApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductDiscountGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductDiscountGetRequest
			ctx.mustBind(&request)

			request.ProductDiscountId = id

			productDiscount := a.productDiscountUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_discount": dto_response.NewProductDiscountResponse(productDiscount),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-discounts/{id} [put]
//	@Summary	Update
//	@tags		Product Discounts
//	@Accept		json
//	@Param		id											path	string										true	"Product Discount Id"
//	@Param		dto_request.ProductDiscountUpdateRequest	body	dto_request.ProductDiscountUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_discount=dto_response.ProductDiscountResponse}}
func (a *ProductDiscountApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductDiscountUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductDiscountUpdateRequest
			ctx.mustBind(&request)

			request.ProductDiscountId = id

			productDiscount := a.productDiscountUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_discount": dto_response.NewProductDiscountResponse(productDiscount),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-discounts/{id} [delete]
//	@Summary	Delete
//	@tags		Product Discounts
//	@Accept		json
//	@Param		id											path	string										true	"Product Discount Id"
//	@Param		dto_request.ProductDiscountDeleteRequest	body	dto_request.ProductDiscountDeleteRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *ProductDiscountApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductDiscountDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.ProductDiscountDeleteRequest
			ctx.mustBind(&request)

			request.ProductDiscountId = id

			a.productDiscountUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterProductDiscountApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductDiscountApi{
		api:                    newApi(useCaseManager),
		productDiscountUseCase: useCaseManager.ProductDiscountUseCase(),
	}

	routerGroup := router.Group("/product-discounts")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())

	routerGroup.GET("/{id}", api.Get())

	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}
