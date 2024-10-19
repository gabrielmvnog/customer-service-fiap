package routes

import (
	application "github.com/gabrielmvnog/customer-service-fiap/src/application/services"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine) {
	var customerService application.CustomerService = *application.NewCustomerService()
	var customerController controllers.CustomerController = *controllers.NewCustomerController(customerService)

	var routerGroup gin.RouterGroup = *router.Group("/v1")
	{
		routerGroup.POST("/customers", customerController.CustomerRegistration)
	}
}
