package ports

import domain "github.com/gabrielmvnog/customer-service-fiap/src/domain/models"

type ICustomerRepository interface {
	InsertCustomer(*domain.Customer) int
	FindCustomerByDocumentNumber(string) (*domain.Customer, error)
}
