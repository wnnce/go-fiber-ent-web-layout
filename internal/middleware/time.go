package middleware

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-fiber-ent-web-layout/internal/factory"
	"time"
)

var logger = factory.GetLogger("time-middleware")

func HandlerTimeMiddleware(ctx *fiber.Ctx) error {
	beforeTime := time.Now().UnixMilli()
	err := ctx.Next()
	afterTime := time.Now().UnixMilli()
	logger.Info(fmt.Sprintf("[Method：%s %s] 处理耗时：%dms", ctx.Method(), ctx.OriginalURL(), afterTime-beforeTime))
	return err
}
