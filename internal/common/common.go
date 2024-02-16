package common

import (
	"errors"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go-fiber-ent-web-layout/internal/factory"
	"net/http"
)

var (
	logger    = factory.GetLogger("common")
	InjectSet = wire.NewSet(NewJwtService)
)

func CustomStackTraceHandler(ctx *fiber.Ctx, e interface{}) {
	trace := fmt.Sprintf("fiber application panic, StackTrace:%v, uri:%s, method:%s", e, ctx.OriginalURL(), ctx.Method())
	logger.Error(trace)
}

func CustomErrorHandler(ctx *fiber.Ctx, err error) error {
	code, message := http.StatusInternalServerError, "server error"
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
		message = e.Message
	}
	result := Fail(code, message)
	return ctx.Status(code).JSON(result)
}
