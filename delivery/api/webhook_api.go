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

func RegisterWebhookApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := WebhookApi{
		api:            newApi(useCaseManager),
		webhookUseCase: useCaseManager.WebhookUseCase(),
	}

	router.POST("/webhook", api.Webhook())
}
