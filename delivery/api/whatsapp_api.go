package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WhatsappApi struct {
	api
	whatsappUseCase use_case.WhatsappUseCase
}

// API:
//
//	@Router		/whatsapp/is-logged-in [get]
//	@Summary	Check whatsapp is logged in
//	@tags		Whatsapps
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{is_logged_in=bool}}
func (a *WhatsappApi) IsLoggedIn() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappIsLoggedIn),
		func(ctx apiContext) {
			isLoggedIn := a.whatsappUseCase.IsLoggedIn(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"is_logged_in": isLoggedIn,
					},
				},
			)
		},
	)

}

// API:
//
//	@Router		/whatsapp/logout [post]
//	@Summary	Logout
//	@tags		Whatsapps
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *WhatsappApi) Logout() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappLogout),
		func(ctx apiContext) {
			a.whatsappUseCase.Logout(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterWhatsappApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := WhatsappApi{
		api:             newApi(useCaseManager),
		whatsappUseCase: useCaseManager.WhatsappUseCase(),
	}

	routerGroup := router.Group("/whatsapp")
	routerGroup.GET("/is-logged-in", api.IsLoggedIn())
	routerGroup.POST("/logout", api.Logout())
}
