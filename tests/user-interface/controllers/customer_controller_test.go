package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gabrielmvnog/customer-service-fiap/src/server"
	"github.com/gabrielmvnog/customer-service-fiap/src/user-interface/dtos"

	"github.com/stretchr/testify/assert"
)

func TestCustomerRegistration(t *testing.T) {
	db, _, _ := sqlmock.New()
	router := server.SetupRouter(db)

	recorder := httptest.NewRecorder()

	newCustomer := dtos.CreateCustomerRequest{Name: "testing 123"}
	customerJson, _ := json.Marshal(newCustomer)

	req, _ := http.NewRequest("POST", "/v1/customers", strings.NewReader(string(customerJson)))
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 201, recorder.Code)
}
