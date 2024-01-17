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

type AdminSupplierTypeApi struct {
	api
	supplierTypeUseCase use_case.SupplierTypeUseCase
}

// API:
//
//	@Router		/admin/supplier-types [post]
//	@Summary	Create
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		dto_request.AdminSupplierTypeCreateRequest	body	dto_request.AdminSupplierTypeCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier_type=dto_response.SupplierTypeResponse}}
func (a *AdminSupplierTypeApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierTypeCreate),
		func(ctx apiContext) {
			var request dto_request.AdminSupplierTypeCreateRequest
			ctx.mustBind(&request)

			supplierType := a.supplierTypeUseCase.AdminCreate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"supplier_type": dto_response.NewSupplierTypeResponse(supplierType),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/admin/supplier-types/filter [post]
//	@Summary	Filter
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		dto_request.AdminSupplierTypeFetchRequest	body	dto_request.AdminSupplierTypeFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.SupplierTypeResponse}}
func (a *AdminSupplierTypeApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierTypeCreate),
		func(ctx apiContext) {
			var request dto_request.AdminSupplierTypeFetchRequest
			ctx.mustBind(&request)

			supplierTypes, total := a.supplierTypeUseCase.AdminFetch(ctx.context(), request)

			nodes := util.ConvertArray(supplierTypes, dto_response.NewSupplierTypeResponse)

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
//	@Router		/admin/supplier-types/{id} [get]
//	@Summary	Get
//	@tags		Admin Suppliers
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier_type=dto_response.SupplierTypeResponse}}
func (a *AdminSupplierTypeApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierTypeGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminSupplierTypeGetRequest
			ctx.mustBind(&request)

			request.SupplierTypeId = id

			supplierType := a.supplierTypeUseCase.AdminGet(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"supplier_type": dto_response.NewSupplierTypeResponse(supplierType),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/admin/supplier-types/{id} [put]
//	@Summary	Update
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.AdminSupplierTypeUpdateRequest	body	dto_request.AdminSupplierTypeUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier_type=dto_response.SupplierTypeResponse}}
func (a *AdminSupplierTypeApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierTypeUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminSupplierTypeUpdateRequest
			ctx.mustBind(&request)

			request.SupplierTypeId = id

			supplierType := a.supplierTypeUseCase.AdminUpdate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"supplier_type": dto_response.NewSupplierTypeResponse(supplierType),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/admin/supplier-types/{id} [delete]
//	@Summary	Delete
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *AdminSupplierTypeApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierTypeDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminSupplierTypeDeleteRequest
			ctx.mustBind(&request)
			request.SupplierTypeId = id

			a.supplierTypeUseCase.AdminDelete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterAdminSupplierTypeApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AdminSupplierTypeApi{
		api:                 newApi(useCaseManager),
		supplierTypeUseCase: useCaseManager.SupplierTypeUseCase(),
	}

	adminRouterGroup := router.Group("/admin/supplier-types")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.DELETE("/:id", api.Delete())
}
