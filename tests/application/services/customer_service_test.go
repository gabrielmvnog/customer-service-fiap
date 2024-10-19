package controllers

import (
	"testing"

	application "github.com/gabrielmvnog/customer-service-fiap/src/application/services"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"
	mock_ports "github.com/gabrielmvnog/customer-service-fiap/tests/mocks/ports"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRegistration(t *testing.T) {
	controller := gomock.NewController(t)

	mock_repository := mock_ports.NewMockICustomerRepository(controller)
	mock_repository.EXPECT().InsertCustomer(gomock.Any()).Return(1)

	service := application.NewCustomerService(mock_repository)

	newCustomer := dtos.CreateCustomerRequest{}
	customerResponse := service.CreateCustomer(&newCustomer)

	assert.Equal(t, "1", customerResponse.Id)
}
