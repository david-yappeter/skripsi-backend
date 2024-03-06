package api

import (
	"fmt"
	"io/ioutil"
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

			reqBody, err := ioutil.ReadAll(ctx.ginCtx.Request.Body)
			if err != nil {
				panic(err)
			}

			fmt.Printf("AAAA \n%+v\n\n", string(reqBody))

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
