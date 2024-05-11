package hand

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"log/slog"
)

func StackTraceHandler(ctx fiber.Ctx, e interface{}) {
	trace := fmt.Sprintf("fiber application panic, StackTrace:%v, uri:%s, method:%s", e, ctx.OriginalURL(), ctx.Method())
	slog.Error(trace)
}
