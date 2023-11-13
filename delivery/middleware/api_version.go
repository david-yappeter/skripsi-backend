package middleware

import (
	"myapp/constant"
	"myapp/global"

	"github.com/gin-gonic/gin"
)

func ApiVersionHandler(router gin.IRouter) {
	router.Use(func(ctx *gin.Context) {
		ctx.Header(constant.HeaderApiVersionKey, global.ApiVersion)
		ctx.Next()
	})
}
