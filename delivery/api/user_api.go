package api

import (
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	api
	userUseCase use_case.UserUseCase
}

// API:
//
//	@Router		/users/me [post]
//	@Summary	Get Me
//	@tags		Users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) GetMe() gin.HandlerFunc {
	return a.Authorize(
		nil,
		func(ctx apiContext) {

			user := a.userUseCase.GetMe(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"user": dto_response.NewUserResponse(user),
					},
				},
			)
		},
	)
}

func RegisterUserApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := UserApi{
		api:         newApi(useCaseManager),
		userUseCase: useCaseManager.UserUseCase(),
	}

	routerGroup := router.Group("/users")
	routerGroup.POST("/me", api.GetMe())
}
