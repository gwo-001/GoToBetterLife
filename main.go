package main

import (
	orm "GoToBetterLife/api/database"
	. "GoToBetterLife/api/router"
)

func main() {
	orm.Init()
	defer orm.Db.Close()
	router := InitRouter()
	router.Run(":8080")
}
