package application

import (
	"errors"
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

func (customerService CustomerService) CreateCustomer(newCustomer *dtos.CreateCustomerRequest) (*dtos.CreateCustomerResponse, error) {
	customer := domain.Customer{
		Name:           newCustomer.Name,
		Email:          newCustomer.Email,
		DocumentNumber: newCustomer.DocumentNumber,
	}

	err := customer.Validate()
	if err != nil {
		return &dtos.CreateCustomerResponse{}, errors.New("invalid customer data")
	}

	customerId := customerService.Repository.InsertCustomer(&customer)

	return &dtos.CreateCustomerResponse{Id: strconv.Itoa(customerId)}, nil
}
