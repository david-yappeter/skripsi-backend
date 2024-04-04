package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TransactionApi struct {
	api
	transactionUseCase use_case.TransactionUseCase
}

// API:
//
//	@Router		/transactions/checkout [post]
//	@Summary	Checkout current active cart
//	@tags		Transactions
//	@Accept		json
//	@Param		dto_request.TransactionCheckoutCartRequest	body	dto_request.TransactionCheckoutCartRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{transaction=dto_response.TransactionResponse,printer_data=[]int16}}
func (a *TransactionApi) Checkout() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTransactionCheckoutCart),
		func(ctx apiContext) {
			var request dto_request.TransactionCheckoutCartRequest
			ctx.mustBind(&request)

			transaction, printerData := a.transactionUseCase.CheckoutCart(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"transaction":  dto_response.NewTransactionResponse(transaction),
						"printer_data": printerData,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/transactions/{id} [get]
//	@Summary	Get Transaction
//	@tags		Transactions
//	@Accept		json
//	@Param		id											path	string										true	"Transaction Id"
//	@Param		dto_request.TransactionCheckoutCartRequest	body	dto_request.TransactionCheckoutCartRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{transaction=dto_response.TransactionResponse}}
func (a *TransactionApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTransactionGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.TransactionGetRequest
			ctx.mustBind(&request)
			request.TransactionId = id

			transaction := a.transactionUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"transaction": dto_response.NewTransactionResponse(transaction),
					},
				},
			)
		},
	)
}

func RegisterTransactionApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := TransactionApi{
		api:                newApi(useCaseManager),
		transactionUseCase: useCaseManager.TransactionUseCase(),
	}

	routerGroup := router.Group("/transactions")
	routerGroup.POST("/checkout", api.Checkout())
	routerGroup.GET("/:id", api.Get())
}
