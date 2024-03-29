package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/factory"
	"log/slog"
	"net/http"
	"time"
)

// AuthMiddleware 用户登录权限验证中间件
type AuthMiddleware struct {
	logger     *slog.Logger
	jwtService *common.JwtService
}

func NewAuthMiddleware(jwtService *common.JwtService) *AuthMiddleware {
	return &AuthMiddleware{
		logger:     factory.GetLogger("AuthMiddleware"),
		jwtService: jwtService,
	}
}

// TokenAuth 登录验证，如果Token验证成功就将Sub参数和Scope权限参数存储到ctx.Locals中
// 后续中间件或者请求处理函数需要使用时，可以直接获取并使用类型转换
func (a *AuthMiddleware) TokenAuth(ctx *fiber.Ctx) error {
	headers := ctx.GetReqHeaders()
	authorization, ok := headers[fiber.HeaderAuthorization]
	if !ok || len(authorization[0]) <= 7 {
		return common.FiberAuthError("The token does not exist")
	}
	claims, err := a.jwtService.VerifyToken(authorization[0][7:])
	currentTime := time.Now()
	if err != nil || claims.NotBefore.After(currentTime) {
		return common.FiberAuthError("Invalid token")
	}
	if claims.ExpiresAt.Before(currentTime) {
		return common.FiberAuthError("The token has expired")
	}
	ctx.Locals("tokenSub", claims.Subject)
	ctx.Locals("tokenScope", claims.Scope)
	return ctx.Next()
}

// VerifyPermissions 用户权限验证
func (a *AuthMiddleware) VerifyPermissions(permissions ...string) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		if scopes, ok := ctx.Locals("tokenScope").(jwt.ClaimStrings); ok {
			for _, value := range scopes {
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
		return fiber.NewError(http.StatusForbidden, "No permission")
	}
}
