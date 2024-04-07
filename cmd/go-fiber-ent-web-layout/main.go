package main

import (
	"flag"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"go-fiber-ent-web-layout/api"
	"go-fiber-ent-web-layout/api/example/v1"
	"go-fiber-ent-web-layout/api/user/v1"
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/conf"
	"go-fiber-ent-web-layout/internal/factory"
	"go-fiber-ent-web-layout/internal/middleware"
	"log/slog"
	"os"
)

var (
	appName  = "go-fiber-ent-web-layout"
	confPath string
)

// 创建fiber app 包含注入中间件、错误处理、路由绑定等操作
func newApp(eApi *example.ExampleApi, uApi *user.UserApi, auth *middleware.AuthMiddleware) *fiber.App {
	app := fiber.New(fiber.Config{
		AppName:      appName,                   // 应用名称
		ErrorHandler: common.CustomErrorHandler, // 自定义错误处理器
		JSONDecoder:  sonic.Unmarshal,           // 使用sonic进行Json序列化
		JSONEncoder:  sonic.Marshal,             // 使用sonic进行Json解析
	})
	// 防止程序panic 使用自定义的处理器 记录异常
	app.Use(recover.New(recover.Config{
		EnableStackTrace:  true,
		StackTraceHandler: common.CustomStackTraceHandler,
	}))
	app.Use(middleware.HandlerTimeMiddleware)
	api.RegisterRoutes(app, eApi, uApi, auth)
	return app
}

func init() {
	flag.StringVar(&confPath, "conf", "/configs/config.yaml", "config path, eg: -conf config.yaml")
}

func main() {
	flag.Parse()
	config := conf.ReadConfig(confPath)
	factory.SetLoggerOptions(&factory.LoggerOptions{
		Level:  slog.LevelInfo,
		Output: os.Stdout,
		Args: map[string]string{
			"app-name": appName,
		},
	})
	app, cleanup, err := wireApp(config.Data, config.Jwt)
	if err != nil {
		panic(err)
	}
	defer cleanup()
	if err = app.Listen(config.Server.Addr); err != nil {
		panic(err)
	}
}
