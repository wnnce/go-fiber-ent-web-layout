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
	exampleRoute.Post("/", eApi.SaveExample, auth.VerifyPermissions("example:save"))
	exampleRoute.Put("/", eApi.UpdateExample, auth.VerifyPermissions("example:update"))
	exampleRoute.Get("/list", eApi.ListExample, auth.VerifyPermissions("example:list"))
	exampleRoute.Get("/:id<int;min(1)>", eApi.QueryExample, auth.VerifyPermissions("example:query"))

	userRoute := app.Group("/user")
	userRoute.Post("/login", uApi.Login)
}
