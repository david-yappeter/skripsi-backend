package api

import (
	"fmt"
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"myapp/util"
	"net/http"

	"github.com/gin-gonic/gin"
)

type DeliveryOrderApi struct {
	api
	deliveryOrderUseCase use_case.DeliveryOrderUseCase
}

// API:
//
//	@Router		/delivery-orders [post]
//	@Summary	Create
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		dto_request.DeliveryOrderCreateRequest	body	dto_request.DeliveryOrderCreateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Create() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderCreate),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderCreateRequest
			ctx.mustBind(&request)

			deliveryOrder := a.deliveryOrderUseCase.Create(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/report [get]
//	@Summary	Download Excel Report
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		dto_request.DeliveryOrderDownloadReportRequest	body	dto_request.DeliveryOrderDownloadReportRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) DownloadReport() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDownloadReport),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderDownloadReportRequest
			ctx.mustBind(&request)

			ioReadCloser, contentLength, contentType, filename := a.deliveryOrderUseCase.DownloadReport(ctx.context(), request)

			ctx.dataFromReader(
				http.StatusOK,
				contentLength,
				contentType,
				ioReadCloser,
				map[string]string{
					"Content-Disposition": fmt.Sprintf("attachment; filename=\"%s\"", filename),
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/upload [post]
//	@Summary	Upload
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		file	formData	file	true	"Body with file"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{path=string}}
func (a *DeliveryOrderApi) Upload() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderUpload),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderUploadRequest
			ctx.mustBind(&request)

			path := a.deliveryOrderUseCase.Upload(ctx.context(), request)

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
//	@Router		/delivery-orders/{id}/items [post]
//	@Summary	Add Item
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.DeliveryOrderAddItemRequest	body	dto_request.DeliveryOrderAddItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) AddItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderAddItem),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderAddItemRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.AddItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/images [post]
//	@Summary	Add Image
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.DeliveryOrderAddImageRequest	body	dto_request.DeliveryOrderAddImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) AddImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderAddImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderAddImageRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.AddImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/drivers [post]
//	@Summary	Add Driver
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.DeliveryOrderAddDriverRequest	body	dto_request.DeliveryOrderAddDriverRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) AddDriver() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderAddDriver),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderAddDriverRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.AddDriver(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id} [put]
//	@Summary	Update
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.DeliveryOrderUpdateRequest	body	dto_request.DeliveryOrderUpdateRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Update() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderUpdate),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderUpdateRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.Update(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/cancel [patch]
//	@Summary	Cancel
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id										path	string									true	"Id"
//	@Param		dto_request.DeliveryOrderCancelRequest	body	dto_request.DeliveryOrderCancelRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Cancel() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderCancel),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderCancelRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.Cancel(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/completed [patch]
//	@Summary	Completed
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		dto_request.DeliveryOrderMarkCompletedRequest	body	dto_request.DeliveryOrderMarkCompletedRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Completed() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderMarkCompleted),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderMarkCompletedRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.MarkCompleted(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/delivery-location [patch]
//	@Summary	Delivery Location
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		dto_request.DeliveryOrderDeliveryLocationRequest	body	dto_request.DeliveryOrderDeliveryLocationRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *DeliveryOrderApi) DeliveryLocation() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDeliveryLocation),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderDeliveryLocationRequest
			ctx.mustBind(&request)

			a.deliveryOrderUseCase.DeliveryLocation(ctx.context(), request)

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
//	@Router		/delivery-orders/{id}/on-going [patch]
//	@Summary	On Going
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.DeliveryOrderMarkOngoingRequest	body	dto_request.DeliveryOrderMarkOngoingRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) OnGoing() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderMarkOngoing),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderMarkOngoingRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.MarkOngoing(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/delivering [patch]
//	@Summary	Delivering
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		dto_request.DeliveryOrderDeliveringRequest	body	dto_request.DeliveryOrderDeliveringRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Delivering() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDelivering),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderDeliveringRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.Delivering(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/filter [post]
//	@Summary	Filter
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		dto_request.DeliveryOrderFetchRequest	body	dto_request.DeliveryOrderFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderFetch),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderFetchRequest
			ctx.mustBind(&request)

			deliveryOrders, total := a.deliveryOrderUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(deliveryOrders, dto_response.NewDeliveryOrderResponse)

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
//	@Router		/delivery-orders/filter-driver [post]
//	@Summary	Filter for driver
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		dto_request.DeliveryOrderFetchDriverRequest	body	dto_request.DeliveryOrderFetchDriverRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.DataResponse{data=dto_response.PaginationResponse{nodes=[]dto_response.DeliveryOrderResponse},have_active_delivery=bool}
func (a *DeliveryOrderApi) FetchDriver() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderFetchDriver),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderFetchDriverRequest
			ctx.mustBind(&request)

			deliveryOrders, total := a.deliveryOrderUseCase.FetchDriver(ctx.context(), request)
			activeDeliveryOrder := a.deliveryOrderUseCase.ActiveForDriver(ctx.context())

			nodes := util.ConvertArray(deliveryOrders, dto_response.NewDeliveryOrderResponse)

			ctx.json(
				http.StatusOK,
				dto_response.DataResponse{
					"data": dto_response.PaginationResponse{
						Page:  request.Page,
						Limit: request.Limit,
						Total: total,
						Nodes: nodes,
					},
					"have_active_delivery": activeDeliveryOrder != nil,
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id} [get]
//	@Summary	Get
//	@tags		Delivery Orders
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderGetRequest
			ctx.mustBind(&request)

			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/guest [get]
//	@Summary	Get For Tracking Maps
//	@tags		Delivery Orders
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) GetGuest() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderGetRequest
			ctx.mustBind(&request)

			request.DeliveryOrderId = id

			deliveryOrder := a.deliveryOrderUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/active-driver [get]
