package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestPing(t *testing.T) {
	req, err := http.NewRequest("GET", "/ping", nil)

	if err != nil {
		t.Fatal("Error running request to /ping:", err)
	}

	rec := httptest.NewRecorder()
	handler := http.HandlerFunc(Ping)

	handler.ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Fatalf("/ping failed. Expected status code, expected %d, got %d", http.StatusOK, status)
	}
}
