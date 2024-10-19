package server

import (
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/controllers"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/routes"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	db := pg.Connect(&pg.Options{
		User: "postgres",
	})
	defer db.Close()

	router.GET("/health", controllers.Healthcheck)

	routes.SetupCustomerRoutes(router)

	return router
}
