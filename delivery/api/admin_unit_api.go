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

type AdminUnitApi struct {
	api
	unitUseCase use_case.UnitUseCase
}

//	@Router		/admin/units [post]
//	@Summary	Create
//	@tags		Admin Units
//	@Accept		json
//	@Param		dto_request.AdminUnitCreateRequest	body	dto_request.AdminUnitCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{unit=dto_response.UnitResponse}}
func (a *AdminUnitApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUnitCreate),
		func(ctx apiContext) {
			var request dto_request.AdminUnitCreateRequest
			ctx.mustBind(&request)

			unit := a.unitUseCase.AdminCreate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"unit": dto_response.NewUnitResponse(unit),
					},
				},
			)
		},
	)
}

//	@Router		/admin/units/filter [post]
//	@Summary	Filter
//	@tags		Admin Units
//	@Accept		json
//	@Param		dto_request.AdminUnitFetchRequest	body	dto_request.AdminUnitFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.UnitResponse}}
func (a *AdminUnitApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUnitCreate),
		func(ctx apiContext) {
			var request dto_request.AdminUnitFetchRequest
			ctx.mustBind(&request)

			units, total := a.unitUseCase.AdminFetch(ctx.context(), request)

			nodes := util.ConvertArray(units, dto_response.NewUnitResponse)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: nodes,
					},
				},
			)
		},
	)
}

//	@Router		/admin/units/{id} [get]
//	@Summary	Update
//	@tags		Admin Units
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{unit=dto_response.UnitResponse}}
func (a *AdminUnitApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUnitGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUnitGetRequest
			ctx.mustBind(&request)

			request.UnitId = id

			unit := a.unitUseCase.AdminGet(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"unit": dto_response.NewUnitResponse(unit),
					},
				},
			)
		},
	)
}

//	@Router		/admin/units/{id} [put]
//	@Summary	Update
//	@tags		Admin Units
//	@Accept		json
//	@Param		id									path	string								true	"Id"
//	@Param		dto_request.AdminUnitUpdateRequest	body	dto_request.AdminUnitUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{unit=dto_response.UnitResponse}}
func (a *AdminUnitApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUnitUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUnitUpdateRequest
			ctx.mustBind(&request)

			request.UnitId = id

			unit := a.unitUseCase.AdminUpdate(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"unit": dto_response.NewUnitResponse(unit),
					},
				},
			)
		},
	)
}

//	@Router		/admin/units/{id} [delete]
//	@Summary	Update Password
//	@tags		Admin Units
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *AdminUnitApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionAdminUnitDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.AdminUnitDeleteRequest
			ctx.mustBind(&request)
			request.UnitId = id

			a.unitUseCase.AdminDelete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterAdminUnitApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := AdminUnitApi{
		api:         newApi(useCaseManager),
		unitUseCase: useCaseManager.UnitUseCase(),
	}

	adminRouterGroup := router.Group("/admin/units")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.PUT("/:id", api.Update())
	adminRouterGroup.DELETE("/:id", api.Delete())
}
