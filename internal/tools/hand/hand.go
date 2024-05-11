package hand

import (
	"errors"
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/internal/tools/res"
	"net/http"
)

var defaultErrorHandlerChain *errorHandlerChain

func init() {
	defaultErrorHandlerChain = &errorHandlerChain{}
	RegisterErrorHandler(FiberErrorHandler)
}

// RegisterErrorHandler 注册错误处理函数
func RegisterErrorHandler(handler ErrorHandler) {
	defaultErrorHandlerChain.RegisterErrorHandler(handler)
}

// FiberErrorHandler fiber.Error 错误处理函数
func FiberErrorHandler(err error) *ErrorHandlerResult {
	var e *fiber.Error
	if errors.As(err, &e) {
		return &ErrorHandlerResult{
			Code:    e.Code,
			Message: e.Message,
		}
	}
	return nil
}

// CustomErrorHandler 自定义的错误处理
func CustomErrorHandler(ctx fiber.Ctx, err error) error {
	code, message := http.StatusInternalServerError, err.Error()
	if result := defaultErrorHandlerChain.DoHandler(err); result != nil {
		code = result.Code
		message = result.Message
	}
	return ctx.Status(code).JSON(res.Fail(code, message))
}
