package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gabrielmvnog/customer-service-fiap/src/application/ports"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/routes"
	mock_ports "github.com/gabrielmvnog/customer-service-fiap/tests/mocks/ports"
	"github.com/gin-gonic/gin"
	"go.uber.org/mock/gomock"

	"github.com/stretchr/testify/assert"
)

func SetupTestCustomerRoutes(customerService ports.ICustomerService) *gin.Engine {
	router := gin.Default()

	routes.SetupCustomerRoutes(router, customerService)

	return router
}

func TestCustomerRegistration(t *testing.T) {
	controller := gomock.NewController(t)

	mock_service := mock_ports.NewMockICustomerService(controller)
	mock_service.EXPECT().CreateCustomer(gomock.Any()).Return(&dtos.CreateCustomerResponse{Id: "1"})

	router := SetupTestCustomerRoutes(mock_service)

	recorder := httptest.NewRecorder()

	newCustomer := dtos.CreateCustomerRequest{Name: "testing 123"}
	customerJson, _ := json.Marshal(newCustomer)

	req, _ := http.NewRequest("POST", "/v1/customers", strings.NewReader(string(customerJson)))
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 201, recorder.Code)
}
