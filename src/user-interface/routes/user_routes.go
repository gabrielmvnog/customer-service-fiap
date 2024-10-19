package routes

import (
	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/controllers"

	"github.com/gin-gonic/gin"
)

func SetupCustomerRoutes(router *gin.Engine, customerService ports.ICustomerService) {
	var customerController *controllers.CustomerController = controllers.NewCustomerController(customerService)

	var routerGroup gin.RouterGroup = *router.Group("/v1")
	{
		routerGroup.POST("/customers", customerController.CustomerRegistration)
	}
}
