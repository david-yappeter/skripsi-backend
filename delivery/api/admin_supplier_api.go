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

type AdminSupplierApi struct {
	api
	supplierUseCase use_case.SupplierUseCase
}

//	@Router		/admin/suppliers [post]
//	@Summary	Create
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		dto_request.AdminSupplierCreateRequest	body	dto_request.AdminSupplierCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier=dto_response.SupplierResponse}}
func (a *AdminSupplierApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierCreate),
		func(ctx apiContext) {
			var request dto_request.AdminSupplierCreateRequest
			ctx.mustBind(&request)

			supplier := a.supplierUseCase.AdminCreate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"supplier": dto_response.NewSupplierResponse(supplier),
					},
				},
			)
		},
	)
}

//	@Router		/admin/suppliers/filter [post]
//	@Summary	Filter
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		dto_request.AdminSupplierFetchRequest	body	dto_request.AdminSupplierFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{suppliers=[]dto_response.SupplierResponse}}
func (a *AdminSupplierApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierCreate),
		func(ctx apiContext) {
			var request dto_request.AdminSupplierFetchRequest
			ctx.mustBind(&request)

			suppliers, total := a.supplierUseCase.AdminFetch(ctx.context(), request)

			nodes := util.ConvertArray(suppliers, dto_response.NewSupplierResponse)

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

//	@Router		/admin/suppliers/{id} [get]
//	@Summary	Update
//	@tags		Admin Suppliers
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier=dto_response.SupplierResponse}}
func (a *AdminSupplierApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminSupplierGetRequest
			ctx.mustBind(&request)

			request.SupplierId = id

			supplier := a.supplierUseCase.AdminGet(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"supplier": dto_response.NewSupplierResponse(supplier),
					},
				},
			)
		},
	)
}

//	@Router		/admin/suppliers/{id} [put]
//	@Summary	Update
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.AdminSupplierUpdateRequest	body	dto_request.AdminSupplierUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier=dto_response.SupplierResponse}}
func (a *AdminSupplierApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminSupplierUpdateRequest
			ctx.mustBind(&request)

			request.SupplierId = id

			supplier := a.supplierUseCase.AdminUpdate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"supplier": dto_response.NewSupplierResponse(supplier),
					},
				},
			)
		},
	)
}

//	@Router		/admin/suppliers/{id} [delete]
//	@Summary	Update Password
//	@tags		Admin Suppliers
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *AdminSupplierApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminSupplierDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminSupplierDeleteRequest
			ctx.mustBind(&request)
			request.SupplierId = id

			a.supplierUseCase.AdminDelete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterAdminSupplierApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AdminSupplierApi{
		api:             newApi(useCaseManager),
		supplierUseCase: useCaseManager.SupplierUseCase(),
	}

	adminRouterGroup := router.Group("/admin/suppliers")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.DELETE("/:id", api.Delete())
}
