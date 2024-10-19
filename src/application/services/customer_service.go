package application

import (
	"strconv"

	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	domain "github.com/gabrielmvnog/customer-service-fiap/src/domain/models"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"
)

type CustomerService struct {
	Repository ports.ICustomerRepository
}

func NewCustomerService(repository ports.ICustomerRepository) ports.ICustomerService {
	return &CustomerService{
		Repository: repository,
	}
}

func (customerService CustomerService) CreateCustomer(newCustomer *dtos.CreateCustomerRequest) *dtos.CreateCustomerResponse {
	customer := domain.Customer{
		Name:           newCustomer.Name,
		Email:          newCustomer.Name,
		DocumentNumber: newCustomer.Name,
	}
	customerId := customerService.Repository.InsertCustomer(&customer)

	return &dtos.CreateCustomerResponse{Id: strconv.Itoa(customerId)}
}
