package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go-fiber-ent-web-layout/api"
	"go-fiber-ent-web-layout/api/example/v1"
	"go-fiber-ent-web-layout/api/user/v1"
	"go-fiber-ent-web-layout/internal/conf"
	"go-fiber-ent-web-layout/internal/middleware/auth"
	"go-fiber-ent-web-layout/internal/middleware/limiter"
	"go-fiber-ent-web-layout/internal/middleware/timeout"
	"go-fiber-ent-web-layout/internal/tools"
	"log/slog"
	"os"
)

var confPath string

// 创建fiber app 包含注入中间件、错误处理、路由绑定等操作
func newApp(ctx context.Context, cf *conf.Server, eApi *example.ExampleApi, uApi *user.UserApi, auth *auth.AuthMiddleware) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      cf.Name,                  // 应用名称
		ErrorHandler: tools.CustomErrorHandler, // 自定义错误处理器
		JSONDecoder:  sonic.Unmarshal,          // 使用sonic进行Json序列化
		JSONEncoder:  sonic.Marshal,            // 使用sonic进行Json解析
	})
	// 防止程序panic 使用自定义的处理器 记录异常
	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: tools.CustomStackTraceHandler,
	}))
	app.Use(timeout.NewMiddleware(cf.Timeout))
	app.Use(limiter.NewMiddleware(limiter.Config{
		KeyGenerate:     limiter.Md5KeyGenerate(),
		CallbackHandler: limiter.DefaultCallbackHandler,
		Sliding:         cf.Limiter.Sliding,
		TokenBucket:     cf.Limiter.TokenBucket,
	}, ctx))
	api.RegisterRoutes(app, eApi, uApi, auth)
	return app
}

func init() {
	flag.StringVar(&confPath, "conf", "/configs/config-prod.yaml", "config path, eg: -conf config-prod.yaml")
}

func main() {
	flag.Parse()
	config := conf.ReadConfig(confPath)
	// 初始化日志
	handler := slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelInfo,
	})
	slog.SetDefault(slog.New(handler).With("app-name", config.Server.Name))
	ctx, cancel := context.WithCancel(context.Background())
	app, cleanup, err := wireApp(ctx, config.Data, config.Jwt, config.Server)
	if err != nil {
		panic(err)
	}
	defer func() {
		cancel()
		cleanup()
	}()
	if err = app.Listen(fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)); err != nil {
		panic(err)
	}
}
