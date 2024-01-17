package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminProductUnitApi struct {
	api
	productUnitUseCase use_case.ProductUnitUseCase
}

// API:
//
//	@Router		/admin/product-units [post]
//	@Summary	Create
//	@tags		Admin Product Units
//	@Accept		json
//	@Param		dto_request.AdminProductUnitCreateRequest	body	dto_request.AdminProductUnitCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_unit=dto_response.ProductUnitResponse}}
func (a *AdminProductUnitApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminProductUnitCreate),
		func(ctx apiContext) {
			var request dto_request.AdminProductUnitCreateRequest
			ctx.mustBind(&request)

			productUnit := a.productUnitUseCase.AdminCreate(ctx.context(), request)

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
//	@Router		/admin/product-units/upload [post]
//	@Summary	Upload
//	@tags		Admin Product Units
//	@Accept		json
//	@Param		dto_request.AdminProductUnitUploadRequest	body	dto_request.AdminProductUnitUploadRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *AdminProductUnitApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminProductUnitUpload),
		func(ctx apiContext) {
			var request dto_request.AdminProductUnitUploadRequest
			ctx.mustBind(&request)

			path := a.productUnitUseCase.AdminUpload(ctx.context(), request)

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
//	@Router		/admin/product-units/{id} [put]
//	@Summary	Update
//	@tags		Admin Product Units
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.AdminProductUnitUpdateRequest	body	dto_request.AdminProductUnitUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_unit=dto_response.ProductUnitResponse}}
func (a *AdminProductUnitApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminProductUnitUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminProductUnitUpdateRequest
			ctx.mustBind(&request)
			request.ProductUnitId = id

			productUnit := a.productUnitUseCase.AdminUpdate(ctx.context(), request)

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
//	@Router		/admin/product-units/{id} [delete]
//	@Summary	Delete
//	@tags		Admin Product Units
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.AdminProductUnitDeleteRequest	body	dto_request.AdminProductUnitDeleteRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *AdminProductUnitApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminProductUnitDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminProductUnitDeleteRequest
			ctx.mustBind(&request)
			request.ProductUnitId = id

			a.productUnitUseCase.AdminDelete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterAdminProductUnitApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AdminProductUnitApi{
		api:                newApi(useCaseManager),
		productUnitUseCase: useCaseManager.ProductUnitUseCase(),
	}

	adminRouterGroup := router.Group("/admin/product-units")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/upload", api.Upload())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.DELETE("/:id", api.Delete())
}
