package middleware

import (
	"context"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log/slog"
	"time"
)

func TimeoutMiddleware(timeout time.Duration) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		beforeTime := time.Now().UnixMilli()
		cont, cancel := context.WithTimeout(ctx.Context(), timeout)
		defer func() {
			cancel()
			afterTime := time.Now().UnixMilli()
			slog.Info(fmt.Sprintf("[method:%s uri:%s] 处理耗时：%dms", ctx.Method(), ctx.OriginalURL(), afterTime-beforeTime))
		}()
		ctx.SetUserContext(cont)
		return ctx.Next()
	}
}
