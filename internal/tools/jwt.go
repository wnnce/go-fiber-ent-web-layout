package tools

import (
	"fmt"
	"github.com/bytedance/sonic"
	"github.com/golang-jwt/jwt/v5"
	"go-fiber-ent-web-layout/internal/conf"
	"log/slog"
	"strconv"
	"time"
)

// CustomClaims 自定义的Token负载结构体
type CustomClaims struct {
	Issuer    string           `json:"iss,omitempty"`
	Subject   string           `json:"sub,omitempty"`
	Audience  jwt.ClaimStrings `json:"aud,omitempty"`
	ExpiresAt *jwt.NumericDate `json:"exp,omitempty"`
	NotBefore *jwt.NumericDate `json:"nbf,omitempty"`
	IssuedAt  *jwt.NumericDate `json:"iat,omitempty"`
	ID        string           `json:"jti,omitempty"`
	Scope     jwt.ClaimStrings `json:"scope,omitempty"`
}

func (c CustomClaims) GetExpirationTime() (*jwt.NumericDate, error) {
	return c.ExpiresAt, nil
}
func (c CustomClaims) GetNotBefore() (*jwt.NumericDate, error) {
	return c.NotBefore, nil
}
func (c CustomClaims) GetIssuedAt() (*jwt.NumericDate, error) {
	return c.IssuedAt, nil
}
func (c CustomClaims) GetAudience() (jwt.ClaimStrings, error) {
	return c.Audience, nil
}
func (c CustomClaims) GetIssuer() (string, error) {
	return c.Issuer, nil
}
func (c CustomClaims) GetSubject() (string, error) {
	return c.Subject, nil
}
func (c CustomClaims) GetScope() (jwt.ClaimStrings, error) {
	return c.Scope, nil
}

var (
	issue      = "jwt"
	expireTime = 24 * time.Hour
	secret     = "secret"
)

func SetJwtConfig(config conf.Jwt) {
	if config.Issue != "" {
		issue = config.Issue
	}
	if config.ExpireTime > 0 {
		expireTime = config.ExpireTime
	}
	if config.Secret != "" {
		secret = config.Secret
	}
}

func GenerateToken(sub interface{}) (string, error) {
	subString, err := sonic.Marshal(sub)
	if err != nil {
		return "", err
	}
	currentTime := time.Now()
	numberDate := &jwt.NumericDate{Time: currentTime}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    issue,
		NotBefore: numberDate,
		IssuedAt:  numberDate,
		ExpiresAt: &jwt.NumericDate{Time: currentTime.Add(expireTime)},
		Subject:   string(subString),
		ID:        strconv.FormatInt(currentTime.UnixMilli(), 10),
	})
	return t.SignedString([]byte(secret))
}

func VerifyToken(tokenString string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			slog.Error("token unexpected signing method")
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		slog.Error("failed to resolve Claims")
		return nil, fmt.Errorf("failed to resolve Claims")
	}
	return claims, nil
}
