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

type RoleApi struct {
	api
	roleUseCase use_case.RoleUseCase
}

// API:
//
//	@Router		/roles/options/user-form [post]
//	@Summary	Option for User Form
//	@tags		Roles
//	@Accept		json
//	@Param		dto_request.RoleOptionForUserFormRequest	body	dto_request.RoleOptionForUserFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.RoleResponse}}
func (a *RoleApi) OptionForUserForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionRoleOptionForUserForm),
		func(ctx apiContext) {
			var request dto_request.RoleOptionForUserFormRequest
			ctx.mustBind(&request)

			roles, total := a.roleUseCase.OptionForUserForm(ctx.context(), request)

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

func RegisterRoleApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := RoleApi{
		api:         newApi(useCaseManager),
		roleUseCase: useCaseManager.RoleUseCase(),
	}

	routerGroup := router.Group("/roles")

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/user-form", api.OptionForUserForm())
}
