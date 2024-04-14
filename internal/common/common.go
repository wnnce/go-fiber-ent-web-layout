package common

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"log/slog"
	"net/http"
)

var (
	InjectSet = wire.NewSet(NewJwtService)
)

func CustomStackTraceHandler(ctx fiber.Ctx, e interface{}) {
	trace := fmt.Sprintf("fiber application panic, StackTrace:%v, uri:%s, method:%s", e, ctx.OriginalURL(), ctx.Method())
	slog.Error(trace)
}

func CustomErrorHandler(ctx fiber.Ctx, err error) error {
	code, message := http.StatusInternalServerError, "server error"
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	result := Fail(code, message)
	return ctx.Status(code).JSON(result)
}
