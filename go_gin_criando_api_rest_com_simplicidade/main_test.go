package main_test

import (
	"go_gin_criando_api_rest_com_simplicidade/src/api"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestListingAllHandlerStudents(t *testing.T) {
	r := api.Init()

	req, _ := http.NewRequest(http.MethodGet, "/v1/students", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)
}
