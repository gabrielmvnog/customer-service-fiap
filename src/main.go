package main

import "github.com/gabrielmvnog/customer-service-fiap/src/server"

func main() {
	router := server.SetupRouter()

	router.Run()
}
