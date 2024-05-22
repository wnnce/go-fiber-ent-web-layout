package auth

import (
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/internal/cache"
	"go-fiber-ent-web-layout/internal/tools"
	"go-fiber-ent-web-layout/internal/usercase"
	"time"
)

// AuthMiddleware 用户登录权限验证中间件
type AuthMiddleware struct {
	loginCache cache.LoginUserCache
}

func NewAuthMiddleware(loginCache cache.LoginUserCache) *AuthMiddleware {
	return &AuthMiddleware{
		loginCache: loginCache,
	}
}

// TokenAuth 登录验证，如果Token验证成功就将Sub参数和Scope权限参数存储到ctx.Locals中
// 后续中间件或者请求处理函数需要使用时，可以直接获取并使用类型转换
func (a *AuthMiddleware) TokenAuth(ctx fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authorization, ok := headers[fiber.HeaderAuthorization]
	if !ok || len(authorization[0]) <= 7 {
		return tools.FiberAuthError("The token does not exist")
	}
	claims, err := tools.VerifyToken(authorization[0][7:])
	// 判断Token时间是否符合要求
	currentTime := time.Now()
	if err != nil || claims.NotBefore.After(currentTime) {
		return tools.FiberAuthError("Invalid token")
	}
	if claims.ExpiresAt.Before(currentTime) {
		return tools.FiberAuthError("The token has expired")
	}
	// 是否能从Token中解析出用户配置
	user := &usercase.User{}
	if err = sonic.UnmarshalString(claims.Subject, user); err != nil {
		return tools.FiberAuthError("Invalid token")
	}
	// 判断用户的登录状态是否还有效
	loginUser := a.loginCache.GetLoginUser(user.GetUserId())
	if loginUser == nil {
		return tools.FiberAuthError("The user login is invalid")
	}

	// 设置请求用户缓存
	requestId := ctx.Context().ID()
	cache.SetRequestUser(requestId, loginUser)
	err = ctx.Next()
	// 请求完成后清除请求用户缓存
	cache.ClearRequestUser(requestId)
	return err
}

// VerifyPermissions 用户权限验证
func (a *AuthMiddleware) VerifyPermissions(permissions ...string) fiber.Handler {
	return func(ctx fiber.Ctx) error {
		if requestUser := cache.GetRequestUser(ctx.Context().ID()); requestUser != nil {
			for _, value := range requestUser.GetPermissions() {
				if value == "all" {
					return ctx.Next()
				}
				for _, p := range permissions {
					if p == value {
						return ctx.Next()
					}
				}
			}
		}
		return fiber.NewError(fiber.StatusForbidden, "No permission")
	}
}
