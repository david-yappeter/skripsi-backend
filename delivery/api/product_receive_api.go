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

// API:
//
//	@Router		/product-receives [post]
//	@Summary	Create
//	@tags		Product Receives
//	@Accept		json
//	@Param		dto_request.ProductReceiveCreateRequest	body	dto_request.ProductReceiveCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
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

// API:
//
//	@Router		/product-receives/upload [post]
//	@Summary	Upload
//	@tags		Product Receives
//	@Accept		json
//	@Param		file	formData	file	true	"Body with file"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
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

// API:
//
//	@Router		/product-receives/{id}/items [post]
//	@Summary	Add Item
//	@tags		Product Receives
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.ProductReceiveAddItemRequest	body	dto_request.ProductReceiveAddItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
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

// API:
//
//	@Router		/product-receives/{id}/images [post]
//	@Summary	Add Image
//	@tags		Product Receives
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.ProductReceiveAddImageRequest	body	dto_request.ProductReceiveAddImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
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

// API:
//
//	@Router		/product-receives/{id} [put]
//	@Summary	Update
//	@tags		Product Receives
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductReceiveUpdateRequest	body	dto_request.ProductReceiveUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveUpdateRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.Update(ctx.context(), request)

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

// API:
//
//	@Router		/product-receives/{id}/cancel [patch]
//	@Summary	Cancel
//	@tags		Product Receives
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductReceiveCancelRequest	body	dto_request.ProductReceiveCancelRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Cancel() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveCancel),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveCancelRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.Cancel(ctx.context(), request)

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

// API:
//
//	@Router		/product-receives/{id}/completed [patch]
//	@Summary	Update Completed
//	@tags		Product Receives
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		dto_request.ProductReceiveMarkCompleteRequest	body	dto_request.ProductReceiveMarkCompleteRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) MarkComplete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveMarkComplete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveMarkCompleteRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.MarkComplete(ctx.context(), request)

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

// API:
//
//	@Router		/product-receives/{id}/returned [patch]
//	@Summary	Returned
//	@tags		Product Receives
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		dto_request.ProductReceiveReturnedRequest	body	dto_request.ProductReceiveReturnedRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) Returned() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveReturned),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReceiveReturnedRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id

			productReceive := a.productReceiveUseCase.Returned(ctx.context(), request)

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

// API:
//
//	@Router		/product-receives/filter [post]
//	@Summary	Filter
//	@tags		Product Receives
//	@Accept		json
//	@Param		dto_request.ProductReceiveFetchRequest	body	dto_request.ProductReceiveFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductReceiveResponse}}
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

// API:
//
//	@Router		/product-receives/{id} [get]
//	@Summary	Get
//	@tags		Product Receives
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
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

// API:
//
//	@Router		/product-receives/{id} [delete]
//	@Summary	Delete
//	@tags		Product Receives
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
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

// API:
//
//	@Router		/product-receives/{id}/items/{product_receive_item_id} [delete]
//	@Summary	Delete Item
//	@tags		Product Receives
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		product_receive_item_id						path	string										true	"Product Receive Item Id"
//	@Param		dto_request.ProductReceiveDeleteItemRequest	body	dto_request.ProductReceiveDeleteItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) DeleteItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveDeleteItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			productReceiveItemId := ctx.getUuidParam("product_receive_item_id")
			var request dto_request.ProductReceiveDeleteItemRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id
			request.ProductReceiveItemId = productReceiveItemId

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

// API:
//
//	@Router		/product-receives/{id}/images/{product_receive_image_id} [post]
//	@Summary	Delete File
//	@tags		Product Receives
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		product_receive_image_id						path	string											true	"Product Receive Image Id"
//	@Param		dto_request.ProductReceiveDeleteImageRequest	body	dto_request.ProductReceiveDeleteImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_receive=dto_response.ProductReceiveResponse}}
func (a *ProductReceiveApi) DeleteImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReceiveDeleteImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			productReceiveImageId := ctx.getUuidParam("product_receive_image_id")
			var request dto_request.ProductReceiveDeleteImageRequest
			ctx.mustBind(&request)
			request.ProductReceiveId = id
			request.ProductReceiveImageId = productReceiveImageId

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

	routerGroup := router.Group("/product-receives")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/upload", api.Upload())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.DELETE("/:id", api.Delete())

	routerGroup.POST("/:id/items", api.AddItem())
	routerGroup.POST("/:id/images", api.AddImage())

	routerGroup.PUT("/:id", api.Update())
	routerGroup.PATCH("/:id/cancel", api.Cancel())
	routerGroup.PATCH("/:id/completed", api.MarkComplete())
	routerGroup.PATCH("/:id/returned", api.Returned())

	routerGroup.DELETE("/:id/items/:product_receive_item_id", api.DeleteItem())
	routerGroup.DELETE("/:id/images/:product_receive_image_id", api.DeleteImage())
}
