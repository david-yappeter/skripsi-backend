package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	api
	userUseCase use_case.UserUseCase
}

//	@Router		/admin/users [post]
//	@Summary	Create
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.AdminUserCreateRequest	body	dto_request.AdminUserCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserCreate),
		func(ctx apiContext) {
			var request dto_request.AdminUserCreateRequest
			ctx.mustBind(&request)

			user := a.userUseCase.Create(ctx.context(), request)

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

//	@Router		/admin/users/{id} [put]
//	@Summary	Update
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdateRequest	body	dto_request.AdminUserUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdateRequest
			ctx.mustBind(&request)
			request.Id = id

			user := a.userUseCase.Update(ctx.context(), request)

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

//	@Router		/admin/users/{id}/password [patch]
//	@Summary	Update Password
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdatePasswordRequest	body	dto_request.AdminUserUpdatePasswordRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) UpdatePassword() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdatePassword),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdatePasswordRequest
			ctx.mustBind(&request)
			request.Id = id

			user := a.userUseCase.UpdatePassword(ctx.context(), request)

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

//	@Router		/admin/users/{id}/active [patch]
//	@Summary	Update Active
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdateActiveRequest	body	dto_request.AdminUserUpdateActiveRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) UpdateActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdateActive),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdateActiveRequest
			request.Id = id

			user := a.userUseCase.UpdateActive(ctx.context(), request)

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

//	@Router		/admin/users/{id}/inactive [patch]
//	@Summary	Update InActive
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdateInActiveRequest	body	dto_request.AdminUserUpdateInActiveRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) UpdateInActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdateInActive),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdateInActiveRequest
			request.Id = id

			user := a.userUseCase.UpdateInActive(ctx.context(), request)

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

func RegisterAdminUserApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := UserApi{
		api:         newApi(useCaseManager),
		userUseCase: useCaseManager.UserUseCase(),
	}

	adminRouterGroup := router.Group("/admin/users")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.PATCH("/:id", api.UpdatePassword())
	adminRouterGroup.PATCH("/:id/active", api.UpdateActive())
	adminRouterGroup.PATCH("/:id/inactive", api.UpdateInActive())
}
