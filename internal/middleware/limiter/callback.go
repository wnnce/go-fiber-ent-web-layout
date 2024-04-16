package limiter

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/internal/tools/res"
	"log/slog"
	"net/http"
	"time"
)

type CallbackHandler func(fiber.Ctx, string) error

// DefaultCallbackHandler 请求被限流后的默认处理器
func DefaultCallbackHandler(ctx fiber.Ctx, limiterName string) error {
	slog.Warn(fmt.Sprintf("request is throttled requestId:%d, method:%s, orgiginUrl:%s, IP:%s, limiterName:%s",
		ctx.Context().ID(), ctx.Method(), ctx.OriginalURL(), ctx.IP(), limiterName))
	ctx.Status(http.StatusTooManyRequests).Set("Retry-After", time.Now().Add(1*time.Hour).String())
	return ctx.JSON(res.Fail(http.StatusTooManyRequests, "Requests are frequent"))
}
