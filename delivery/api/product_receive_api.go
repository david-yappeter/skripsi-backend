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

type ProductReceiveApi struct {
	api
	productReceiveUseCase use_case.ProductReceiveUseCase
}

// @Router		/product-receives [post]
// @Summary	Create
// @tags		Product Receives
// @Accept		json
// @Param		dto_request.ProductReceiveCreateRequest	body	dto_request.ProductReceiveCreateRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveCreate),
		func(ctx apiContext) {
			var request dto_request.ProductReceiveCreateRequest
			ctx.mustBind(&request)

			productReceive := a.productReceiveUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

// @Router		/product-receives/upload [post]
// @Summary	Upload
// @tags		Product Receives
// @Accept		json
// @Param		dto_request.ProductReceiveUploadRequest	body	dto_request.ProductReceiveUploadRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *ProductReceiveApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveUpload),
		func(ctx apiContext) {
			var request dto_request.ProductReceiveUploadRequest
			ctx.mustBind(&request)

			path := a.productReceiveUseCase.Upload(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"path": path,
					},
				},
			)
		},
	)
}

// @Router		/product-receives/{id}/items [post]
// @Summary	Add Item
// @tags		Product Receives
// @Accept		json
// @Param		id											path	string										true	"Id"
// @Param		dto_request.ProductReceiveAddItemRequest	body	dto_request.ProductReceiveAddItemRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) AddItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveAddItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveAddItemRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.AddItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

// @Router		/product-receives/{id}/images [post]
// @Summary	Add Image
// @tags		Product Receives
// @Accept		json
// @Param		id											path	string										true	"Id"
// @Param		dto_request.ProductReceiveAddImageRequest	body	dto_request.ProductReceiveAddImageRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) AddImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveAddImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveAddImageRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.AddImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

// @Router		/product-receives/filter [post]
// @Summary	Filter
// @tags		Product Receives
// @Accept		json
// @Param		dto_request.ProductReceiveFetchRequest	body	dto_request.ProductReceiveFetchRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveCreate),
		func(ctx apiContext) {
			var request dto_request.ProductReceiveFetchRequest
			ctx.mustBind(&request)

			productReceives, total := a.productReceiveUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(productReceives, dto_response.NewProductReceiveResponse)

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

// @Router		/product-receives/{id} [get]
// @Summary	Update
// @tags		Product Receives
// @Param		id	path	string	true	"Id"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveGetRequest
			ctx.mustBind(&request)

			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

// @Router		/product-receives/{id} [delete]
// @Summary	Update Password
// @tags		Product Receives
// @Accept		json
// @Param		id	path	string	true	"Id"
// @Produce	json
// @Success	200	{object}	dto_response.SuccessResponse
func (a *ProductReceiveApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveDeleteRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			a.productReceiveUseCase.Delete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

// @Router		/product-receives/{id}/items/{product_unit_id} [delete]
// @Summary	Delete Item
// @tags		Product Receives
// @Accept		json
// @Param		id											path	string										true	"Id"
// @Param		product_unit_id											path	string										true	"Product Unit Id"
// @Param		dto_request.ProductReceiveDeleteItemRequest	body	dto_request.ProductReceiveDeleteItemRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) DeleteItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveDeleteItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			productUnitId := ctx.getUuidParam("product_unit_id")
			var request dto_request.ProductReceiveDeleteItemRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id
			request.ProductUnitId = productUnitId

			productReceive := a.productReceiveUseCase.DeleteItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

// @Router		/product-receives/{id}/images/{file_id} [post]
// @Summary	Delete File
// @tags		Product Receives
// @Accept		json
// @Param		id											path	string										true	"Id"
// @Param		dto_request.ProductReceiveDeleteImageRequest	body	dto_request.ProductReceiveDeleteImageRequest	true	"Body Request"
// @Produce	json
// @Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) DeleteImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveDeleteImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			fileId := ctx.getUuidParam("file_id")
			var request dto_request.ProductReceiveDeleteImageRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id
			request.FileId = fileId

			productReceive := a.productReceiveUseCase.DeleteImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_receive": dto_response.NewProductReceiveResponse(productReceive),
					},
				},
			)
		},
	)
}

func RegisterProductReceiveApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductReceiveApi{
		api:                   newApi(useCaseManager),
		productReceiveUseCase: useCaseManager.ProductReceiveUseCase(),
	}

	adminRouterGroup := router.Group("/product-receives")
	adminRouterGroup.POST("", api.Create())
	adminRouterGroup.POST("/upload", api.Upload())
	adminRouterGroup.POST("/filter", api.Fetch())
	adminRouterGroup.GET("/:id", api.Get())
	adminRouterGroup.DELETE("/:id", api.Delete())

	adminRouterGroup.POST("/:id/items", api.AddItem())
	adminRouterGroup.POST("/:id/images", api.AddImage())

	adminRouterGroup.DELETE("/:id/items/:product_unit_id", api.DeleteItem())
	adminRouterGroup.DELETE("/:id/images/:file_id", api.DeleteImage())
}
