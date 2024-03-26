package api

import (
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthApi struct {
	api
	authUseCase use_case.AuthUseCase
}

// API:
//
//	@Router		/auth/login [post]
//	@Summary	Username Login
//	@tags		Auth
//	@Accept		json
//	@Param		dto_request.AuthUsernameLoginRequest	body	dto_request.AuthUsernameLoginRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.AuthTokenResponse}
func (a *AuthApi) LoginUsername() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.AuthUsernameLoginRequest
			ctx.mustBind(&request)

			data := a.authUseCase.LoginUsername(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.NewAuthTokenResponse(data),
				},
			)
		},
	)
}

// API:
//
//	@Router		/auth/login-driver [post]
//	@Summary	Username Login Driver (For Mobile)
//	@tags		Auth
//	@Accept		json
//	@Param		dto_request.AuthUsernameLoginDriverRequest	body	dto_request.AuthUsernameLoginDriverRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.AuthTokenResponse}
func (a *AuthApi) LoginUsernameDriver() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.AuthUsernameLoginDriverRequest
			ctx.mustBind(&request)

			data := a.authUseCase.LoginUsernameDriver(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.NewAuthTokenResponse(data),
				},
			)
		},
	)
}

// API:
//
//	@Router		/auth/logout [post]
//	@Summary	Logout
//	@tags		Auth
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *AuthApi) Logout() gin.HandlerFunc {
	return a.Authorize(
		nil,
		func(ctx apiContext) {
			a.authUseCase.Logout(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterAuthApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AuthApi{
		api:         newApi(useCaseManager),
		authUseCase: useCaseManager.AuthUseCase(),
	}

	routerGroup := router.Group("/auth")
	routerGroup.POST("/login", api.LoginUsername())
	routerGroup.POST("/login-driver", api.LoginUsernameDriver())
	routerGroup.POST("/logout", api.Logout())
}
