package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type TiktokProductApi struct {
	api
	tiktokProductUseCase use_case.TiktokProductUseCase
}

// API:
//
//	@Router		/tiktok-products [post]
//	@Summary	Create
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		dto_request.TiktokProductCreateRequest	body	dto_request.TiktokProductCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *TiktokProductApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductCreate),
		func(ctx apiContext) {
			var request dto_request.TiktokProductCreateRequest
			ctx.mustBind(&request)

			a.tiktokProductUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterTiktokProductApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := TiktokProductApi{
		api:                  newApi(useCaseManager),
		tiktokProductUseCase: useCaseManager.TiktokProductUseCase(),
	}

	routerGroup := router.Group("/tiktok-products")
	routerGroup.POST("", api.Create())

}
