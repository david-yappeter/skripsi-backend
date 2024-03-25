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

type SupplierApi struct {
	api
	supplierUseCase use_case.SupplierUseCase
}

// API:
//
//	@Router		/suppliers [post]
//	@Summary	Create
//	@tags		Suppliers
//	@Accept		json
//	@Param		dto_request.SupplierCreateRequest	body	dto_request.SupplierCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier=dto_response.SupplierResponse}}
func (a *SupplierApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierCreate),
		func(ctx apiContext) {
			var request dto_request.SupplierCreateRequest
			ctx.mustBind(&request)

			supplier := a.supplierUseCase.Create(ctx.context(), request)

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

// API:
//
//	@Router		/suppliers/filter [post]
//	@Summary	Filter
//	@tags		Suppliers
//	@Accept		json
//	@Param		dto_request.SupplierFetchRequest	body	dto_request.SupplierFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.SupplierResponse}}
func (a *SupplierApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierCreate),
		func(ctx apiContext) {
			var request dto_request.SupplierFetchRequest
			ctx.mustBind(&request)

			suppliers, total := a.supplierUseCase.Fetch(ctx.context(), request)

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

// API:
//
//	@Router		/suppliers/{id} [get]
//	@Summary	Get
//	@tags		Suppliers
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier=dto_response.SupplierResponse}}
func (a *SupplierApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.SupplierGetRequest
			ctx.mustBind(&request)

			request.SupplierId = id

			supplier := a.supplierUseCase.Get(ctx.context(), request)

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

// API:
//
//	@Router		/suppliers/{id} [put]
//	@Summary	Update
//	@tags		Suppliers
//	@Accept		json
//	@Param		id									path	string								true	"Id"
//	@Param		dto_request.SupplierUpdateRequest	body	dto_request.SupplierUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{supplier=dto_response.SupplierResponse}}
func (a *SupplierApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.SupplierUpdateRequest
			ctx.mustBind(&request)

			request.SupplierId = id

			supplier := a.supplierUseCase.Update(ctx.context(), request)

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

// API:
//
//	@Router		/suppliers/{id} [delete]
//	@Summary	Delete
//	@tags		Suppliers
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *SupplierApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.SupplierDeleteRequest
			ctx.mustBind(&request)
			request.SupplierId = id

			a.supplierUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// API:
//
//	@Router		/suppliers/options/product-receive-item-form [post]
//	@Summary	Option for Product Receive Form
//	@tags		Suppliers
//	@Accept		json
//	@Param		dto_request.SupplierOptionForProductReceiveItemFormRequest	body	dto_request.SupplierOptionForProductReceiveItemFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *SupplierApi) OptionForProductReceiveItemForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierOptionForProductReceiveItemForm),
		func(ctx apiContext) {
			var request dto_request.SupplierOptionForProductReceiveItemFormRequest
			ctx.mustBind(&request)

			suppliers, total := a.supplierUseCase.OptionForProductReceiveItemForm(ctx.context(), request)

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

// API:
//
//	@Router		/suppliers/options/product-receive-filter [post]
//	@Summary	Option for Product Receive Filter
//	@tags		Suppliers
//	@Accept		json
//	@Param		dto_request.SupplierOptionForProductReceiveFilterRequest	body	dto_request.SupplierOptionForProductReceiveFilterRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *SupplierApi) OptionForProductReceiveFilter() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSupplierOptionForProductReceiveFilter),
		func(ctx apiContext) {
			var request dto_request.SupplierOptionForProductReceiveFilterRequest
			ctx.mustBind(&request)

			suppliers, total := a.supplierUseCase.OptionForProductReceiveFilter(ctx.context(), request)

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

func RegisterSupplierApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := SupplierApi{
		api:             newApi(useCaseManager),
		supplierUseCase: useCaseManager.SupplierUseCase(),
	}

	routerGroup := router.Group("/suppliers")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/product-receive-item-form", api.OptionForProductReceiveItemForm())
	optionRouterGroup.POST("/product-receive-filter", api.OptionForProductReceiveFilter())
}
