package controllers

import (
	"testing"

	application "github.com/gabrielmvnog/customer-service-fiap/src/application/services"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"
	mock_ports "github.com/gabrielmvnog/customer-service-fiap/tests/mocks/ports"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func TestCustomerRegistrationShouldSucceed(t *testing.T) {
	controller := gomock.NewController(t)

	mock_repository := mock_ports.NewMockICustomerRepository(controller)
	mock_repository.EXPECT().InsertCustomer(gomock.Any()).Return(1)

	service := application.NewCustomerService(mock_repository)

	newCustomer := dtos.CreateCustomerRequest{
		Name:           "Xablau",
		Email:          "xablau@gmail.com",
		DocumentNumber: "904.878.540-55",
	}
	customerResponse, _ := service.CreateCustomer(&newCustomer)

	assert.Equal(t, "1", customerResponse.Id)
}

func TestCustomerRegistrationShouldReturnError(t *testing.T) {
	controller := gomock.NewController(t)

	mock_repository := mock_ports.NewMockICustomerRepository(controller)

	service := application.NewCustomerService(mock_repository)

	newCustomer := dtos.CreateCustomerRequest{}
	_, err := service.CreateCustomer(&newCustomer)

	assert.EqualError(t, errors.New("invalid customer data"), err.Error())
}
