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

type CustomerTypeApi struct {
	api
	customerTypeUseCase use_case.CustomerTypeUseCase
}

// API:
//
//	@Router		/customer-types [post]
//	@Summary	Create
//	@tags		Customer Types
//	@Accept		json
//	@Param		dto_request.CustomerTypeCreateRequest	body	dto_request.CustomerTypeCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer_type=dto_response.CustomerTypeResponse}}
func (a *CustomerTypeApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeCreate),
		func(ctx apiContext) {
			var request dto_request.CustomerTypeCreateRequest
			ctx.mustBind(&request)

			customerType := a.customerTypeUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer_type": dto_response.NewCustomerTypeResponse(customerType),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customer-types/filter [post]
//	@Summary	Filter
//	@tags		Customer Types
//	@Accept		json
//	@Param		dto_request.CustomerTypeFetchRequest	body	dto_request.CustomerTypeFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerTypeResponse}}
func (a *CustomerTypeApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeFetch),
		func(ctx apiContext) {
			var request dto_request.CustomerTypeFetchRequest
			ctx.mustBind(&request)

			customerTypes, total := a.customerTypeUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(customerTypes, dto_response.NewCustomerTypeResponse)

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
//	@Router		/customer-types/{id} [get]
//	@Summary	Get
//	@tags		Customer Types
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer_type=dto_response.CustomerTypeResponse}}
func (a *CustomerTypeApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerTypeGetRequest
			ctx.mustBind(&request)

			request.CustomerTypeId = id

			customerType := a.customerTypeUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer_type": dto_response.NewCustomerTypeResponse(customerType),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customer-types/{id} [put]
//	@Summary	Update
//	@tags		Customer Types
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.CustomerTypeUpdateRequest	body	dto_request.CustomerTypeUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer_type=dto_response.CustomerTypeResponse}}
func (a *CustomerTypeApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerTypeUpdateRequest
			ctx.mustBind(&request)

			request.CustomerTypeId = id

			customerType := a.customerTypeUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer_type": dto_response.NewCustomerTypeResponse(customerType),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customer-types/{id} [delete]
//	@Summary	Delete
//	@tags		Customer Types
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *CustomerTypeApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerTypeDeleteRequest
			ctx.mustBind(&request)
			request.CustomerTypeId = id

			a.customerTypeUseCase.Delete(ctx.context(), request)

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
//	@Router		/customer-types/options/customer-form [post]
//	@Summary	Option for Customer Form
//	@tags		Customer Types
//	@Accept		json
//	@Param		dto_request.CustomerTypeOptionForCustomerFormRequest	body	dto_request.CustomerTypeOptionForCustomerFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerTypeResponse}}
func (a *CustomerTypeApi) OptionForCustomerForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeOptionForCustomerForm),
		func(ctx apiContext) {
			var request dto_request.CustomerTypeOptionForCustomerFormRequest
			ctx.mustBind(&request)

			customerTypes, total := a.customerTypeUseCase.OptionForCustomerForm(ctx.context(), request)

			nodes := util.ConvertArray(customerTypes, dto_response.NewCustomerTypeResponse)

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

func RegisterCustomerTypeApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CustomerTypeApi{
		api:                 newApi(useCaseManager),
		customerTypeUseCase: useCaseManager.CustomerTypeUseCase(),
	}

	routerGroup := router.Group("/customer-types")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/customer-form", api.OptionForCustomerForm())
}
