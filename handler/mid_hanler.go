package handler

import (
	"GoToBetterLife/api"
	"github.com/gin-gonic/gin"
)

var token string

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		cookies := c.Request.Cookies()
		// 如果请求header中没有cookie，直接拦截请求
		if cookies == nil || len(cookies) == 0 {
			c.Abort()
			return
		}
		for _, v := range cookies {
			if v.Name == "token" {
				// 先校验token有效性
				if  _, err := api.VerifyToken(v.Value);err!=nil{
					c.Abort()
				}
				break
			}
			continue
		}
		c.Next()
	}

}
