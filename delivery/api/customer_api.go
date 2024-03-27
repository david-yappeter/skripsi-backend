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

type CustomerApi struct {
	api
	customerUseCase use_case.CustomerUseCase
}

// API:
//
//	@Router		/customers [post]
//	@Summary	Create
//	@tags		Customers
//	@Accept		json
//	@Param		dto_request.CustomerCreateRequest	body	dto_request.CustomerCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer=dto_response.CustomerResponse}}
func (a *CustomerApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerCreate),
		func(ctx apiContext) {
			var request dto_request.CustomerCreateRequest
			ctx.mustBind(&request)

			customer := a.customerUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer": dto_response.NewCustomerResponse(customer),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customers/filter [post]
//	@Summary	Filter
//	@tags		Customers
//	@Accept		json
//	@Param		dto_request.CustomerFetchRequest	body	dto_request.CustomerFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerResponse}}
func (a *CustomerApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerFetch),
		func(ctx apiContext) {
			var request dto_request.CustomerFetchRequest
			ctx.mustBind(&request)

			customers, total := a.customerUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(customers, dto_response.NewCustomerResponse)

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
//	@Router		/customers/{id} [get]
//	@Summary	Get
//	@tags		Customers
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer=dto_response.CustomerResponse}}
func (a *CustomerApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerGetRequest
			ctx.mustBind(&request)

			request.CustomerId = id

			customer := a.customerUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer": dto_response.NewCustomerResponse(customer),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customers/{id} [put]
//	@Summary	Update
//	@tags		Customers
//	@Accept		json
//	@Param		id									path	string								true	"Id"
//	@Param		dto_request.CustomerUpdateRequest	body	dto_request.CustomerUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer=dto_response.CustomerResponse}}
func (a *CustomerApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerUpdateRequest
			ctx.mustBind(&request)

			request.CustomerId = id

			customer := a.customerUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer": dto_response.NewCustomerResponse(customer),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customers/{id} [delete]
//	@Summary	Delete
//	@tags		Customers
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *CustomerApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerDeleteRequest
			ctx.mustBind(&request)
			request.CustomerId = id

			a.customerUseCase.Delete(ctx.context(), request)

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
//	@Router		/customers/options/delivery-order-form [post]
//	@Summary	Option for Delivery Order Form
//	@tags		Customers
//	@Accept		json
//	@Param		dto_request.CustomerOptionForDeliveryOrderFormRequest	body	dto_request.CustomerOptionForDeliveryOrderFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerResponse}}
func (a *CustomerApi) OptionForDeliveryOrderForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerOptionForDeliveryOrderForm),
		func(ctx apiContext) {
			var request dto_request.CustomerOptionForDeliveryOrderFormRequest
			ctx.mustBind(&request)

			customers, total := a.customerUseCase.OptionForDeliveryOrderForm(ctx.context(), request)

			nodes := util.ConvertArray(customers, dto_response.NewCustomerResponse)

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
//	@Router		/customers/options/delivery-order-filter [post]
//	@Summary	Option for Delivery Order Filter
//	@tags		Customers
//	@Accept		json
//	@Param		dto_request.CustomerOptionForDeliveryOrderFilterRequest	body	dto_request.CustomerOptionForDeliveryOrderFilterRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerResponse}}
func (a *CustomerApi) OptionForDeliveryOrderFilter() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerOptionForDeliveryOrderFilter),
		func(ctx apiContext) {
			var request dto_request.CustomerOptionForDeliveryOrderFilterRequest
			ctx.mustBind(&request)

			customers, total := a.customerUseCase.OptionForDeliveryOrderFilter(ctx.context(), request)

			nodes := util.ConvertArray(customers, dto_response.NewCustomerResponse)

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

func RegisterCustomerApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CustomerApi{
		api:             newApi(useCaseManager),
		customerUseCase: useCaseManager.CustomerUseCase(),
	}

	routerGroup := router.Group("/customers")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/delivery-order-form", api.OptionForDeliveryOrderForm())
	optionRouterGroup.POST("/delivery-order-filter", api.OptionForDeliveryOrderFilter())
}
