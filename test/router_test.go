package test

import (
	"app/server"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRoute(t *testing.T) {
	router := server.NewRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ishealthy", nil)
	router.ServeHTTP(w, req)
	if w.Code != 200 {
		t.Errorf("Response code of /ishealthy should be 200")
	}
}
