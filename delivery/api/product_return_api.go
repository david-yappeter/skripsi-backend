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

type ProductReturnApi struct {
	api
	productReturnUseCase use_case.ProductReturnUseCase
}

// API:
//
//	@Router		/product-returns [post]
//	@Summary	Create
//	@tags		Product Return
//	@Accept		json
//	@Param		dto_request.ProductReturnCreateRequest	body	dto_request.ProductReturnCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnCreate),
		func(ctx apiContext) {
			var request dto_request.ProductReturnCreateRequest
			ctx.mustBind(&request)

			productReturn := a.productReturnUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/upload [post]
//	@Summary	Upload
//	@tags		Product Return
//	@Accept		json
//	@Param		file	formData	file	true	"Body with file"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *ProductReturnApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnUpload),
		func(ctx apiContext) {
			var request dto_request.ProductReturnUploadRequest
			ctx.mustBind(&request)

			path := a.productReturnUseCase.Upload(ctx.context(), request)

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
//	@Router		/product-returns/{id}/items [post]
//	@Summary	Add Item
//	@tags		Product Return
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductReturnAddItemRequest	body	dto_request.ProductReturnAddItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) AddItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnAddItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReturnAddItemRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id

			productReturn := a.productReturnUseCase.AddItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/{id}/images [post]
//	@Summary	Add Image
//	@tags		Product Return
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.ProductReturnAddImageRequest	body	dto_request.ProductReturnAddImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) AddImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnAddImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReturnAddImageRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id

			productReturn := a.productReturnUseCase.AddImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/{id} [put]
//	@Summary	Update
//	@tags		Product Return
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.ProductReturnUpdateRequest	body	dto_request.ProductReturnUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReturnUpdateRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id

			productReturn := a.productReturnUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/{id}/completed [patch]
//	@Summary	Update Completed
//	@tags		Product Return
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		dto_request.ProductReturnMarkCompleteRequest	body	dto_request.ProductReturnMarkCompleteRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) MarkComplete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnMarkComplete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReturnMarkCompleteRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id

			productReturn := a.productReturnUseCase.MarkComplete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/filter [post]
//	@Summary	Filter
//	@tags		Product Return
//	@Accept		json
//	@Param		dto_request.ProductReturnFetchRequest	body	dto_request.ProductReturnFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnCreate),
		func(ctx apiContext) {
			var request dto_request.ProductReturnFetchRequest
			ctx.mustBind(&request)

			productReturns, total := a.productReturnUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(productReturns, dto_response.NewProductReturnResponse)

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
//	@Router		/product-returns/{id} [get]
//	@Summary	Get
//	@tags		Product Return
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReturnGetRequest
			ctx.mustBind(&request)

			request.ProductReturnId = id

			productReturn := a.productReturnUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/{id} [delete]
//	@Summary	Delete
//	@tags		Product Return
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *ProductReturnApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.ProductReturnDeleteRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id

			a.productReturnUseCase.Delete(ctx.context(), request)

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
//	@Router		/product-returns/{id}/items/{product_return_item_id} [delete]
//	@Summary	Delete Item
//	@tags		Product Return
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		product_return_item_id						path	string										true	"Product Return Item Id"
//	@Param		dto_request.ProductReturnDeleteItemRequest	body	dto_request.ProductReturnDeleteItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) DeleteItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnDeleteItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			productReturnItemId := ctx.getUuidParam("product_return_item_id")
			var request dto_request.ProductReturnDeleteItemRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id
			request.ProductReturnItemId = productReturnItemId

			productReturn := a.productReturnUseCase.DeleteItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/product-returns/{id}/images/{product_return_image_id} [post]
//	@Summary	Delete File
//	@tags		Product Return
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		product_return_image_id						path	string										true	"Product Receive Image Id"
//	@Param		dto_request.ProductReturnDeleteImageRequest	body	dto_request.ProductReturnDeleteImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{product_return=dto_response.ProductReturnResponse}}
func (a *ProductReturnApi) DeleteImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionProductReturnDeleteImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			productReturnImageId := ctx.getUuidParam("product_return_image_id")
			var request dto_request.ProductReturnDeleteImageRequest
			ctx.mustBind(&request)
			request.ProductReturnId = id
			request.ProductReturnImageId = productReturnImageId

			productReturn := a.productReturnUseCase.DeleteImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"product_return": dto_response.NewProductReturnResponse(productReturn),
					},
				},
			)
		},
	)
}

func RegisterProductReturnApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := ProductReturnApi{
		api:                  newApi(useCaseManager),
		productReturnUseCase: useCaseManager.ProductReturnUseCase(),
	}

	routerGroup := router.Group("/product-returns")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/upload", api.Upload())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.DELETE("/:id", api.Delete())

	routerGroup.POST("/:id/items", api.AddItem())
	routerGroup.POST("/:id/images", api.AddImage())

	routerGroup.PUT("/:id", api.Update())
	routerGroup.PATCH("/:id/completed", api.MarkComplete())

	routerGroup.DELETE("/:id/items/:product_return_item_id", api.DeleteItem())
	routerGroup.DELETE("/:id/images/:product_return_image_id", api.DeleteImage())
}
