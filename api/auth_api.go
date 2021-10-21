package api

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
)

type JWTClaims struct {
	jwt.StandardClaims
	Username int    `json:"username"`
	Password string `json:"password"`
}

var (
	Secret     = "gwo" // salt
	ExpireTime = 3600  // 过期时间
)

// GetToken 获取token
func GetToken(claims *JWTClaims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenStr, err = token.SignedString([]byte(Secret))

	if err != nil {
		return "",errors.New("Server busy")
	}
	return
}

// VerifyToken 校验token
func VerifyToken(tokenStr string)(*JWTClaims,error)  {
	// todo 增加token校验方法
	return nil,nil
}
