package router

import (
	. "GoToBetterLife/api/apis"
	_ "GoToBetterLife/api/models"
	"github.com/gin-gonic/gin"
)

func InitRouter()*gin.Engine {

	router:=gin.Default()
	router.GET("/users",Users)
	router.POST("/user",Update)
	router.PUT("/user/:id",Store)
	router.DELETE("/user/:id",Destroy)

	return router
}
