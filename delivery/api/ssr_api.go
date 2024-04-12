package api

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"image/png"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"time"

	"github.com/boombuler/barcode"
	"github.com/boombuler/barcode/qr"
	"github.com/gin-gonic/gin"
)

type SsrApi struct {
	api
	deliveryOrderUseCase use_case.DeliveryOrderUseCase
	whatsappUseCase      use_case.WhatsappUseCase
}

// API:
//
//	@Router		/ssr/maps/{delivery_order_id} [get]
//	@Summary	Get SSR Maps Data
//	@tags		Server Sent Event
//	@Accept		json
//	@Param		delivery_order_id	path	string	true	"Delivery Order Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order_position=dto_response.DeliveryOrderPositionResponse}}
func (a *SsrApi) SsrMaps() gin.HandlerFunc {
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

					// this format 'data: ${JSON_DATA}\n\n" is important
					event := fmt.Sprintf("data: %s\n\n", string(jsonData))
					ctx.ginCtx.Writer.WriteString(string(event))
					ctx.ginCtx.Writer.Flush()
				}
			}
		},
	)

}

// API:
//
//	@Router		/ssr/whatsapp/login [get]
//	@Summary	Get SSR Whatsapp Login Base64 QR Code
//	@tags		Server Sent Event
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{base64_qr=string}}
func (a *SsrApi) SsrWhatsappLogin() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionSsrWhatsappLogin),
		func(ctx apiContext) {
			// set response headers for SSE
			ctx.ginCtx.Header("Content-Type", "text/event-stream")
			ctx.ginCtx.Header("Cache-Control", "no-cache")
			ctx.ginCtx.Header("Connection", "keep-alive")
			ctx.ginCtx.Header("X-Accel-Buffering", "no")

			// create a channel to receive notifications when the client connection is closed
			closeNotify := ctx.ginCtx.Writer.CloseNotify()

			qrChan := a.whatsappUseCase.Login(ctx.context())

			for {
				select {
				case <-closeNotify:
					// Client connection closed, exit the handler
					return
				case qrString := <-qrChan:
					if qrString == "" {
						// disconnect
						return
					}

					// Create the barcode
					qrCode, _ := qr.Encode(qrString, qr.L, qr.Auto)

					// Scale the barcode to 400 x 400 pixels
					qrCode, _ = barcode.Scale(qrCode, 400, 400)

					buf := new(bytes.Buffer)

					png.Encode(buf, qrCode)

					base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

					jsonData, _ := json.Marshal(dto_response.Response{
						Data: dto_response.DataResponse{
							"base64_qr": base64Str,
						},
					})

					// this format 'data: ${JSON_DATA}\n\n" is important
					event := fmt.Sprintf("data: %s\n\n", string(jsonData))
					ctx.ginCtx.Writer.WriteString(string(event))
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
		whatsappUseCase:      useCaseManager.WhatsappUseCase(),
	}

	router.GET("/ssr/maps/:delivery_order_id", api.SsrMaps())
	router.GET("/ssr/whatsapp/login", api.SsrWhatsappLogin())
}
