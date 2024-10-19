package infrastructure

import (
	"database/sql"
	"log"

	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	domain "github.com/gabrielmvnog/customer-service-fiap/src/domain/models"
)

type CustomerRepository struct {
	Database *sql.DB
}

func NewCustomerRepository(db *sql.DB) ports.ICustomerRepository {
	return &CustomerRepository{
		Database: db,
	}
}

func (customerRepository CustomerRepository) InsertCustomer(newCustomer *domain.Customer) int {
	var customerId int

	err := customerRepository.Database.QueryRow(
		`INSERT INTO customers(name, email, documentNumber) VALUES ($1, $2, $3) RETURNING id`,
		newCustomer.Name,
		newCustomer.Email,
		newCustomer.DocumentNumber,
	).Scan(&customerId)
	if err != nil {
		log.Fatal(err)
	}

	return customerId
}
