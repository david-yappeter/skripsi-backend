package util

import (
	"context"
	"time"
)

type noCancel struct {
	ctx context.Context
}

func (c noCancel) Deadline() (time.Time, bool)       { return time.Time{}, false }
func (c noCancel) Done() <-chan struct{}             { return nil }
func (c noCancel) Err() error                        { return nil }
func (c noCancel) Value(key interface{}) interface{} { return c.ctx.Value(key) }

// WithoutCancelCtx returns a context that is never canceled.
func WithoutCancelCtx(ctx context.Context) context.Context {
	return noCancel{ctx: ctx}
}
