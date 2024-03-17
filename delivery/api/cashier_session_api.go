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

type CashierSessionApi struct {
	api
	cashierSessionUseCase use_case.CashierSessionUseCase
}

// API:
//
//	@Router		/cashier-sessions/filter [post]
//	@Summary	Fetch Cashier Sessions
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		dto_request.CashierSessionFetchRequest	body	dto_request.CashierSessionFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CashierSessionResponse}}
func (a *CashierSessionApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionFetch),
		func(ctx apiContext) {
			var request dto_request.CashierSessionFetchRequest
			ctx.mustBind(&request)

			cashierSessions, total := a.cashierSessionUseCase.Fetch(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Limit: request.Limit,
						Page:  request.Page,
						Total: total,
						Nodes: util.ConvertArray(cashierSessions, dto_response.NewCashierSessionResponse),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/cashier-sessions/start [post]
//	@Summary	Start Cashier Session
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		dto_request.CashierSessionStartRequest	body	dto_request.CashierSessionStartRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cashier_session=dto_response.CashierSessionResponse}}
func (a *CashierSessionApi) Start() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionStart),
		func(ctx apiContext) {
			var request dto_request.CashierSessionStartRequest
			ctx.mustBind(&request)

			cashierSession := a.cashierSessionUseCase.Start(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cashier_session": dto_response.NewCashierSessionResponse(cashierSession),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/cashier-sessions/get-current [get]
//	@Summary	Get current user cashier session
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		dto_request.CashierSessionStartRequest	body	dto_request.CashierSessionStartRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cashier_session=dto_response.CashierSessionResponse}}
func (a *CashierSessionApi) GetCurrent() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionGetCurrent),
		func(ctx apiContext) {
			var request dto_request.CashierSessionStartRequest
			ctx.mustBind(&request)

			cashierSession := a.cashierSessionUseCase.GetByCurrentUser(ctx.context())

			var resp *dto_response.CashierSessionResponse
			if cashierSession != nil {
				resp = dto_response.NewCashierSessionResponseP(*cashierSession)
			}

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cashier_session": resp,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/cashier-sessions/end [post]
//	@Summary	End Cashier Session
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		dto_request.CashierSessionEndRequest	body	dto_request.CashierSessionEndRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cashier_session=dto_response.CashierSessionResponse}}
func (a *CashierSessionApi) End() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionEnd),
		func(ctx apiContext) {
			var request dto_request.CashierSessionEndRequest
			ctx.mustBind(&request)

			cashierSession := a.cashierSessionUseCase.End(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"cashier_session": dto_response.NewCashierSessionResponse(cashierSession),
					},
				},
			)
		},
	)
}

func RegisterCashierSessionApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CashierSessionApi{
		api:                   newApi(useCaseManager),
		cashierSessionUseCase: useCaseManager.CashierSessionUseCase(),
	}

	routerGroup := router.Group("/cashier-sessions")
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.POST("/start", api.Start())
	routerGroup.GET("/current", api.GetCurrent())
	routerGroup.POST("/end", api.End())
}
