// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/api/example/v1"
	"go-fiber-ent-web-layout/api/user/v1"
	"go-fiber-ent-web-layout/internal/cache"
	"go-fiber-ent-web-layout/internal/conf"
	"go-fiber-ent-web-layout/internal/data"
	"go-fiber-ent-web-layout/internal/middleware/auth"
	"go-fiber-ent-web-layout/internal/service"
)

// Injectors from wire.go:

// wireApp generate inject code
func wireApp(contextContext context.Context, confData *conf.Data, jwt *conf.Jwt, server *conf.Server) (*fiber.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData)
	if err != nil {
		return nil, nil, err
	}
	iExampleRepo := data.NewExampleRepo(dataData)
	iExampleService := service.NewExampleService(iExampleRepo)
	exampleApi := example.NewExampleApi(iExampleService)
	loginUserCache := cache.NewLoginUserCache()
	iUserService := service.NewUserService(loginUserCache)
	userApi := user.NewUserApi(iUserService)
	authMiddleware := auth.NewAuthMiddleware(loginUserCache)
	app := newApp(contextContext, server, exampleApi, userApi, authMiddleware)
	return app, func() {
		cleanup()
	}, nil
}
