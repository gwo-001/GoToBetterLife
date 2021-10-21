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

const (
	ServerBusy = "Server busy"
)

// GetToken 获取token
func GetToken(claims *JWTClaims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodES256, claims)
	tokenStr, err = token.SignedString([]byte(Secret))

	if err != nil {
		return "",err
	}
	return
}

// VerifyToken 校验token
func VerifyToken(tokenStr string)(*JWTClaims,error)  {
	token,err:=jwt.ParseWithClaims(tokenStr,&JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret),nil
	})
	if err != nil {
		return nil, errors.New(ServerBusy)
	}
	claims,ok:=token.Claims.(*JWTClaims)
	if !ok {
		return nil,errors.New(ServerBusy)
	}
	if err:=token.Claims.Valid();err!=nil {
		return nil, errors.New(ServerBusy)
	}
	return claims,nil
}
