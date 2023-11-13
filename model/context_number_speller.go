package model

import (
	"context"
	"myapp/internal/number_speller"
)

type numberSpellerCtxKeyType string

var numberSpellerCtxKey = numberSpellerCtxKeyType("number-speller")

func SetNumberSpellerCtx(ctx context.Context, numberSpeller number_speller.Interface) context.Context {
	return context.WithValue(ctx, numberSpellerCtxKey, numberSpeller)
}

func MustGetNumberSpellerCtx(ctx context.Context) number_speller.Interface {
	v, ok := ctx.Value(numberSpellerCtxKey).(number_speller.Interface)
	if !ok {
		return number_speller.DefaultLanguage
	}

	return v
}
