package main

import (
	"main/db"
	"main/router"
)

func main() {
	db.InitPostgresDB()
	r := router.InitRouter()
	r.Run(":3333")
}
