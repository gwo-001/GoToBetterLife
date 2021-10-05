package router

import (
	"GoToBetterLife/api/apis"
	_ "GoToBetterLife/api/models"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine {
	router:=gin.Default()
	router.GET("/users",)
	router.POST("/user",)
	router.PUT("/user/:id",apis.Store)
	router.DELETE("/user/:id",apis.Destroy)

	return router
}
