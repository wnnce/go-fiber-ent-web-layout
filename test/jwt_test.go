package test

import (
	"fmt"
	"github.com/bytedance/sonic"
	"go-fiber-ent-web-layout/internal/conf"
	"go-fiber-ent-web-layout/internal/tools"
	"go-fiber-ent-web-layout/internal/usercase"
	"testing"
	"time"
)

func TestTokenGenerate(t *testing.T) {
	jwtConf := &conf.Jwt{
		Issue:      "layout",
		Secret:     "hello world",
		ExpireTime: 24 * time.Hour,
	}
	jwtService := tools.NewJwtService(jwtConf)

	token, err := jwtService.CreateToken(&usercase.User{
		UserId:   1,
		Username: "admin",
		Scopes:   []string{"select", "create"},
	})
	if err != nil {
		fmt.Printf("%v\n", err)
	} else {
		println(token)
	}

	claims, err := jwtService.VerifyToken(token)
	if err != nil {
		println(err)
	} else {
		at, _ := claims.GetExpirationTime()
		println(at.UnixMilli())
		sub, _ := claims.GetSubject()
		tokenUser := &usercase.User{}
		err := sonic.Unmarshal([]byte(sub), tokenUser)
		if err == nil {
			fmt.Printf("%v\n", tokenUser)
		}
		for _, v := range tokenUser.GetPermissions() {
			fmt.Println(v)
		}
	}
}
