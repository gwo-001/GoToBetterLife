package api

import (
	"GoToBetterLife/util"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type JWTClaims struct {
	jwt.StandardClaims
	Username string `json:"username"`
	Password string `json:"password"`
}

var (
	Secret     = []byte("gwo") // salt
	ExpireTime = 3600          // 过期时间
)

const (
	ServerBusy = "Server busy"
)

// Login 登陆
func Login(c *gin.Context) {
	claims := &JWTClaims{}
	err := c.BindJSON(claims)
	claims.IssuedAt = time.Now().Unix()
	claims.ExpiresAt = time.Now().Add(time.Second * time.Duration(ExpireTime)).Unix()
	signedToken, err := GetToken(claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.One,
			util.Message:    util.Fail,
			"token":         "",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		util.StatusCode: util.One,
		util.Message:    util.Success,
		"token":         signedToken,
	})
}

// Verify 校验token
func Verify(c *gin.Context) {
	tokenStr := c.Request.FormValue("token")
	claim, err := VerifyToken(tokenStr)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			util.StatusCode: util.Zero,
			util.Message:    util.Fail,
			"data":          "invalid token",
		})
		return
	}
	data := "valid token, " + claim.Username
	c.JSON(http.StatusOK, gin.H{
		util.StatusCode: util.One,
		util.Message:    util.Success,
		util.Data:       data,
	})
}

// GetToken 获取token
func GetToken(claims *JWTClaims) (tokenStr string, err error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenStr, err = token.SignedString(Secret)

	if err != nil {
		return "", err
	}
	return
}

// VerifyToken 校验token
func VerifyToken(tokenStr string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
		return Secret, nil
	})
	if err != nil {
		return nil, errors.New(ServerBusy)
	}
	claims, ok := token.Claims.(*JWTClaims)
	if !ok {
		return nil, errors.New(ServerBusy)
	}
	if err := token.Claims.Valid(); err != nil {
		return nil, errors.New(ServerBusy)
	}
	return claims, nil
}
