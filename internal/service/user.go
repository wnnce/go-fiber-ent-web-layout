package service

import (
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/factory"
	"go-fiber-ent-web-layout/internal/usercase"
	"log/slog"
)

var users = []*usercase.User{
	{
		Username: "example",
		Password: "admin",
		Scopes: []string{
			"example_select",
			"example_create",
		},
	},
	{
		Username: "admin",
		Password: "admin",
		Scopes: []string{
			"all",
		},
	},
}

type UserService struct {
	logger     *slog.Logger
	jwtService *common.JwtService
}

func NewUserService(jwtService *common.JwtService) usercase.IUserService {
	return &UserService{
		logger:     factory.GetLogger("user-service"),
		jwtService: jwtService,
	}
}

func (u *UserService) Login(user *usercase.User) (string, error) {
	for _, val := range users {
		if val.Username == user.Username && val.Password == user.Password {
			if token, err := u.jwtService.CreateToken(user, val.Scopes); err != nil {
				return "", common.FiberServerError("登录失败")
			} else {
				return token, nil
			}
		}
	}
	return "", common.FiberRequestError("用户名或密码错误")
}
