package tools

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v3"
	res "go-fiber-ent-web-layout/internal/tools/res"
	"log/slog"
	"net/http"
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
	result := res.Fail(code, message)
	return ctx.Status(code).JSON(result)
}

func FiberRequestError(message string) *fiber.Error {
	return fiber.NewError(http.StatusBadRequest, message)
}

func FiberAuthError(message string) *fiber.Error {
	return fiber.NewError(http.StatusUnauthorized, message)
}

func FiberServerError(message string) *fiber.Error {
	return fiber.NewError(http.StatusInternalServerError, message)
}