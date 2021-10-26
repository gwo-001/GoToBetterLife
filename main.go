package main

import (
	"GoToBetterLife/dal/database"
	"GoToBetterLife/dal/models"
	"GoToBetterLife/routers"
)

func main() {
	database.InitMysql()
	database.InitRedis()
	defer database.Db.Close()
	models.InitAllTables()
	router := routers.InitRouter()
	router.Run(":8080")
}
