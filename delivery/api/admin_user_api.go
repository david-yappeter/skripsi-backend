package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminUserApi struct {
	api
	userUseCase use_case.UserUseCase
}

// API:
//
//	@Router		/admin/users [post]
//	@Summary	Create
//	@tags		Admin Users
//	@Accept		json
//	@Param		dto_request.AdminUserCreateRequest	body	dto_request.AdminUserCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserCreate),
		func(ctx apiContext) {
			var request dto_request.AdminUserCreateRequest
			ctx.mustBind(&request)

			user := a.userUseCase.AdminCreate(ctx.context(), request)

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

// API:
//
//	@Router		/admin/users/{id} [put]
//	@Summary	Update
//	@tags		Admin Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdateRequest	body	dto_request.AdminUserUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdateRequest
			ctx.mustBind(&request)
			request.UserId = id

			user := a.userUseCase.AdminUpdate(ctx.context(), request)

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

// API:
//
//	@Router		/admin/users/{id}/password [patch]
//	@Summary	Change Password
//	@tags		Admin Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdatePasswordRequest	body	dto_request.AdminUserUpdatePasswordRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) UpdatePassword() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdatePassword),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdatePasswordRequest
			ctx.mustBind(&request)
			request.UserId = id

			user := a.userUseCase.AdminUpdatePassword(ctx.context(), request)

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

// API:
//
//	@Router		/admin/users/{id}/active [patch]
//	@Summary	Update Active
//	@tags		Admin Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdateActiveRequest	body	dto_request.AdminUserUpdateActiveRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) UpdateActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdateActive),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdateActiveRequest
			request.UserId = id

			user := a.userUseCase.AdminUpdateActive(ctx.context(), request)

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

// API:
//
//	@Router		/admin/users/{id}/inactive [patch]
//	@Summary	Update InActive
//	@tags		Admin Users
//	@Accept		json
//	@Param		dto_request.AdminUserUpdateInActiveRequest	body	dto_request.AdminUserUpdateInActiveRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) UpdateInActive() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserUpdateInActive),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUserUpdateInActiveRequest
			request.UserId = id

			user := a.userUseCase.AdminUpdateInActive(ctx.context(), request)

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

// API:
//
//	@Router		/admin/users/{id}/roles [post]
//	@Summary	Add Role
//	@tags		Admin Users
//	@Accept		json
//	@Param		id									path	string								true	"Id"
//	@Param		dto_request.AdminUserAddRoleRequest	body	dto_request.AdminUserAddRoleRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) AddRole() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserAddRole),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.AdminUserAddRoleRequest
			ctx.mustBind(&request)

			request.UserId = id

			user := a.userUseCase.AdminAddRole(ctx.context(), request)

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

// API:
//
//	@Router		/admin/users/{id}/roles/{role_id} [delete]
//	@Summary	Delete Role
//	@tags		Admin Users
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		role_id									path	string									true	"Role Id"
//	@Param		dto_request.AdminUserDeleteRoleRequest	body	dto_request.AdminUserDeleteRoleRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *AdminUserApi) DeleteRole() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUserDeleteRole),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			roleId := ctx.getUuidParam("role_id")

			var request dto_request.AdminUserDeleteRoleRequest
			ctx.mustBind(&request)

			request.UserId = id
			request.RoleId = roleId

			user := a.userUseCase.AdminDeleteRole(ctx.context(), request)

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
	api := AdminUserApi{
		api:         newApi(useCaseManager),
		userUseCase: useCaseManager.UserUseCase(),
	}

	adminRouterGroup := router.Group("/admin/users")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.PATCH("/:id", api.UpdatePassword())
	adminRouterGroup.PATCH("/:id/active", api.UpdateActive())
	adminRouterGroup.PATCH("/:id/inactive", api.UpdateInActive())

	adminRouterGroup.POST("/:id/roles", api.AddRole())
	adminRouterGroup.DELETE("/:id/roles/:role_id", api.DeleteRole())
}