//	@Summary	Get Active For Driver (Mobile)
//	@tags		Delivery Orders
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) ActiveForDriver() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderActiveForDriver),
		func(ctx apiContext) {
			deliveryOrder := a.deliveryOrderUseCase.ActiveForDriver(ctx.context())

			var resp *dto_response.DeliveryOrderResponse
			if deliveryOrder != nil {
				resp = dto_response.NewDeliveryOrderResponseP(*deliveryOrder)
			}

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": resp,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id} [delete]
//	@Summary	Delete
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *DeliveryOrderApi) Delete() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDelete),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderDeleteRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id

			a.deliveryOrderUseCase.Delete(ctx.context(), request)

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
//	@Router		/delivery-orders/{id}/items/{delivery_order_item_id} [delete]
//	@Summary	Delete Item
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		delivery_order_item_id						path	string										true	"Delivery Order Item Id"
//	@Param		dto_request.DeliveryOrderDeleteItemRequest	body	dto_request.DeliveryOrderDeleteItemRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) DeleteItem() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDeleteItem),
		func(ctx apiContext) {
			deliveryOrderId := ctx.getUuidParam("id")
			deliveryOrderItemId := ctx.getUuidParam("delivery_order_item_id")
			var request dto_request.DeliveryOrderDeleteItemRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = deliveryOrderId
			request.DeliveryOrderItemId = deliveryOrderItemId

			deliveryOrder := a.deliveryOrderUseCase.DeleteItem(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/images/{delivery_order_image_id} [delete]
//	@Summary	Delete File
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id											path	string										true	"Id"
//	@Param		delivery_order_image_id						path	string										true	"Id"
//	@Param		dto_request.DeliveryOrderDeleteImageRequest	body	dto_request.DeliveryOrderDeleteImageRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) DeleteImage() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDeleteImage),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			deliveryOrderImageId := ctx.getUuidParam("delivery_order_image_id")
			var request dto_request.DeliveryOrderDeleteImageRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id
			request.DeliveryOrderImageId = deliveryOrderImageId

			deliveryOrder := a.deliveryOrderUseCase.DeleteImage(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-orders/{id}/drivers/{driver_user_id} [delete]
//	@Summary	Delete Driver
//	@tags		Delivery Orders
//	@Accept		json
//	@Param		id												path	string											true	"Id"
//	@Param		driver_user_id									path	string											true	"Id"
//	@Param		dto_request.DeliveryOrderDeleteDriverRequest	body	dto_request.DeliveryOrderDeleteDriverRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order=dto_response.DeliveryOrderResponse}}
func (a *DeliveryOrderApi) DeleteDriver() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderDeleteDriver),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			driverUserId := ctx.getUuidParam("driver_user_id")
			var request dto_request.DeliveryOrderDeleteDriverRequest
			ctx.mustBind(&request)
			request.DeliveryOrderId = id
			request.DriverUserId = driverUserId

			deliveryOrder := a.deliveryOrderUseCase.DeleteDriver(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order": dto_response.NewDeliveryOrderResponse(deliveryOrder),
					},
				},
			)
		},
	)
}

func RegisterDeliveryOrderApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := DeliveryOrderApi{
		api:                  newApi(useCaseManager),
		deliveryOrderUseCase: useCaseManager.DeliveryOrderUseCase(),
	}

	routerGroup := router.Group("/delivery-orders")
	routerGroup.POST("", api.Create())
	routerGroup.GET("/report", api.DownloadReport())
	routerGroup.POST("/upload", api.Upload())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.POST("/filter-driver", api.FetchDriver())
	routerGroup.GET("/:id", api.Get())
	routerGroup.GET("/:id/guest", api.GetGuest())
	routerGroup.GET("/active-driver", api.ActiveForDriver())
	routerGroup.DELETE("/:id", api.Delete())

	routerGroup.POST("/:id/items", api.AddItem())
	routerGroup.POST("/:id/images", api.AddImage())
	routerGroup.POST("/:id/drivers", api.AddDriver())

	routerGroup.PUT("/:id", api.Update())
	routerGroup.PATCH("/:id/cancel", api.Cancel())
	routerGroup.PATCH("/:id/on-going", api.OnGoing())
	routerGroup.PATCH("/:id/delivering", api.Delivering())
	routerGroup.PATCH("/:id/completed", api.Completed())
	// routerGroup.PATCH("/:id/returned", api.Returned())
	routerGroup.PATCH("delivery-location", api.DeliveryLocation())

	routerGroup.DELETE("/:id/items/:delivery_order_item_id", api.DeleteItem())
	routerGroup.DELETE("/:id/images/:delivery_order_image_id", api.DeleteImage())
	routerGroup.DELETE("/:id/drivers/:driver_user_id", api.DeleteDriver())
}
