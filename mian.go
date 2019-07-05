package main

import (
	db "gin/ginDemo/database"
    router "gin/ginDemo/route"
)

func main() {
	defer db.SqlDB.Close()
	router := router.InitRouter()
	router.Run(":8000")
}
