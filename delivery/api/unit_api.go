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

type UnitApi struct {
	api
	unitUseCase use_case.UnitUseCase
}

// API:
//
//	@Router		/units [post]
//	@Summary	Create
//	@tags		Units
//	@Accept		json
//	@Param		dto_request.UnitCreateRequest	body	dto_request.UnitCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{unit=dto_response.UnitResponse}}
func (a *UnitApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUnitCreate),
		func(ctx apiContext) {
			var request dto_request.UnitCreateRequest
			ctx.mustBind(&request)

			unit := a.unitUseCase.Create(ctx.context(), request)

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

// API:
//
//	@Router		/units/filter [post]
//	@Summary	Filter
//	@tags		Units
//	@Accept		json
//	@Param		dto_request.UnitFetchRequest	body	dto_request.UnitFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.UnitResponse}}
func (a *UnitApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUnitFetch),
		func(ctx apiContext) {
			var request dto_request.UnitFetchRequest
			ctx.mustBind(&request)

			units, total := a.unitUseCase.Fetch(ctx.context(), request)

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

// API:
//
//	@Router		/units/{id} [get]
//	@Summary	Get
//	@tags		Units
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{unit=dto_response.UnitResponse}}
func (a *UnitApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUnitGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.UnitGetRequest
			ctx.mustBind(&request)

			request.UnitId = id

			unit := a.unitUseCase.Get(ctx.context(), request)

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

// API:
//
//	@Router		/units/{id} [put]
//	@Summary	Update
//	@tags		Units
//	@Accept		json
//	@Param		id								path	string							true	"Id"
//	@Param		dto_request.UnitUpdateRequest	body	dto_request.UnitUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{unit=dto_response.UnitResponse}}
func (a *UnitApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUnitUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.UnitUpdateRequest
			ctx.mustBind(&request)

			request.UnitId = id

			unit := a.unitUseCase.Update(ctx.context(), request)

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

// API:
//
//	@Router		/units/{id} [delete]
//	@Summary	Delete
//	@tags		Units
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *UnitApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUnitDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.UnitDeleteRequest
			ctx.mustBind(&request)
			request.UnitId = id

			a.unitUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// API:
//
//	@Router		/units/options/product-unit-form [post]
//	@Summary	Option for product unit form
//	@tags		Units
//	@Accept		json
//	@Param		dto_request.UnitOptionForProductUnitFormRequest	body	dto_request.UnitOptionForProductUnitFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.UnitResponse}}
func (a *UnitApi) OptionForProductUnitForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionUnitOptionForProductUnitForm),
		func(ctx apiContext) {
			var request dto_request.UnitOptionForProductUnitFormRequest
			ctx.mustBind(&request)

			units, total := a.unitUseCase.OptionForProductUnitForm(ctx.context(), request)

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

func RegisterUnitApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := UnitApi{
		api:         newApi(useCaseManager),
		unitUseCase: useCaseManager.UnitUseCase(),
	}

	routerGroup := router.Group("/units")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.PUT("/:id", api.Update())
	routerGroup.DELETE("/:id", api.Delete())

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/product-unit-form", api.OptionForProductUnitForm())
}
