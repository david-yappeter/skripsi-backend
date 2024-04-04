package api

import (
	"fmt"
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
//	@Router		/cashier-sessions/{id} [get]
//	@Summary	Get
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		id	path	string	true	"Cashier Session Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cashier_session=dto_response.CashierSessionResponse}}
func (a *CashierSessionApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CashierSessionGetRequest
			request.CashierSessionId = id

			cashierSession := a.cashierSessionUseCase.Get(ctx.context(), request)

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
//	@Router		/cashier-sessions/{id}/transactions [post]
//	@Summary	Fetch Transaction
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		id													path	string												true	"Cashier Session Id"
//	@Param		dto_request.CashierSessionFetchTransactionRequest	body	dto_request.CashierSessionFetchTransactionRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.TransactionResponse}}
func (a *CashierSessionApi) FetchTransaction() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionFetchTransaction),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CashierSessionFetchTransactionRequest
			request.CashierSessionId = id

			transactions, total := a.cashierSessionUseCase.FetchTransaction(ctx.context(), request)

			nodes := util.ConvertArray(transactions, dto_response.NewTransactionResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Limit: request.Limit,
						Page:  request.Page,
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
//	@Router		/cashier-sessions/{id}/report [get]
//	@Summary	Download Report
//	@tags		Cashier Sessions
//	@Accept		json
//	@Param		id	path	string	true	"Cashier Session Id"
//	@Produce	json
func (a *CashierSessionApi) DownloadReport() gin.HandlerFunc {
	return a.Guest(
		// data_type.PermissionP(data_type.PermissionCashierSessionDownloadReport),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CashierSessionDownloadReportRequest
			request.CashierSessionId = id

			ioReadCloser, contentLength, contentType, filename := a.cashierSessionUseCase.DownloadReport(ctx.context(), request)

			ctx.dataFromReader(
				http.StatusOK,
				contentLength,
				contentType,
				ioReadCloser,
				map[string]string{
					"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
				},
			)
		},
	)
}

// API:
//
//	@Router		/cashier-sessions/current [get]
//	@Summary	Get current user cashier session
//	@tags		Cashier Sessions
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{cashier_session=dto_response.CashierSessionResponse}}
func (a *CashierSessionApi) GetCurrent() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCashierSessionGetCurrent),
		func(ctx apiContext) {
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
	routerGroup.GET("/:id", api.Get())
	routerGroup.POST("/:id/transactions", api.FetchTransaction())
	routerGroup.GET("/:id/report", api.DownloadReport())
	routerGroup.GET("/current", api.GetCurrent())
	routerGroup.POST("/end", api.End())
}
