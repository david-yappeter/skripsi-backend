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

type SupplierTypeApi struct {
	api
	supplierTypeUseCase use_case.SupplierTypeUseCase
}

// API:
//
//	@Router		/supplier-types [post]
//	@Summary	Create
//	@tags		Suppliers
//	@Accept		json
//	@Param		dto_request.SupplierTypeCreateRequest	body	dto_request.SupplierTypeCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier_type=dto_response.SupplierTypeResponse}}
func (a *SupplierTypeApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierTypeCreate),
		func(ctx apiContext) {
			var request dto_request.SupplierTypeCreateRequest
			ctx.mustBind(&request)

			supplierType := a.supplierTypeUseCase.Create(ctx.context(), request)

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
//	@Router		/supplier-types/filter [post]
//	@Summary	Filter
//	@tags		Suppliers
//	@Accept		json
//	@Param		dto_request.SupplierTypeFetchRequest	body	dto_request.SupplierTypeFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.SupplierTypeResponse}}
func (a *SupplierTypeApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierTypeCreate),
		func(ctx apiContext) {
			var request dto_request.SupplierTypeFetchRequest
			ctx.mustBind(&request)

			supplierTypes, total := a.supplierTypeUseCase.Fetch(ctx.context(), request)

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
//	@Router		/supplier-types/{id} [get]
//	@Summary	Get
//	@tags		Suppliers
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier_type=dto_response.SupplierTypeResponse}}
func (a *SupplierTypeApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierTypeGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.SupplierTypeGetRequest
			ctx.mustBind(&request)

			request.SupplierTypeId = id

			supplierType := a.supplierTypeUseCase.Get(ctx.context(), request)

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
//	@Router		/supplier-types/{id} [put]
//	@Summary	Update
//	@tags		Suppliers
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.SupplierTypeUpdateRequest	body	dto_request.SupplierTypeUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier_type=dto_response.SupplierTypeResponse}}
func (a *SupplierTypeApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierTypeUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.SupplierTypeUpdateRequest
			ctx.mustBind(&request)

			request.SupplierTypeId = id

			supplierType := a.supplierTypeUseCase.Update(ctx.context(), request)

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
//	@Router		/supplier-types/{id} [delete]
//	@Summary	Delete
//	@tags		Suppliers
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *SupplierTypeApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierTypeDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.SupplierTypeDeleteRequest
			ctx.mustBind(&request)
			request.SupplierTypeId = id

			a.supplierTypeUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterSupplierTypeApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := SupplierTypeApi{
		api:                 newApi(useCaseManager),
		supplierTypeUseCase: useCaseManager.SupplierTypeUseCase(),
	}

	adminRouterGroup := router.Group("/supplier-types")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.DELETE("/:id", api.Delete())
}
