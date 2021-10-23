package routers

import (
	"GoToBetterLife/api"
	_ "GoToBetterLife/dal/models"
	"GoToBetterLife/handler"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {

	router := gin.Default()

	// 注册与登陆接口
	login := router.Group("/auth")
	{
		login.POST("/signUp",)
		login.POST("/login", api.Login)
		login.GET("/verify", api.Verify)
	}

	router.Use(handler.Authorize())
	// 这里分出一个user路由组，专门操作用户
	user := router.Group("/user")
	{
		user.GET("/", api.Users)
		user.POST("/", api.Update)
		user.PUT("/", api.UserSignUp)
		user.DELETE("/", api.Destroy)
	}

	// 日记相关接口
	dairy := router.Group("/dairy")
	{
		dairy.GET("/")
		dairy.PUT("/", api.AddNewDairy)
	}

	return router
}
