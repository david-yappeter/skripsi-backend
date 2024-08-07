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

type DebtApi struct {
	api
	debtUseCase use_case.DebtUseCase
}

// API:
//
//	@Router		/debts/upload [post]
//	@Summary	Upload Image
//	@tags		Debts
//	@Accept		json
//	@Param		dto_request.DebtUploadImageRequest	body	dto_request.DebtUploadImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *DebtApi) UploadImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDebtUploadImage),
		func(ctx apiContext) {
			var request dto_request.DebtUploadImageRequest
			ctx.mustBind(&request)

			path := a.debtUseCase.UploadImage(ctx.context(), request)

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
//	@Router		/debts/filter [post]
//	@Summary	Filter
//	@tags		Debts
//	@Accept		json
//	@Param		dto_request.DebtFetchRequest	body	dto_request.DebtFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.DebtResponse}}
func (a *DebtApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDebtFetch),
		func(ctx apiContext) {
			var request dto_request.DebtFetchRequest
			ctx.mustBind(&request)

			debts, total := a.debtUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(debts, dto_response.NewDebtResponse)

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
//	@Router		/debts/report [get]
//	@Summary	Download Excel Report
//	@tags		Debts
//	@Accept		json
//	@Param		dto_request.DebtDownloadReportRequest	body	dto_request.DebtDownloadReportRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.DebtResponse}}
func (a *DebtApi) DownloadReport() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDebtDownloadReport),
		func(ctx apiContext) {
			var request dto_request.DebtDownloadReportRequest
			ctx.mustBind(&request)

			ioReadCloser, contentLength, contentType, filename := a.debtUseCase.DownloadReport(ctx.context(), request)

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
//	@Router		/debts/{id} [get]
//	@Summary	Get
//	@tags		Debts
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{debt=dto_response.DebtResponse}}
func (a *DebtApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDebtGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DebtGetRequest
			ctx.mustBind(&request)

			request.DebtId = id

			debt := a.debtUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"debt": dto_response.NewDebtResponse(debt),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/debts/{id}/payment [patch]
//	@Summary	Payment
//	@tags		Debts
//	@Accept		json
//	@Param		id								path	string							true	"Id"
//	@Param		dto_request.DebtPaymentRequest	body	dto_request.DebtPaymentRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{debt=dto_response.DebtResponse}}
func (a *DebtApi) Payment() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDebtPayment),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DebtPaymentRequest
			ctx.mustBind(&request)

			request.DebtId = id

			debt := a.debtUseCase.Payment(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"debt": dto_response.NewDebtResponse(debt),
					},
				},
			)
		},
	)
}

func RegisterDebtApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := DebtApi{
		api:         newApi(useCaseManager),
		debtUseCase: useCaseManager.DebtUseCase(),
	}

	routerGroup := router.Group("/debts")
	routerGroup.POST("/upload", api.UploadImage())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/report", api.DownloadReport())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PATCH("/:id/payment", api.Payment())
}
