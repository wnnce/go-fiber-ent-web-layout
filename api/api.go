package api

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"go-fiber-ent-web-layout/api/example/v1"
	"go-fiber-ent-web-layout/api/manage"
	"go-fiber-ent-web-layout/api/user/v1"
	"go-fiber-ent-web-layout/internal/middleware/auth"
)

var InjectSet = wire.NewSet(example.NewExampleApi, user.NewUserApi)

// RegisterRoutes 全局路由绑定处理函数 在newApp函数中调用 不然wire无法处理依赖注入
func RegisterRoutes(app *fiber.App, eApi *example.ExampleApi, uApi *user.UserApi) {
	manageRoute := app.Group("/manage")
	manageRoute.Get("/logger/sse/:interval<int;min<100>>", manage.LoggerPush)

	exampleRoute := app.Group("/example", auth.TokenAuth)
	exampleRoute.Post("/", auth.VerifyPermissions("example:save"), eApi.SaveExample)
	exampleRoute.Put("/", auth.VerifyPermissions("example:update"), eApi.UpdateExample)
	exampleRoute.Get("/list", auth.VerifyPermissions("example:list"), eApi.ListExample)
	exampleRoute.Get("/:id<int;min(1)>", eApi.QueryExample)

	userRoute := app.Group("/user")
	userRoute.Post("/login", uApi.Login)
}
