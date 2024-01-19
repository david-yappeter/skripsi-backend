package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
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
//	@Router		/product-units/upload [post]
//	@Summary	Upload
//	@tags		Product Units
//	@Accept		json
//	@Param		dto_request.ProductUnitUploadRequest	body	dto_request.ProductUnitUploadRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *ProductUnitApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUnitUpload),
		func(ctx apiContext) {
			var request dto_request.ProductUnitUploadRequest
			ctx.mustBind(&request)

			path := a.productUnitUseCase.Upload(ctx.context(), request)

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
//	@Router		/product-units/{id} [put]
//	@Summary	Update
//	@tags		Product Units
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductUnitUpdateRequest	body	dto_request.ProductUnitUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_unit=dto_response.ProductUnitResponse}}
func (a *ProductUnitApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductUnitUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductUnitUpdateRequest
			ctx.mustBind(&request)
			request.ProductUnitId = id

			productUnit := a.productUnitUseCase.Update(ctx.context(), request)

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

func RegisterProductUnitApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductUnitApi{
		api:                newApi(useCaseManager),
		productUnitUseCase: useCaseManager.ProductUnitUseCase(),
	}

	routerGroup := router.Group("/product-units")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/upload", api.Upload())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}
