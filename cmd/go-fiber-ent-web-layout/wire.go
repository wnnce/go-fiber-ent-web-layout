//go:build wireinject
// +build wireinject

package main

import (
	"context"
	"github.com/gofiber/fiber/v3"
	"github.com/google/wire"
	"go-fiber-ent-web-layout/api"
	"go-fiber-ent-web-layout/internal/conf"
	"go-fiber-ent-web-layout/internal/data"
	"go-fiber-ent-web-layout/internal/service"
)

// wireApp generate inject code
func wireApp(context.Context, *conf.Data, *conf.Jwt, *conf.Server) (*fiber.App, func(), error) {
	panic(wire.Build(api.InjectSet, data.InjectSet, service.InjectSet, newApp))
}
