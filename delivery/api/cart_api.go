package api

// import (
// 	"myapp/data_type"
// 	"myapp/delivery/dto_request"
// 	"myapp/delivery/dto_response"
// 	"myapp/use_case"
// 	"myapp/util"
// 	"net/http"

// 	"github.com/gin-gonic/gin"
// )

// type CartApi struct {
// 	api
// 	cartUseCase use_case.CartUseCase
// }

// // API:
// //
// //	@Router		/carts/options/user-form [post]
// //	@Summary	Create
// //	@tags		Carts
// //	@Accept		json
// //	@Param		dto_request.CartOptionForUserFormRequest	body	dto_request.CartOptionForUserFormRequest	true	"Body Request"
// //	@Produce	json
// //	@Success	200	{object}	dto_response.Response{data=dto_response.PaginationResponse{nodes=[]dto_response.CartResponse}}
// func (a *CartApi) Create() gin.HandlerFunc {
// 	return a.Authorize(
// 		data_type.PermissionP(data_type.PermissionCartOptionForUserForm),
// 		func(ctx apiContext) {
// 			var request dto_request.CartOptionForUserFormRequest
// 			ctx.mustBind(&request)

// 			carts, total := a.cartUseCase.OptionForUserForm(ctx.context(), request)

// 			ctx.json(
// 				http.StatusOK,
// 				dto_response.Response{
// 					Data: dto_response.PaginationResponse{
// 						Page:  request.Page,
// 						Limit: request.Limit,
// 						Total: total,
// 						Nodes: util.ConvertArray(carts, dto_response.NewCartResponse),
// 					},
// 				},
// 			)
// 		},
// 	)
// }

// func RegisterCartApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
// 	api := CartApi{
// 		api:         newApi(useCaseManager),
// 		cartUseCase: useCaseManager.CartUseCase(),
// 	}

// 	routerGroup := router.Group("/carts")
// 	routerGroup.POST("", api.Create())
// }
