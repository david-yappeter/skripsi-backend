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

// API:
//
//	@Router		/tiktok-products/upload-image [post]
//	@Summary	Upload Image
//	@tags		Tiktok Products
//	@Accept		json
//	@Param		dto_request.TiktokProductUploadImageRequest	body	dto_request.TiktokProductUploadImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{url=string,uri=string}}
func (a *TiktokProductApi) UploadImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionTiktokProductUploadImage),
		func(ctx apiContext) {
			var request dto_request.TiktokProductUploadImageRequest
			ctx.mustBind(&request)

			url, uri := a.tiktokProductUseCase.UploadImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"url": url,
						"uri": uri,
					},
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
	routerGroup.POST("/upload-image", api.UploadImage())

}
