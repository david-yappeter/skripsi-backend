package api

import (
	"fmt"
	"myapp/use_case"
	"time"

	"github.com/gin-gonic/gin"
)

type SsrApi struct {
	api
}

// API:
//
//	@Router		/ssr/maps [get]
//	@Summary	Get SSR Maps Data
//	@tags		Suppliers
//	@Accept		json
//	@Produce	json
func (a *SsrApi) Ssr() gin.HandlerFunc {
	return a.Guest(
		func(ctx apiContext) {
			// Set response headers for SSE
			ctx.ginCtx.Header("Content-Type", "text/event-stream")
			ctx.ginCtx.Header("Cache-Control", "no-cache")
			ctx.ginCtx.Header("Connection", "keep-alive")

			// Send a message every second
			for {
				data := fmt.Sprintf("data: %s\n\n", time.Now().Format("15:04:05"))
				ctx.ginCtx.Writer.WriteString(data)
				ctx.ginCtx.Writer.Flush()
				time.Sleep(1 * time.Second)
			}
		},
	)
}

func RegisterSsrApi(router gin.IRouter, useCaseManager use_case.UseCaseManager) {
	api := SsrApi{
		api: newApi(useCaseManager),
	}

	router.GET("/ssr/maps", api.Ssr())
}
