package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/wire"
	"go-fiber-ent-web-layout/api/example/v1"
	"go-fiber-ent-web-layout/api/user/v1"
	"go-fiber-ent-web-layout/internal/middlewares"
)

var InjectSet = wire.NewSet(example.NewExampleApi, user.NewUserApi)

// RegisterRoutes 全局路由绑定处理函数 在newApp函数中调用 不然wire无法处理依赖注入
func RegisterRoutes(app *fiber.App, eApi *example.ExampleApi, uApi *user.UserApi, auth *middlewares.AuthMiddleware) {
	exampleRoute := app.Group("/example", auth.TokenAuth)
	exampleRoute.Post("/", auth.ExampleCreatePermissions, eApi.SaveExample)
	exampleRoute.Put("/", auth.ExampleUpdatePermissions, eApi.UpdateExample)
	exampleRoute.Get("/list", auth.ExampleSelectPermissions, eApi.ListExample)
	exampleRoute.Get("/:id<int;min(1)>", eApi.QueryExample)

	userRoute := app.Group("/user")
	userRoute.Post("/login", uApi.Login)
}
