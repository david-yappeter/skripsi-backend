package middleware

import (
	"fmt"
	"myapp/constant"
	"myapp/delivery/dto_response"
	"myapp/infrastructure"
	"runtime/debug"

	internalI18n "myapp/internal/i18n"

	"github.com/gin-gonic/gin"
)

func PanicHandler(router gin.IRouter, loggerStack infrastructure.LoggerStack) {
	router.Use(func(ctx *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				var (
					debugStackResponse *dto_response.ErrorResponse
					logMessage         = ""
				)

				switch v := r.(type) {
				case dto_response.ErrorResponse:
					instantResponse := v

					localizer, _ := ctx.MustGet("i18n").(*internalI18n.Localizer)
					localization, err := localizer.Translate(instantResponse.Message)
					if err == nil {
						instantResponse.Message = localization
					}

					ctx.AbortWithStatusJSON(instantResponse.Code, instantResponse)
					return

				case error:
					switch v {
					case constant.ErrNotAuthenticated:
						instantResponse := dto_response.NewUnauthorizedErrorResponse("Unauthorized")
						ctx.AbortWithStatusJSON(instantResponse.Code, instantResponse)
						return

					case constant.ErrForbidden:
						instantResponse := dto_response.NewForbiddenErrorResponseP("Forbidden")
						ctx.AbortWithStatusJSON(instantResponse.Code, instantResponse)
						return

					default:
						logMessage = fmt.Sprintf("Captured error: %s", v.Error())
					}

				default:
					logMessage = fmt.Sprintf("Unhandled err type %T, Content: %+v", v, v)
				}

				// write all error to loggers
				if len(logMessage) != 0 {
					logMessage += "\n"
				}
				logMessage += string(debug.Stack())
				loggerStack.WriteAll(logMessage)

				if debugStackResponse == nil {
					debugStackResponse = dto_response.NewInternalServerErrorResponseP()
				}
				ctx.AbortWithStatusJSON(debugStackResponse.Code, debugStackResponse)
			}
		}()

		ctx.Next()
	})
}
