package handler

import (
	"GoToBetterLife/api"
	"fmt"
	"github.com/gin-gonic/gin"
)

var token string

func Authorize() gin.HandlerFunc {
	return func(c *gin.Context) {
		//headerMap := c.Request.Header
		cookies := c.Request.Cookies()
		for _, v := range cookies {
			if v.Name == "token" {
				// 先校验token有效性
				jwtClaim, err := api.VerifyToken(v.Value)
				if err != nil {
					c.Abort()
					return
				}
				// 这里校验用户密码的正确性
				//username:=jwtClaim.Username
				//password:=jwtClaim.Password

				fmt.Println(jwtClaim)
				c.Next()
			}
		}

		//cookie := headerMap.Get("Cookie")
		//if !util.HasLength(cookie) {
		//	fmt.Println("没有cookie")
		//	c.Abort()
		//}
		//
		//arr := strings.Split(cookie, ";")
		//flag := false
		//for i,v := range arr {
		//	tmpArr := strings.Split(v, "=")
		//	if tmpArr[0] == "token" {
		//		_, err := api.VerifyToken(arr[i+1])
		//		if err == nil {
		//			flag=true
		//		}
		//	}
		//	continue
		//}
		//if flag {
		//	c.Next()
		//	return
		//}
		//c.Abort()
		//for k, v := range c.Request.Header {
		//	if "token"== k{
		//		fmt.Println(v)
		//		jwtClaim,err:=api.VerifyToken(string(v))
		//		if err == nil {
		//			c.Abort()
		//		}
		//
		//		c.Next()
		//	}
		//	continue
		//}
	}

}
