package service

import (
	"go-fiber-ent-web-layout/internal/cache"
	"go-fiber-ent-web-layout/internal/tools"
	"go-fiber-ent-web-layout/internal/usercase"
	"log/slog"
)

var users = []*usercase.User{
	{
		UserId:   1,
		Username: "example",
		Password: "admin",
		Scopes: []string{
			"example:list",
		},
	},
	{
		UserId:   2,
		Username: "admin",
		Password: "admin",
		Scopes: []string{
			"all",
		},
	},
}

type UserService struct {
	logger     *slog.Logger
	jwtService *tools.JwtService
	loginCache cache.LoginUserCache
}

func NewUserService(jwtService *tools.JwtService, loginCache cache.LoginUserCache) usercase.IUserService {
	return &UserService{
		logger:     slog.Default().With("trace-name", "user-service"),
		jwtService: jwtService,
		loginCache: loginCache,
	}
}

func (u *UserService) Login(user *usercase.User) (string, error) {
	for _, val := range users {
		if val.Username == user.Username && val.Password == user.Password {
			if token, err := u.jwtService.CreateToken(val); err != nil {
				return "", tools.FiberServerError("登录失败")
			} else {
				// 登录成功后添加到登录用户缓存
				u.loginCache.AddLoginUser(val)
				return token, nil
			}
		}
	}
	return "", tools.FiberRequestError("用户名或密码错误")
}
