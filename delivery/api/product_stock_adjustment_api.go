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

type ProductStockAdjustmentApi struct {
	api
	productStockAdjustmentUseCase use_case.ProductStockAdjustmentUseCase
}

// API:
//
//	@Router		/product-stock-adjustments [post]
//	@Summary	Fetch
//	@tags		Product Stock Adjustments
//	@Accept		json
//	@Param		dto_request.ProductStockAdjustmentFetchRequest	body	dto_request.ProductStockAdjustmentFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductStockAdjustmentResponse,printer_data=[]int16}}
func (a *ProductStockAdjustmentApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductStockAdjustmentFetch),
		func(ctx apiContext) {
			var request dto_request.ProductStockAdjustmentFetchRequest
			ctx.mustBind(&request)

			productStockAdjustments, total := a.productStockAdjustmentUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(productStockAdjustments, dto_response.NewProductStockAdjustmentResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Nodes: nodes,
						Total: total,
						Page:  request.Page,
						Limit: request.Limit,
					},
				},
			)
		},
	)
}

func RegisterProductStockAdjustmentApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductStockAdjustmentApi{
		api:                           newApi(useCaseManager),
		productStockAdjustmentUseCase: useCaseManager.ProductStockAdjustmentUseCase(),
	}

	routerGroup := router.Group("/product-stock-adjustments")
	routerGroup.POST("", api.Fetch())
}
