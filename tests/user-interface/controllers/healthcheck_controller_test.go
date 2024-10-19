package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gabrielmvnog/customer-service-fiap/src/server"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	router := server.SetupRouter()

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
