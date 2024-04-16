package limiter

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"path"
)

// Limiter 限流器接口
type Limiter interface {
	// DoLimit 限流方法 返回该请求是否被限流
	DoLimit(key string) bool

	// Name 返回限流器的名称
	Name() string
}

type LimitHandler struct {
	keyGenerate     KeyGenerate
	callbackHandler CallbackHandler
	limiters        []Limiter
	excludePaths    []string
	excludeIPs      []string
}

func (l *LimitHandler) addLimiter(limiters ...Limiter) {
	l.limiters = append(l.limiters, limiters...)
}

func NewMiddleware(config Config, ctx context.Context) fiber.Handler {
	config = configDefault(config)
	limitHandler := &LimitHandler{
		keyGenerate:     config.KeyGenerate,
		callbackHandler: config.CallbackHandler,
		excludePaths:    config.ExcludePaths,
		excludeIPs:      config.ExcludeIPs,
	}
	if config.Sliding.Enable {
		slidingLimiter := NewSlidingWindow(config.Sliding)
		slidingLimiter.TimingSideWindow(ctx)
		limitHandler.addLimiter(slidingLimiter)
	}
	if config.TokenBucket.Enable {
		tokenBucket := NewTokenBucket(config.TokenBucket)
		tokenBucket.TimingReleaseToken(ctx)
		limitHandler.addLimiter(tokenBucket)
	}
	limitHandler.addLimiter(config.CustomLimiters...)
	iPMatch := NewIPMatch()
	return func(fiberCtx fiber.Ctx) error {
		// 判断请求路径是否被排除
		url := fiberCtx.OriginalURL()
		for _, pattern := range limitHandler.excludePaths {
			if result, _ := path.Match(pattern, url); result {
				return fiberCtx.Next()
			}
		}
		// 判断请求IP是否被排除
		ip := fiberCtx.IP()
		for _, pattern := range limitHandler.excludeIPs {
			if iPMatch.Match(pattern, ip) {
				return fiberCtx.Next()
			}
		}
		// 生成用于限流的请求Key
		limitKey := limitHandler.keyGenerate(fiberCtx)
		// 遍历限流器 逐一执行
		for _, limiter := range limitHandler.limiters {
			if limiter.DoLimit(limitKey) {
				return limitHandler.callbackHandler(fiberCtx, limiter.Name())
			}
		}
		return fiberCtx.Next()
	}
}
