package router

import (
	. "GoToBetterLife/api/apis"
	_ "GoToBetterLife/api/models"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()
	// 这里分出一个user路由组，专门操作用户
	user := router.Group("/user")
	{
		user.GET("/", Users)
		user.POST("/", Update)
		user.PUT("/", Store)
		user.DELETE("/", Destroy)
	}

	dairy:=router.Group("/dairy")
	{
		dairy.GET("/")
	}

	return router
}
