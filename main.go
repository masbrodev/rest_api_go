package main

import (
	"rest_api/db"
	"rest_api/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
