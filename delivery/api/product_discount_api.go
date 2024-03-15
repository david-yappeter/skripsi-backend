package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
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
			var request dto_request.ProductDiscountUpdateRequest
			ctx.mustBind(&request)

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
			var request dto_request.ProductDiscountDeleteRequest
			ctx.mustBind(&request)

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
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}
