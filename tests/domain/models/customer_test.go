package models

import (
	"errors"
	"testing"

	domain "github.com/gabrielmvnog/customer-service-fiap/src/domain/models"
	"github.com/stretchr/testify/assert"
)

func TestCustomerShouldSucced(t *testing.T) {
	customer := domain.Customer{
		Name:           "Xablau",
		Email:          "testing@gmail.com",
		DocumentNumber: "60563724722",
	}
	err := customer.Validate()

	assert.Nil(t, err)
}

func TestCustomerShouldReturnErrorToInvalidEmail(t *testing.T) {
	customer := domain.Customer{
		Name:           "Xablau",
		Email:          "testinggmail.com",
		DocumentNumber: "60563724722",
	}
	err := customer.Validate()

	assert.EqualError(t, errors.New("invalid email"), err.Error())
}

func TestCustomerShouldReturnErrorToInvalidName(t *testing.T) {
	customer := domain.Customer{
		Name:           "",
		Email:          "testing@gmail.com",
		DocumentNumber: "60563724722",
	}
	err := customer.Validate()

	assert.EqualError(t, errors.New("invalid name"), err.Error())
}

func TestCustomerShouldReturnErrorToInvalidDocumentNumber(t *testing.T) {
	customer := domain.Customer{
		Name:           "",
		Email:          "testing@gmail.com",
		DocumentNumber: "111",
	}
	err := customer.Validate()

	assert.EqualError(t, errors.New("invalid document number"), err.Error())
}
