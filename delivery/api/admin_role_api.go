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

type AdminRoleApi struct {
	api
	roleUseCase use_case.RoleUseCase
}

// API:
//
//	@Router		/admin/roles/options/user-form [post]
//	@Summary	Option for User Form
//	@tags		Admin Roles
//	@Accept		json
//	@Param		dto_request.AdminRoleOptionForUserFormRequest	body	dto_request.AdminRoleOptionForUserFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.RoleResponse}}
func (a *AdminRoleApi) OptionForUserForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminRoleOptionForUserForm),
		func(ctx apiContext) {
			var request dto_request.AdminRoleOptionForUserFormRequest
			ctx.mustBind(&request)

			roles, total := a.roleUseCase.AdminOptionForUserForm(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: util.ConvertArray(roles, dto_response.NewRoleResponse),
					},
				},
			)
		},
	)
}

func RegisterAdminRoleApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AdminRoleApi{
		api:         newApi(useCaseManager),
		roleUseCase: useCaseManager.RoleUseCase(),
	}

	adminRouterGroup := router.Group("/admin/roles")
	adminRouterGroup.POST("/options/user-form", api.OptionForUserForm())
}
