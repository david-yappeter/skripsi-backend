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

type BalanceApi struct {
	api
	balanceUseCase use_case.BalanceUseCase
}

// API:
//
//	@Router		/balances [post]
//	@Summary	Create
//	@tags		Balances
//	@Accept		json
//	@Param		dto_request.BalanceCreateRequest	body	dto_request.BalanceCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{balance=dto_response.BalanceResponse}}
func (a *BalanceApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionBalanceCreate),
		func(ctx apiContext) {
			var request dto_request.BalanceCreateRequest
			ctx.mustBind(&request)

			balance := a.balanceUseCase.Create(ctx.context(), request)

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

//	@Router		/balances/filter [post]
//	@Summary	Filter
//	@tags		Balances
//	@Accept		json
//	@Param		dto_request.BalanceFetchRequest	body	dto_request.BalanceFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.BalanceResponse}}
func (a *BalanceApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionBalanceCreate),
		func(ctx apiContext) {
			var request dto_request.BalanceFetchRequest
			ctx.mustBind(&request)

			balances, total := a.balanceUseCase.Fetch(ctx.context(), request)

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

//	@Router		/balances/{id} [get]
//	@Summary	Get
//	@tags		Balances
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{balance=dto_response.BalanceResponse}}
func (a *BalanceApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionBalanceGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.BalanceGetRequest
			ctx.mustBind(&request)

			request.BalanceId = id

			balance := a.balanceUseCase.Get(ctx.context(), request)

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

//	@Router		/balances/{id} [put]
//	@Summary	Update
//	@tags		Balances
//	@Accept		json
//	@Param		id									path	string								true	"Id"
//	@Param		dto_request.BalanceUpdateRequest	body	dto_request.BalanceUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{balance=dto_response.BalanceResponse}}
func (a *BalanceApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionBalanceUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.BalanceUpdateRequest
			ctx.mustBind(&request)

			request.BalanceId = id

			balance := a.balanceUseCase.Update(ctx.context(), request)

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

//	@Router		/balances/{id} [delete]
//	@Summary	Delete
//	@tags		Balances
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *BalanceApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionBalanceDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.BalanceDeleteRequest
			ctx.mustBind(&request)
			request.BalanceId = id

			a.balanceUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterBalanceApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := BalanceApi{
		api:            newApi(useCaseManager),
		balanceUseCase: useCaseManager.BalanceUseCase(),
	}

	routerGroup := router.Group("/balances")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())
}
