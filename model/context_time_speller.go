package model

import (
	"context"
	"myapp/internal/time_speller"
)

type timeSpellerCtxTypeKey string

var timeSpellerCtxKey = timeSpellerCtxTypeKey("time-speller")

func SetTimeSpellerCtx(ctx context.Context, timeSpeller time_speller.Interface) context.Context {
	return context.WithValue(ctx, timeSpellerCtxKey, timeSpeller)
}

func MustGetTimeSpellerCtx(ctx context.Context) time_speller.Interface {
	v, ok := ctx.Value(timeSpellerCtxKey).(time_speller.Interface)
	if !ok {
		return time_speller.DefaultLanguage
	}

	return v
}
