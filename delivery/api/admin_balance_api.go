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

type AdminBalanceApi struct {
	api
	balanceUseCase use_case.BalanceUseCase
}

// API:
//
//	@Router		/admin/balances [post]
//	@Summary	Create
//	@tags		Admin Balances
//	@Accept		json
//	@Param		dto_request.AdminBalanceCreateRequest	body	dto_request.AdminBalanceCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{balance=dto_response.BalanceResponse}}
func (a *AdminBalanceApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminBalanceCreate),
		func(ctx apiContext) {
			var request dto_request.AdminBalanceCreateRequest
			ctx.mustBind(&request)

			balance := a.balanceUseCase.AdminCreate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"balance": dto_response.NewBalanceResponse(balance),
					},
				},
			)
		},
	)
}

//	@Router		/admin/balances/filter [post]
//	@Summary	Filter
//	@tags		Admin Balances
//	@Accept		json
//	@Param		dto_request.AdminBalanceFetchRequest	body	dto_request.AdminBalanceFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.BalanceResponse}}
func (a *AdminBalanceApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminBalanceCreate),
		func(ctx apiContext) {
			var request dto_request.AdminBalanceFetchRequest
			ctx.mustBind(&request)

			balances, total := a.balanceUseCase.AdminFetch(ctx.context(), request)

			nodes := util.ConvertArray(balances, dto_response.NewBalanceResponse)

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

//	@Router		/admin/balances/{id} [get]
//	@Summary	Get
//	@tags		Admin Balances
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{balance=dto_response.BalanceResponse}}
func (a *AdminBalanceApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminBalanceGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminBalanceGetRequest
			ctx.mustBind(&request)

			request.BalanceId = id

			balance := a.balanceUseCase.AdminGet(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"balance": dto_response.NewBalanceResponse(balance),
					},
				},
			)
		},
	)
}

//	@Router		/admin/balances/{id} [put]
//	@Summary	Update
//	@tags		Admin Balances
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.AdminBalanceUpdateRequest	body	dto_request.AdminBalanceUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{balance=dto_response.BalanceResponse}}
func (a *AdminBalanceApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminBalanceUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminBalanceUpdateRequest
			ctx.mustBind(&request)

			request.BalanceId = id

			balance := a.balanceUseCase.AdminUpdate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"balance": dto_response.NewBalanceResponse(balance),
					},
				},
			)
		},
	)
}

//	@Router		/admin/balances/{id} [delete]
//	@Summary	Delete
//	@tags		Admin Balances
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *AdminBalanceApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminBalanceDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminBalanceDeleteRequest
			ctx.mustBind(&request)
			request.BalanceId = id

			a.balanceUseCase.AdminDelete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterAdminBalanceApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AdminBalanceApi{
		api:            newApi(useCaseManager),
		balanceUseCase: useCaseManager.BalanceUseCase(),
	}

	adminRouterGroup := router.Group("/admin/balances")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.DELETE("/:id", api.Delete())
}
