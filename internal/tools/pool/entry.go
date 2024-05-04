package pool

import (
	"context"
	"math"
)

var defaultPool Pool

func init() {
	defaultPool = NewPool("defaultPool", math.MaxInt16)
}

func Go(fn func()) {
	CtxGo(context.Background(), fn)
}
func CtxGo(ctx context.Context, fn func()) {
	defaultPool.CtxGo(ctx, fn)
}
func DoGo(ctx context.Context, handler PanicHandler, fn func()) {
	defaultPool.DoGo(ctx, handler, fn)
}
func SetCap(cap int32) {
	defaultPool.SetCap(cap)
}
func WorkerCount() int32 {
	return defaultPool.WorkerCount()
}
