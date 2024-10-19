package routes

import (
	"database/sql"

	application "github.com/gabrielmvnog/customer-service-fiap/src/application/services"
	"github.com/gabrielmvnog/customer-service-fiap/src/infrastructure"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, db *sql.DB) {
	var customerRepository infrastructure.CustomerRepository = *infrastructure.NewCustomerRepository(db)
	var customerService application.CustomerService = *application.NewCustomerService(customerRepository)
	var customerController controllers.CustomerController = *controllers.NewCustomerController(customerService)

	var routerGroup gin.RouterGroup = *router.Group("/v1")
	{
		routerGroup.POST("/customers", customerController.CustomerRegistration)
	}
}
