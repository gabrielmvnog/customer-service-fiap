package server

import (
	"database/sql"

	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	application "github.com/gabrielmvnog/customer-service-fiap/src/application/services"
	"github.com/gabrielmvnog/customer-service-fiap/src/infrastructure"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/controllers"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/routes"

	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	router := gin.Default()

	router.GET("/health", controllers.Healthcheck)

	var customerRepository ports.ICustomerRepository = infrastructure.NewCustomerRepository(db)
	var customerService ports.ICustomerService = application.NewCustomerService(customerRepository)
	routes.SetupCustomerRoutes(router, customerService)

	return router
}
