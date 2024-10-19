package controllers

import (
	"net/http"

	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"

	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	CustomerService ports.ICustomerService
}

func NewCustomerController(customerService ports.ICustomerService) *CustomerController {
	return &CustomerController{
		CustomerService: customerService,
	}
}

func (customerController CustomerController) CustomerRegistration(context *gin.Context) {
	var newCustomer dtos.CreateCustomerRequest

	if err := context.BindJSON(&newCustomer); err != nil {
		return
	}

	var response *dtos.CreateCustomerResponse = customerController.CustomerService.CreateCustomer(&newCustomer)

	context.IndentedJSON(http.StatusCreated, &response)
}
