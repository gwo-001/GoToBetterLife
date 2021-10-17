package main

import (
	"GoToBetterLife/dal/database"
	"GoToBetterLife/dal/models"
	"GoToBetterLife/routers"
)

func main() {
	database.Init()
	defer database.Db.Close()
	models.InitAllTables()
	router := routers.InitRouter()
	router.Run(":8080")
}
