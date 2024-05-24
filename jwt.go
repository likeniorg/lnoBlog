package main

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = "xniasxmioanmoxasndocja3fmoxahxax"

// 用于存储jwt中的信息
type Claims struct {
	ID       string `json:"id"`
	Identify string `json:"identify"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// 生成JWT
func GenerateJWT(id, identify, username string) (string, error) {
	claims := Claims{
		id,
		identify,
		username,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // 过期时间24小时
			IssuedAt:  jwt.NewNumericDate(time.Now()),                     // 签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                     // 生效时间
		},
	}
	// 使用HS256签名算法
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	s, err := t.SignedString([]byte(secretKey))

	return s, err
}

// 解析JWT
func ParseJwt(tokenstring string) (*Claims, error) {
	t, err := jwt.ParseWithClaims(tokenstring, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if claims, ok := t.Claims.(*Claims); ok && t.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
