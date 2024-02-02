package model

import (
	"context"
	"myapp/constant"
	"myapp/infrastructure"
)

type dbtxCtxKeyType string

var dbtxCtxKey = dbtxCtxKeyType("dbtx")

func SetDbtxCtx(ctx context.Context, dbtx infrastructure.DBTX) (context.Context, error) {
	if _, err := GetDbtxCtx(ctx); err == nil {
		return nil, constant.ErrDbtxAlreadyExist
	}

	return context.WithValue(ctx, dbtxCtxKey, dbtx), nil
}

func GetDbtxCtx(ctx context.Context) (infrastructure.DBTX, error) {
	v, ok := ctx.Value(dbtxCtxKey).(infrastructure.DBTX)
	if !ok {
		return nil, constant.ErrDbtxNotFound
	}

	return v, nil
}
