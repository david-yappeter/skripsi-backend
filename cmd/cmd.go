package cmd

import (
	"context"
	"fmt"
	"myapp/model"
	"myapp/util"
)

func newAuditLogCtx(action string) context.Context {
	ctx := context.Background()
	ctx = model.SetRequestIdCtx(ctx, util.NewUuid())
	ctx = model.SetRequestActionCtx(ctx, fmt.Sprintf("CMD %s", action))

	return ctx
}
