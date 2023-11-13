package middleware

import (
	"fmt"
	"myapp/constant"
	"myapp/model"
	"myapp/util"

	"github.com/gin-gonic/gin"
)

func RequestIdHandler(router gin.IRouter) {
	router.Use(func(ctx *gin.Context) {
		requestId := util.NewUuid()

		ctx.Request = ctx.Request.WithContext(model.SetRequestIdCtx(ctx.Request.Context(), requestId))

		ctx.Request = ctx.Request.WithContext(
			model.SetRequestActionCtx(
				ctx.Request.Context(),
				fmt.Sprintf("%s %s", ctx.Request.Method, ctx.Request.URL.Path),
			),
		)

		ctx.Header(constant.HeaderRequestIdKey, requestId)

		ctx.Next()
	})
}
