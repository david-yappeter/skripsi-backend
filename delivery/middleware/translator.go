package middleware

import (
	"myapp/internal/gin/validator"
	internalI18n "myapp/internal/i18n"
	"myapp/internal/number_speller"
	"myapp/internal/time_speller"
	"myapp/model"

	"github.com/gin-gonic/gin"
	"golang.org/x/text/language"
)

func TranslatorHandler(router gin.IRouter) {
	router.Use(func(ctx *gin.Context) {
		var (
			matcher = language.NewMatcher([]language.Tag{
				language.Indonesian,
				language.English,
			})
			accept = ctx.GetHeader("Accept-Language")
		)

		tag, _ := language.MatchStrings(matcher, accept)
		locale := tag.String()

		if timeSpeller, ok := time_speller.Languages[locale]; ok {
			ctx.Request = ctx.Request.WithContext(model.SetTimeSpellerCtx(ctx.Request.Context(), timeSpeller))
		}

		if numberSpeller, ok := number_speller.Languages[locale]; ok {
			ctx.Request = ctx.Request.WithContext(model.SetNumberSpellerCtx(ctx.Request.Context(), numberSpeller))
		}

		if validatorTranslator, ok := validator.Translators[locale]; ok {
			ctx.Request = ctx.Request.WithContext(model.SetValidatorTranslatorCtx(ctx.Request.Context(), validatorTranslator))
		}

		// i18n
		ctx.Set("i18n", internalI18n.NewLocalizer(locale))

		ctx.Next()
	})
}
