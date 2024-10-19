package application

import (
	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"
)

type CustomerService struct {
	Repository ports.ICustomerRepository
}

func NewCustomerService(repository ports.ICustomerRepository) *CustomerService {
	return &CustomerService{
		Repository: repository,
	}
}

func (customerService CustomerService) CreateCustomer(newCustomer dtos.CreateCustomerRequest) dtos.CreateCustomerResponse {
	return dtos.CreateCustomerResponse{Id: "1"}
}
