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

type DashboardApi struct {
	api
	dashboardUseCase use_case.DashboardUseCase
}

// API:
//
//	@Router		/dashboards/summarize-debt [post]
//	@Summary	Summarize Debt
//	@tags		Dashboards
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer_debt_summaries=[]dto_response.CustomerDebtSummaryResponse,supplier_debt_summaries=[]dto_response.SupplierDebtSummaryResponse}}
func (a *DashboardApi) SummarizeDebt() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDashboardSummarizeDebt),
		func(ctx apiContext) {
			customerDebtSummaries, supplierDebtSummaries := a.dashboardUseCase.SummarizeDebt(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer_debt_summaries": util.ConvertArray(customerDebtSummaries, dto_response.NewCustomerDebtSummaryResponse),
						"supplier_debt_summaries": util.ConvertArray(supplierDebtSummaries, dto_response.NewSupplierDebtSummaryResponse),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/dashboards/summarize-transaction [post]
//	@Summary	Summarize Transactions
//	@tags		Dashboards
//	@Accept		json
//	@Param		dto_request.DashboardSummarizeTransactionRequest	body	dto_request.DashboardSummarizeTransactionRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{transaction_summaries=[]dto_response.TransactionSummaryResponse}}
func (a *DashboardApi) SummarizeTransaction() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDashboardSummarizeTransaction),
		func(ctx apiContext) {
			var request dto_request.DashboardSummarizeTransactionRequest
			ctx.mustBind(&request)

			transactionSummaries := a.dashboardUseCase.SummarizeTransaction(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"transaction_summaries": util.ConvertArray(transactionSummaries, dto_response.NewTransactionSummaryResponse),
					},
				},
			)
		},
	)
}

func RegisterDashboardApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := DashboardApi{
		api:              newApi(useCaseManager),
		dashboardUseCase: useCaseManager.DashboardUseCase(),
	}

	routerGroup := router.Group("/dashboards")
	routerGroup.POST("/summarize-debt", api.SummarizeDebt())
	routerGroup.POST("/summarize-transaction", api.SummarizeTransaction())
}
