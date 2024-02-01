package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CashierSessionApi struct {
	api
	cashierSessionUseCase use_case.CashierSessionUseCase
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

func RegisterCashierSessionApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CashierSessionApi{
		api:                   newApi(useCaseManager),
		cashierSessionUseCase: useCaseManager.CashierSessionUseCase(),
	}

	routerGroup := router.Group("/cashier-sessions")
	routerGroup.POST("/start", api.Start())
}
