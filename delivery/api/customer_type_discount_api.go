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

type CustomerTypeDiscountApi struct {
	api
	customerTypeDiscountUseCase use_case.CustomerTypeDiscountUseCase
}

// API:
//
//	@Router		/customer-type-discounts/options/whatsapp-customer-type-discount-change-broadcast-form [post]
//	@Summary	Option for Whatsapp Customer Type Discount Change Broadcast Form
//	@tags		Customer Types
//	@Accept		json
//	@Param		dto_request.CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest	body	dto_request.CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest	true	"Body Request"
//	@Produce	json
//	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CustomerTypeDiscountResponse}}
func (a *CustomerTypeDiscountApi) OptionForWhatsappCustomerTypeDiscountChangeBroadcastForm() gin.HandlerFunc {
	return a.Authorize(
		data_type.PermissionP(data_type.PermissionCustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastForm),
		func(ctx apiContext) {
			var request dto_request.CustomerTypeDiscountOptionForWhatsappCustomerTypeDiscountChangeBroadcastFormRequest
			ctx.mustBind(&request)

			customerTypeDiscounts, total := a.customerTypeDiscountUseCase.OptionForWhatsappCustomerTypeDiscountChangeBroadcastForm(ctx.context(), request)

			nodes := util.ConvertArray(customerTypeDiscounts, dto_response.NewCustomerTypeDiscountResponse)

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

func RegisterCustomerTypeDiscountApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := CustomerTypeDiscountApi{
		api:                         newApi(useCaseManager),
		customerTypeDiscountUseCase: useCaseManager.CustomerTypeDiscountUseCase(),
	}

	routerGroup := router.Group("/customer-type-discounts")

	optionRouterGroup := routerGroup.Group("/options")
	optionRouterGroup.POST("/whatsapp-customer-type-discount-change-broadcast-form", api.OptionForWhatsappCustomerTypeDiscountChangeBroadcastForm())
}
