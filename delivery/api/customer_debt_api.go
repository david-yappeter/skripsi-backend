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

type CustomerDebtApi struct {
	api
	customerDebtUseCase use_case.CustomerDebtUseCase
}

// API:
//
//	@Router		/customer-debts/upload [post]
//	@Summary	Upload Image
//	@tags		Customer Debts
//	@Accept		json
//	@Param		dto_request.CustomerDebtUploadImageRequest	body	dto_request.CustomerDebtUploadImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *CustomerDebtApi) UploadImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerDebtUploadImage),
		func(ctx apiContext) {
			var request dto_request.CustomerDebtUploadImageRequest
			ctx.mustBind(&request)

			path := a.customerDebtUseCase.UploadImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"path": path,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customer-debts/report [get]
//	@Summary	Download Report
//	@tags		Customer Debts
//	@Accept		json
//	@Param		dto_request.CustomerDebtDownloadReportRequest	body	dto_request.CustomerDebtDownloadReportRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *CustomerDebtApi) DownloadReport() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerDebtDownloadReport),
		func(ctx apiContext) {
			var request dto_request.CustomerDebtDownloadReportRequest
			ctx.mustBind(&request)

			ioReadCloser, contentLength, contentType, filename := a.customerDebtUseCase.DownloadReport(ctx.context(), request)

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
//	@Router		/customer-debts/filter [post]
//	@Summary	Filter
//	@tags		Customer Debts
//	@Accept		json
//	@Param		dto_request.CustomerDebtFetchRequest	body	dto_request.CustomerDebtFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerDebtResponse}}
func (a *CustomerDebtApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerDebtFetch),
		func(ctx apiContext) {
			var request dto_request.CustomerDebtFetchRequest
			ctx.mustBind(&request)

			customerDebts, total := a.customerDebtUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(customerDebts, dto_response.NewCustomerDebtResponse)

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
//	@Router		/customer-debts/{id} [get]
//	@Summary	Get
//	@tags		Customer Debts
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer_debt=dto_response.CustomerDebtResponse}}
func (a *CustomerDebtApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerDebtGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerDebtGetRequest
			ctx.mustBind(&request)

			request.CustomerDebtId = id

			customerDebt := a.customerDebtUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer_debt": dto_response.NewCustomerDebtResponse(customerDebt),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/customer-debts/{id}/payment [patch]
//	@Summary	Payment
//	@tags		Customer Debts
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.CustomerDebtPaymentRequest	body	dto_request.CustomerDebtPaymentRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{customer_debt=dto_response.CustomerDebtResponse}}
func (a *CustomerDebtApi) Payment() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerDebtPayment),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.CustomerDebtPaymentRequest
			ctx.mustBind(&request)

			request.CustomerDebtId = id

			customerDebt := a.customerDebtUseCase.Payment(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"customer_debt": dto_response.NewCustomerDebtResponse(customerDebt),
					},
				},
			)
		},
	)
}

func RegisterCustomerDebtApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CustomerDebtApi{
		api:                 newApi(useCaseManager),
		customerDebtUseCase: useCaseManager.CustomerDebtUseCase(),
	}

	routerGroup := router.Group("/customer-debts")
	routerGroup.POST("/upload", api.UploadImage())
	routerGroup.GET("/report", api.DownloadReport())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PATCH("/:id/payment", api.Payment())
}
