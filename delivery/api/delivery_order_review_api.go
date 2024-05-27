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

type DeliveryOrderReviewApi struct {
	api
	deliveryOrderReviewUseCase use_case.DeliveryOrderReviewUseCase
}

// API:
//
//	@Router		/delivery-order-reviews/guest [post]
//	@Summary	Create for Guest (Maps)
//	@tags		Delivery Order Reviews
//	@Accept		json
//	@Param		dto_request.DeliveryOrderReviewCreateGuestRequest	body	dto_request.DeliveryOrderReviewCreateGuestRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order_review=dto_response.DeliveryOrderReviewResponse}}
func (a *DeliveryOrderReviewApi) CreateGuest() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderReviewCreateGuestRequest
			ctx.mustBind(&request)

			deliveryOrderReview := a.deliveryOrderReviewUseCase.CreateGuest(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order_review": dto_response.NewDeliveryOrderReviewResponse(deliveryOrderReview),
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-order-reviews/delivery-orders/{delivery_order_id} [post]
//	@Summary	Get Is Exist By Delivery Order Id
//	@tags		Delivery Order Reviews
//	@Accept		json
//	@Param		delivery_order_id	path	string	true	"Delivery Order Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{review_exist=bool}}
func (a *DeliveryOrderReviewApi) IsExistByDeliveryOrder() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			id := ctx.getUuidParam("delivery_order_id")
			var request dto_request.DeliveryOrderReviewIsExistByDeliveryOrderRequest
			ctx.mustBind(&request)

			request.DeliveryOrderId = id

			exist := a.deliveryOrderReviewUseCase.IsExistByDeliveryOrder(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"review_exist": exist,
					},
				},
			)
		},
	)
}

// API:
//
//	@Router		/delivery-order-reviews/filter [post]
//	@Summary	Filter
//	@tags		Delivery Order Reviews
//	@Accept		json
//	@Param		dto_request.DeliveryOrderReviewFetchRequest	body	dto_request.DeliveryOrderReviewFetchRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.DeliveryOrderReviewResponse}}
func (a *DeliveryOrderReviewApi) Fetch() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderReviewFetch),
		func(ctx apiContext) {
			var request dto_request.DeliveryOrderReviewFetchRequest
			ctx.mustBind(&request)

			deliveryOrderReviews, total := a.deliveryOrderReviewUseCase.Fetch(ctx.context(), request)

			nodes := util.ConvertArray(deliveryOrderReviews, dto_response.NewDeliveryOrderReviewResponse)

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
//	@Router		/delivery-order-reviews/{id} [get]
//	@Summary	Get
//	@tags		Delivery Order Reviews
//	@Param		id	path	string	true	"Id"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{delivery_order_review=dto_response.DeliveryOrderReviewResponse}}
func (a *DeliveryOrderReviewApi) Get() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionDeliveryOrderReviewGet),
		func(ctx apiContext) {
			id := ctx.getUuidParam("id")
			var request dto_request.DeliveryOrderReviewGetRequest
			ctx.mustBind(&request)

			request.DeliveryOrderReviewId = id

			deliveryOrderReview := a.deliveryOrderReviewUseCase.Get(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"delivery_order_review": dto_response.NewDeliveryOrderReviewResponse(deliveryOrderReview),
					},
				},
			)
		},
	)
}

func RegisterDeliveryOrderReviewApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := DeliveryOrderReviewApi{
		api:                        newApi(useCaseManager),
		deliveryOrderReviewUseCase: useCaseManager.DeliveryOrderReviewUseCase(),
	}

	routerGroup := router.Group("/delivery-order-reviews")
	routerGroup.POST("/guest", api.CreateGuest())
	routerGroup.POST("/delivery-orders/:delivery_order_id", api.IsExistByDeliveryOrder())
	routerGroup.POST("/filter", api.Fetch())
	routerGroup.GET("/:id", api.Get())
}
