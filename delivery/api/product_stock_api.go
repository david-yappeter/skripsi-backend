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

type ProductStockApi struct {
	api
	productStockUseCase use_case.ProductStockUseCase
}

// API:
//
//	@Router		/product-stocks/filter [post]
//	@Summary	Filter
//	@tags		Product Stocks
//	@Accept		json
//	@Param		dto_request.ProductStockFetchRequest	body	dto_request.ProductStockFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductStockResponse}}
func (a *ProductStockApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductStockFetch),
		func(ctx apiContext) {
			var request dto_request.ProductStockFetchRequest
			ctx.mustBind(&request)

			productStocks, total := a.productStockUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(productStocks, dto_response.NewProductStockResponse)

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
//	@Router		/product-stocks/{id} [get]
//	@Summary	Get
//	@tags		Product Stocks
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_stock=dto_response.ProductStockResponse}}
func (a *ProductStockApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductStockGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductStockGetRequest
			ctx.mustBind(&request)

			request.ProductStockId = id

			productStock := a.productStockUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_stock": dto_response.NewProductStockResponse(productStock),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-stocks/report [get]
//	@Summary	Download Report
//	@tags		Product Stocks
//	@Produce	json
func (a *ProductStockApi) DownloadReport() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductStockGet),
		func(ctx apiContext) {

			ioReadCloser, contentLength, contentType, filename := a.productStockUseCase.DownloadReport(ctx.context())

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
//	@Router		/product-stocks/{id}/adjustment [patch]
//	@Summary	Adjustment
//	@tags		Product Stocks
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.ProductStockAdjustmentRequest	body	dto_request.ProductStockAdjustmentRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_stock=dto_response.ProductStockResponse}}
func (a *ProductStockApi) Adjustment() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductStockAdjustment),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductStockAdjustmentRequest
			ctx.mustBind(&request)

			request.ProductStockId = id

			productStock := a.productStockUseCase.Adjustment(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_stock": dto_response.NewProductStockResponse(productStock),
					},
				},
			)
		},
	)
}

func RegisterProductStockApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductStockApi{
		api:                 newApi(useCaseManager),
		productStockUseCase: useCaseManager.ProductStockUseCase(),
	}

	routerGroup := router.Group("/product-stocks")
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.GET("/report", api.DownloadReport())
	routerGroup.PATCH("/:id/adjustment", api.Adjustment())
}
