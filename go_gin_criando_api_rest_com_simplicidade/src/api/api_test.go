package api_test

import (
	"go_gin_criando_api_rest_com_simplicidade/src/api"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInit(t *testing.T) {
	r := api.Init()

	req, _ := http.NewRequest("GET", "/health", nil)
	resp := httptest.NewRecorder()

	r.ServeHTTP(resp, req)

	if resp.Code != http.StatusOK {
		t.Errorf("expected status %d, but it came %d", http.StatusOK, resp.Code)
	}
}
