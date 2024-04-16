package user

import (
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/common/res"
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
	if err := ctx.Bind().JSON(user); err != nil {
		return err
	}
	if errorMessage := common.StructFieldValidation(user); len(errorMessage) > 0 {
		return res.FiberRequestError(errorMessage)
	}
	token, err := ua.service.Login(user)
	if err != nil {
		return err
	}
	return ctx.JSON(res.OkByData(map[string]string{"token": token}))
}
