package handler

import (
	"GoToBetterLife/api"
	"GoToBetterLife/dal/models"
	"github.com/gin-gonic/gin"
)


// Authorize 鉴权中间件
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
				claim, err := api.VerifyToken(v.Value)
				if err != nil {
					c.Abort()
				}
				models.Username = claim.Username
				break
			}
			continue
		}
		c.Next()
	}

}
