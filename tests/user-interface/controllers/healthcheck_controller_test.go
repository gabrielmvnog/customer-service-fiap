package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gabrielmvnog/customer-service-fiap/src/server"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	db, _, _ := sqlmock.New()
	router := server.SetupRouter(db)

	recorder := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/health", nil)
	router.ServeHTTP(recorder, req)

	assert.Equal(t, 200, recorder.Code)
}
