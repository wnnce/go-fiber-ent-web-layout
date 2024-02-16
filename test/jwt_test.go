package test

import (
	"fmt"
	"github.com/bytedance/sonic"
	"go-fiber-ent-web-layout/internal/common"
	"go-fiber-ent-web-layout/internal/conf"
	"testing"
	"time"
)

type user struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func TestTokenGenerate(t *testing.T) {
	jwtConf := &conf.Jwt{
		Issue:      "layout",
		Secret:     "hello world",
		ExpireTime: 24 * time.Hour,
	}
	jwtService := common.NewJwtService(jwtConf)

	token, err := jwtService.CreateToken(&user{
		ID:   1,
		Name: "xin",
	}, []string{"all", "user"})
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
		tokenUser := &user{}
		err := sonic.Unmarshal([]byte(sub), tokenUser)
		if err == nil {
			fmt.Printf("%v\n", tokenUser)
		}
		scope, _ := claims.GetScope()
		println(scope)
	}
}
