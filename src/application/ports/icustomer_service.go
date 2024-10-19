package ports

import "github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"

type ICustomerService interface {
	CreateCustomer(*dtos.CreateCustomerRequest) (*dtos.CreateCustomerResponse, error)
	FindCustomerByDocumentNumber(string) (*dtos.FindCustomerResponse, error)
}
