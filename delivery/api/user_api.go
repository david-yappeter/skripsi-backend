package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"myapp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserApi struct {
	api
	userUseCase use_case.UserUseCase
}

// API:
//
//	@Router		/users [post]
//	@Summary	Create
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.UserCreateRequest	body	dto_request.UserCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserCreate),
		func(ctx apiContext) {
			var request dto_request.UserCreateRequest
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

// API:
//
//	@Router		/users/filter [post]
//	@Summary	Fetch
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.UserFetchRequest	body	dto_request.UserFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.UserResponse}}
func (a *UserApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserFetch),
		func(ctx apiContext) {
			var request dto_request.UserFetchRequest
			ctx.mustBind(&request)

			users, total := a.userUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(users, dto_response.NewUserResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Nodes: nodes,
						Total: total,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/users/me [post]
//	@Summary	Get Me
//	@tags		Users
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserMeResponse}}
func (a *UserApi) GetMe() gin.HandlerFunc {
	return a.Authorize(
		nil,
		func(ctx apiContext) {

			user := a.userUseCase.GetMe(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"user": dto_response.NewUserMeResponse(user),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/users/{id} [get]
//	@Summary	Get
//	@tags		Users
//	@Accept		json
//	@Param		id	path	string	true	"User Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) Get() gin.HandlerFunc {
	return a.Authorize(
		nil,
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.UserGetRequest
			request.UserId = id

			user := a.userUseCase.Get(ctx.context(), request)

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
//	@Router		/users/{id} [put]
//	@Summary	Update
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.UserUpdateRequest	body	dto_request.UserUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.UserUpdateRequest
			ctx.mustBind(&request)
			request.UserId = id

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

// API:
//
//	@Router		/users/{id}/password [patch]
//	@Summary	Change Password
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.UserUpdatePasswordRequest	body	dto_request.UserUpdatePasswordRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) UpdatePassword() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserUpdatePassword),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.UserUpdatePasswordRequest
			ctx.mustBind(&request)
			request.UserId = id

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

// API:
//
//	@Router		/users/{id}/roles [post]
//	@Summary	Add Role
//	@tags		Users
//	@Accept		json
//	@Param		id								path	string							true	"Id"
//	@Param		dto_request.UserAddRoleRequest	body	dto_request.UserAddRoleRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) AddRole() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserAddRole),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")

			var request dto_request.UserAddRoleRequest
			ctx.mustBind(&request)

			request.UserId = id

			user := a.userUseCase.AddRole(ctx.context(), request)

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
//	@Router		/users/{id}/roles/{role_id} [delete]
//	@Summary	Delete Role
//	@tags		Users
//	@Accept		json
//	@Param		id		path	string	true	"Id"
//	@Param		role_id	path	string	true	"Role Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{user=dto_response.UserResponse}}
func (a *UserApi) DeleteRole() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserDeleteRole),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			roleId := ctx.getUuidParam("role_id")

			var request dto_request.UserDeleteRoleRequest

			request.UserId = id
			request.RoleId = roleId

			user := a.userUseCase.DeleteRole(ctx.context(), request)

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
//	@Router		/users/options/cashier-session-filter [post]
//	@Summary	Option for Cashier Session Filter
//	@tags		Users
//	@Accept		json
//	@Param		dto_request.UserOptionForCashierSessionFilterRequest	body	dto_request.UserOptionForCashierSessionFilterRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.UserResponse}}
func (a *UserApi) OptionForCashierSessionFilter() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUserOptionForCashierSessionFilter),
		func(ctx apiContext) {
			var request dto_request.UserOptionForCashierSessionFilterRequest
			ctx.mustBind(&request)

			users, total := a.userUseCase.OptionForCashierSessionFilter(ctx.context(), request)

			nodes := util.ConvertArray(users, dto_response.NewUserResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Nodes: nodes,
						Total: total,
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
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.POST("/me", api.GetMe())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.PATCH("/:id", api.UpdatePassword())

	routerGroup.POST("/:id/roles", api.AddRole())
	routerGroup.DELETE("/:id/roles/:role_id", api.DeleteRole())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/cashier-session-filter", api.OptionForCashierSessionFilter())
}
