package user

import (
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/internal/tools/res"
	"go-fiber-ent-web-layout/internal/usercase"
)

type UserApi struct {
	service usercase.IUserService
}

func NewUserApi(service usercase.IUserService) *UserApi {
	return &UserApi{
		service: service,
	}
}

func (ua *UserApi) Login(ctx fiber.Ctx) error {
	user := &usercase.User{}
	if err := ctx.Bind().Body(user); err != nil {
		return err
	}
	token, err := ua.service.Login(user)
	if err != nil {
		return err
	}
	return ctx.JSON(res.OkByData(map[string]string{"token": token}))
}
