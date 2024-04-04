package api

import (
	"encoding/json"
	"fmt"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"time"

	"github.com/gin-gonic/gin"
)

type SsrApi struct {
	api
	deliveryOrderUseCase use_case.DeliveryOrderUseCase
}

// API:
//
//	@Router		/ssr/maps/{delivery_order_id} [get]
//	@Summary	Get SSR Maps Data
//	@tags		Suppliers
//	@Accept		json
//	@Param		delivery_order_id	path	string	true	"Delivery Order Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order_position=dto_response.DeliveryOrderPositionResponse}}
func (a *SsrApi) Ssr() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			id := ctx.getUuidParam("delivery_order_id")
			var request dto_request.LatestDeliveryLocationRequest
			ctx.mustBind(&request)

			request.DeliveryOrderId = id

			// set response headers for SSE
			ctx.ginCtx.Header("Content-Type", "text/event-stream")
			ctx.ginCtx.Header("Cache-Control", "no-cache")
			ctx.ginCtx.Header("Connection", "keep-alive")
			ctx.ginCtx.Header("X-Accel-Buffering", "no")

			// create a channel to receive notifications when the client connection is closed
			closeNotify := ctx.ginCtx.Writer.CloseNotify()

			// tick every 3 second
			ticker := time.NewTicker(3 * time.Second)
			defer ticker.Stop()

			for {
				select {
				case <-closeNotify:
					// Client connection closed, exit the handler
					return
				case <-ticker.C:
					deliveryOrderPosition := a.deliveryOrderUseCase.LatestDeliveryLocation(ctx.context(), request)

					var resp *dto_response.DeliveryOrderPositionResponse
					if deliveryOrderPosition != nil {
						resp = dto_response.NewDeliveryOrderPositionResponseP(*deliveryOrderPosition)
					}

					jsonData, err := json.Marshal(dto_response.Response{
						Data: dto_response.DataResponse{
							"delivery_order_position": resp,
						},
					})
					if err != nil {
						fmt.Println("Error encoding JSON:", err)
						return
					}
					ctx.ginCtx.Writer.WriteString(string(jsonData))
					ctx.ginCtx.Writer.Flush()
				}
			}
		},
	)
}

func RegisterSsrApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := SsrApi{
		api:                  newApi(useCaseManager),
		deliveryOrderUseCase: useCaseManager.DeliveryOrderUseCase(),
	}

	router.GET("/ssr/maps/:delivery_order_id", api.Ssr())
}
