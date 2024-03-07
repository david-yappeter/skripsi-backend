package api

import (
	"fmt"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WebhookApi struct {
	api
	authUseCase use_case.AuthUseCase
}

func (a *WebhookApi) Webhook() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.TiktokWebhookBaseRequest[dto_request.WebhookOrderStatusChangeRequest]
			ctx.mustBind(&request)

			fmt.Printf("AAAA \n%+v\n\n", request)

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
		api:         newApi(useCaseManager),
		authUseCase: useCaseManager.AuthUseCase(),
	}

	router.POST("/webhook", api.Webhook())
}
