package main

import (
	"GoToBetterLife/api/database"
	"GoToBetterLife/api/models"
	"GoToBetterLife/api/routers"
)

func main() {
	database.Init()
	defer database.Db.Close()
	models.InitAllTables()
	router := routers.InitRouter()
	router.Run(":8080")
}
