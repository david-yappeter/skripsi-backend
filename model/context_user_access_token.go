package model

import (
	"context"
	"myapp/constant"
)

type userAccessTokenCtxKeyType string

var userAccessTokenCtxKey = userAccessTokenCtxKeyType("user-access-token")

func SetUserAccessTokenCtx(ctx context.Context, userAccessToken *UserAccessToken) context.Context {
	return context.WithValue(ctx, userAccessTokenCtxKey, userAccessToken)
}

func GetUserAccessTokenCtx(ctx context.Context) (*UserAccessToken, error) {
	v, ok := ctx.Value(userAccessTokenCtxKey).(*UserAccessToken)
	if !ok {
		return nil, constant.ErrNotAuthenticated
	}

	return v, nil
}

func MustGetUserAccessTokenCtx(ctx context.Context) UserAccessToken {
	v, err := GetUserAccessTokenCtx(ctx)
	if err != nil {
		panic(err)
	}

	if v == nil || v.Id == "" {
		panic(constant.ErrNotAuthenticated)
	}

	return *v
}
