package server

import (
	"database/sql"

	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/controllers"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/health", controllers.Healthcheck)

	routes.SetupCustomerRoutes(router, db)

	return router
}
