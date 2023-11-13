package model

import (
	"context"
	"myapp/constant"
)

type requestActionCtxKeyType string

var requestActionCtxKey = requestActionCtxKeyType("request-action")

func SetRequestActionCtx(ctx context.Context, requestAction string) context.Context {
	return context.WithValue(ctx, requestActionCtxKey, requestAction)
}

func GetRequestActionCtx(ctx context.Context) (string, error) {
	v, ok := ctx.Value(requestActionCtxKey).(string)
	if !ok {
		return "", constant.ErrRequestActionNotFound
	}

	return v, nil
}

func MustGetRequestActionCtx(ctx context.Context) string {
	v, err := GetRequestActionCtx(ctx)
	if err != nil {
		panic(err)
	}

	return v
}
