package main

import (
	"github.com/gabrielmvnog/customer-service-fiap/src/config"
	"github.com/gabrielmvnog/customer-service-fiap/src/server"
)

func main() {
	db := config.SetupDatabase()

	router := server.SetupRouter(db)

	router.Run()
}
