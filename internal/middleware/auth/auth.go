package auth

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v3"
	"go-fiber-ent-web-layout/internal/tools"
	"go-fiber-ent-web-layout/internal/usercase"
	"time"
)

// TokenAuth 登录验证，如果Token验证成功就将Sub参数和Scope权限参数存储到ctx.Locals中
// 后续中间件或者请求处理函数需要使用时，可以直接获取并使用类型转换
func TokenAuth(ctx fiber.Ctx) error {
	authorization := ctx.Get(fiber.HeaderAuthorization)
	if len(authorization) <= 7 {
		return tools.FiberAuthError("The token does not exist")
	}
	claims, err := tools.VerifyToken(authorization[7:])
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
	loginUser := GetLoginUser(user.UserId)
	if loginUser == nil {
		return tools.FiberAuthError("The user login is invalid")
	}
	ctx.Locals("loginUser", loginUser)
	return ctx.Next()
}

// VerifyRoles 用户角色验证
func VerifyRoles(roles ...string) fiber.Handler {
	roleMap := make(map[string]struct{}, len(roles))
	for _, role := range roles {
		roleMap[role] = struct{}{}
	}
	return func(ctx fiber.Ctx) error {
		if loginUser := fiber.Locals[LoginUser](ctx, "loginUser"); loginUser != nil {
			for _, value := range loginUser.GetRoles() {
				if _, ok := roleMap[value]; ok || value == "admin" {
					return ctx.Next()
				}
			}
		}
		return fiber.NewError(fiber.StatusForbidden, "Current role has no permissions")
	}
}

// VerifyPermissions 用户权限验证
func VerifyPermissions(permissions ...string) fiber.Handler {
	permissionMap := make(map[string]struct{}, len(permissions))
	for _, permission := range permissions {
		permissionMap[permission] = struct{}{}
	}
	return func(ctx fiber.Ctx) error {
		fmt.Println("触发")
		if loginUser := fiber.Locals[LoginUser](ctx, "loginUser"); loginUser != nil {
			for _, value := range loginUser.GetPermissions() {
				if _, ok := permissionMap[value]; ok || value == "all" {
					return ctx.Next()
				}
			}
		}
		return fiber.NewError(fiber.StatusForbidden, "No permission")
	}
}
