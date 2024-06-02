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

type PurchaseOrderApi struct {
	api
	purchaseOrderUseCase use_case.PurchaseOrderUseCase
}

// API:
//
//	@Router		/purchase-orders [post]
//	@Summary	Create
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		dto_request.PurchaseOrderCreateRequest	body	dto_request.PurchaseOrderCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderCreate),
		func(ctx apiContext) {
			var request dto_request.PurchaseOrderCreateRequest
			ctx.mustBind(&request)

			purchaseOrder := a.purchaseOrderUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/upload [post]
//	@Summary	Upload
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		file	formData	file	true	"Body with file"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *PurchaseOrderApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderUpload),
		func(ctx apiContext) {
			var request dto_request.PurchaseOrderUploadRequest
			ctx.mustBind(&request)

			path := a.purchaseOrderUseCase.Upload(ctx.context(), request)

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
//	@Router		/purchase-orders/{id}/items [post]
//	@Summary	Add Item
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.PurchaseOrderAddItemRequest	body	dto_request.PurchaseOrderAddItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) AddItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderAddItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderAddItemRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.AddItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id}/images [post]
//	@Summary	Add Image
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.PurchaseOrderAddImageRequest	body	dto_request.PurchaseOrderAddImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) AddImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderAddImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderAddImageRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.AddImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id} [put]
//	@Summary	Update
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.PurchaseOrderUpdateRequest	body	dto_request.PurchaseOrderUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderUpdateRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id}/ongoing [patch]
//	@Summary	Ongoing
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.PurchaseOrderOngoingRequest	body	dto_request.PurchaseOrderOngoingRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) Ongoing() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderOngoing),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderOngoingRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.Ongoing(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id}/cancel [patch]
//	@Summary	Cancel
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.PurchaseOrderCancelRequest	body	dto_request.PurchaseOrderCancelRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) Cancel() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderCancel),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderCancelRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.Cancel(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id}/completed [patch]
//	@Summary	Update Completed
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		dto_request.PurchaseOrderMarkCompleteRequest	body	dto_request.PurchaseOrderMarkCompleteRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) MarkComplete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderMarkComplete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderMarkCompleteRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.MarkComplete(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/filter [post]
//	@Summary	Filter
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		dto_request.PurchaseOrderFetchRequest	body	dto_request.PurchaseOrderFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderCreate),
		func(ctx apiContext) {
			var request dto_request.PurchaseOrderFetchRequest
			ctx.mustBind(&request)

			purchaseOrders, total := a.purchaseOrderUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(purchaseOrders, dto_response.NewPurchaseOrderResponse)

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
//	@Router		/purchase-orders/{id} [get]
//	@Summary	Get
//	@tags		Purchase Orders
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderGetRequest
			ctx.mustBind(&request)

			request.PurchaseOrderId = id

			purchaseOrder := a.purchaseOrderUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id} [delete]
//	@Summary	Delete
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *PurchaseOrderApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.PurchaseOrderDeleteRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id

			a.purchaseOrderUseCase.Delete(ctx.context(), request)

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
//	@Router		/purchase-orders/{id}/items/{purchase_order_item_id} [delete]
//	@Summary	Delete Item
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		purchase_order_item_id						path	string										true	"Product Receive Item Id"
//	@Param		dto_request.PurchaseOrderDeleteItemRequest	body	dto_request.PurchaseOrderDeleteItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) DeleteItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderDeleteItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			purchaseOrderItemId := ctx.getUuidParam("purchase_order_item_id")
			var request dto_request.PurchaseOrderDeleteItemRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id
			request.PurchaseOrderItemId = purchaseOrderItemId

			purchaseOrder := a.purchaseOrderUseCase.DeleteItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/purchase-orders/{id}/images/{purchase_order_image_id} [post]
//	@Summary	Delete File
//	@tags		Purchase Orders
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		purchase_order_image_id						path	string											true	"Product Receive Image Id"
//	@Param		dto_request.PurchaseOrderDeleteImageRequest	body	dto_request.PurchaseOrderDeleteImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{purchase_order=dto_response.PurchaseOrderResponse}}
func (a *PurchaseOrderApi) DeleteImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionPurchaseOrderDeleteImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			purchaseOrderImageId := ctx.getUuidParam("purchase_order_image_id")
			var request dto_request.PurchaseOrderDeleteImageRequest
			ctx.mustBind(&request)
			request.PurchaseOrderId = id
			request.PurchaseOrderImageId = purchaseOrderImageId

			purchaseOrder := a.purchaseOrderUseCase.DeleteImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"purchase_order": dto_response.NewPurchaseOrderResponse(purchaseOrder),
					},
				},
			)
		},
	)
}

func RegisterPurchaseOrderApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := PurchaseOrderApi{
		api:                  newApi(useCaseManager),
		purchaseOrderUseCase: useCaseManager.PurchaseOrderUseCase(),
	}

	routerGroup := router.Group("/purchase-orders")
	routerGroup.POST("", api.Create())
	routerGroup.POST("/upload", api.Upload())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
	routerGroup.DELETE("/:id", api.Delete())

	routerGroup.POST("/:id/items", api.AddItem())
	routerGroup.POST("/:id/images", api.AddImage())

	routerGroup.PUT("/:id", api.Update())
	routerGroup.PATCH("/:id/ongoing", api.Ongoing())
	routerGroup.PATCH("/:id/cancel", api.Cancel())
	routerGroup.PATCH("/:id/completed", api.MarkComplete())

	routerGroup.DELETE("/:id/items/:purchase_order_item_id", api.DeleteItem())
	routerGroup.DELETE("/:id/images/:purchase_order_image_id", api.DeleteImage())
}
