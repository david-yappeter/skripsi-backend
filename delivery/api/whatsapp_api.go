package api

import (
	"myapp/data_type"
	"myapp/delivery/dto_request"
	"myapp/delivery/dto_response"
	"myapp/use_case"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WhatsappApi struct {
	api
	whatsappUseCase use_case.WhatsappUseCase
}

// API:
//
//	@Router		/whatsapp/is-logged-in [get]
//	@Summary	Check whatsapp is logged in
//	@tags		Whatsapps
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.DataResponse{is_logged_in=bool}}
func (a *WhatsappApi) IsLoggedIn() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappIsLoggedIn),
		func(ctx apiContext) {
			isLoggedIn := a.whatsappUseCase.IsLoggedIn(ctx.context())

			ctx.json(
				http.StatusOK,
				dto_response.Response{
					Data: dto_response.DataResponse{
						"is_logged_in": isLoggedIn,
					},
				},
			)
		},
	)

}

// API:
//
//	@Router		/whatsapp/logout [post]
//	@Summary	Logout
//	@tags		Whatsapps
//	@Accept		json
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *WhatsappApi) Logout() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappLogout),
		func(ctx apiContext) {
			a.whatsappUseCase.Logout(ctx.context())

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
//	@Router		/whatsapp/broadcast/product-price-change [post]
//	@Summary	Broadcast Product Price Change
//	@tags		Whatsapps
//	@Accept		json
//	@Param		dto_request.WhatsappProductPriceChangeBroadcastRequest	body	dto_request.WhatsappProductPriceChangeBroadcastRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *WhatsappApi) ProductPriceChangeBroadcast() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappProductPriceChangeBroadcast),
		func(ctx apiContext) {
			var request dto_request.WhatsappProductPriceChangeBroadcastRequest
			ctx.mustBind(&request)

			a.whatsappUseCase.ProductPriceChangeBroadcast(ctx.context(), request)

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
//	@Router		/whatsapp/broadcast/customer-debt [post]
//	@Summary	Broadcast Customer Debt
//	@tags		Whatsapps
//	@Accept		json
//	@Param		dto_request.WhatsappCustomerDebtBroadcastRequest	body	dto_request.WhatsappCustomerDebtBroadcastRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *WhatsappApi) CustomerDebtBroadcast() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappCustomerDebtBroadcast),
		func(ctx apiContext) {
			var request dto_request.WhatsappCustomerDebtBroadcastRequest
			ctx.mustBind(&request)

			a.whatsappUseCase.CustomerDebtBroadcast(ctx.context(), request)

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
//	@Router		/whatsapp/broadcast/customer-type-discount [post]
//	@Summary	Broadcast Customer Type Discount
//	@tags		Whatsapps
//	@Accept		json
//	@Param		dto_request.WhatsappCustomerTypeDiscountBroadcastRequest	body	dto_request.WhatsappCustomerTypeDiscountBroadcastRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.SuccessResponse
func (a *WhatsappApi) CustomerTypeDiscountBroadcast() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionWhatsappCustomerTypeDiscountBroadcast),
		func(ctx apiContext) {
			var request dto_request.WhatsappCustomerTypeDiscountBroadcastRequest
			ctx.mustBind(&request)

			a.whatsappUseCase.CustomerTypeDiscountBroadcast(ctx.context(), request)

			ctx.json(
				http.StatusOK,
				dto_response.SuccessResponse{
					Message: "OK",
				},
			)
		},
	)
}

func RegisterWhatsappApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := WhatsappApi{
		api:             newApi(useCaseManager),
		whatsappUseCase: useCaseManager.WhatsappUseCase(),
	}

	routerGroup := router.Group("/whatsapp")
	routerGroup.GET("/is-logged-in", api.IsLoggedIn())
	routerGroup.POST("/logout", api.Logout())

	broadcastRouterGroup := routerGroup.Group("/broadcast")
	broadcastRouterGroup.POST("/product-price-change", api.ProductPriceChangeBroadcast())
	broadcastRouterGroup.POST("/customer-debt", api.CustomerDebtBroadcast())
	broadcastRouterGroup.POST("/customer-type-discount", api.CustomerTypeDiscountBroadcast())
}
