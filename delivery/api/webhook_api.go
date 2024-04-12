package api

import (
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookApi struct {
	api
	webhookUseCase use_case.WebhookUseCase
}

func (a *WebhookApi) Webhook() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest]
			ctx.mustBind(&request)

			a.webhookUseCase.OrderStatusChange(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// func (a *WebhookApi) WhatsappLogin() gin.HandlerFunc {
// 	return a.Guest(
// 		func(ctx apiContext) {
// 			// var request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest]
// 			// ctx.mustBind(&request)

// 			qrString := a.webhookUseCase.WhatsappLogin(ctx.context())

// 			// Create the barcode
// 			qrCode, _ := qr.Encode(qrString, qr.L, qr.Auto)

// 			// Scale the barcode to 200x200 pixels
// 			qrCode, _ = barcode.Scale(qrCode, 200, 200)

// 			// create the output file
// 			// file, _ := os.Create("qrcode.png")
// 			// defer file.Close()
// 			buf := new(bytes.Buffer)

// 			// qrterminal.GenerateHalfBlock(qrString, qrterminal.L, os.Stdout)

// 			// encode the barcode as png
// 			png.Encode(buf, qrCode)

// 			base64Str := base64.StdEncoding.EncodeToString(buf.Bytes())

// 			ctx.json(
// 				http.StatusOK,
// 				dto_response.Response{
// 					Data: dto_response.DataResponse{
// 						"qr_base64": base64Str,
// 					},
// 				},
// 			)

// 			// reader := bytes.NewReader(buf.Bytes())

// 			// readSeekCloser := util.ReadSeekNopCloser(reader)

// 			// seeker, ok := readSeekCloser.(io.Seeker)
// 			// if !ok {
// 			// 	panic("does not support seeking")
// 			// }

// 			// contentLength, err := seeker.Seek(0, io.SeekEnd)
// 			// if err != nil {
// 			// 	panic(err)
// 			// }
// 			// _, err = seeker.Seek(0, io.SeekStart)
// 			// if err != nil {
// 			// 	panic(err)
// 			// }

// 			// ctx.dataFromReader(
// 			// 	http.StatusOK,
// 			// 	contentLength,
// 			// 	"image/png",
// 			// 	readSeekCloser,
// 			// 	map[string]string{
// 			// 		"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", "qr.png"),
// 			// 	},
// 			// )
// 		},
// 	)
// }

func RegisterWebhookApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := WebhookApi{
		api:            newApi(useCaseManager),
		webhookUseCase: useCaseManager.WebhookUseCase(),
	}

	router.POST("/webhook", api.Webhook())
}
