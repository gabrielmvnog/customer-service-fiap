package controllers

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	domain "github.com/gabrielmvnog/customer-service-fiap/src/domain/models"
	"github.com/gabrielmvnog/customer-service-fiap/src/infrastructure"
	"github.com/stretchr/testify/assert"

	_ "github.com/lib/pq"
)

func TestCustomerInsertion(t *testing.T) {
	db, mock, _ := sqlmock.New()

	repository := infrastructure.NewCustomerRepository(db)
	customer := domain.Customer{
		Name:           "test",
		Email:          "test@test.com",
		DocumentNumber: "12345678",
	}

	mock.ExpectQuery("INSERT INTO customers").
		WithArgs("test", "test@test.com", "12345678").
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	customerId := repository.InsertCustomer(&customer)

	assert.Equal(t, 1, customerId)
}

func TestCustomerFindByDocumentNumber(t *testing.T) {
	db, mock, _ := sqlmock.New()

	repository := infrastructure.NewCustomerRepository(db)

	mock.ExpectQuery("SELECT name, email, documentNumber FROM customers").
		WithArgs("12345678910").
		WillReturnRows(sqlmock.NewRows([]string{"name", "email", "documentNumber"}).
			AddRow("Xablau", "test@gmail.com", "12345678910"))

	cusotmer, _ := repository.FindCustomerByDocumentNumber("12345678910")

	assert.Equal(t, "Xablau", cusotmer.Name)
}
