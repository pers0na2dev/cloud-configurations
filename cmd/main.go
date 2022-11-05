package main

import (
	"cloud-configurations/internal/database"
	"cloud-configurations/internal/environment"
	"cloud-configurations/internal/http_router"
)

func init() {
	environment.Initialize()

	database.DB = database.Database()
	database.Migration()
}

func main() {
	if err := http_router.Router().Listen(":4000"); err != nil {
		panic(err.Error())
	}
}
