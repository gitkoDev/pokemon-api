package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gitkoDev/pokemon-api/models"
	"github.com/gitkoDev/pokemon-api/pkg/service"
)

var mockTrainer = models.Trainer{Name: "test", Password: "test"}

var services = &service.Service{}
var handler = &Handler{}

var token string

func init() {

	// log.Println(token)
}

func TestPing(t *testing.T) {
	// Init request
	req, err := http.NewRequest("GET", "/health", nil)
	req = req.WithContext(context.WithValue(req.Context(), "Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA4MDc0NjksImlhdCI6MTcxMDc2NDI2OSwidHJhaW5lcl9pZCI6MX0.gG7UOZksiscT618x1uan5FA8F5YbDDivy2z17ouRcIY"))
	if err != nil {
		t.Fatal(err)
	}

	// Init request recorder and handler
	rr := httptest.NewRecorder()
	hnd := handler
	hnd.InitRoutes().ServeHTTP(rr, req)

	req.Header.Set("Authorization", "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MTA4MDc0NjksImlhdCI6MTcxMDc2NDI2OSwidHJhaW5lcl9pZCI6MX0.gG7UOZksiscT618x1uan5FA8F5YbDDivy2z17ouRcIY")
	// log.Println(req.Header.Values("Authorization"), "here")

	// Check for status code
	if status := rr.Code; status != 200 {
		t.Fatalf("wrong status code: expected: %v, got: %v", http.StatusOK, status)
	}

	// Check for response
	expectedBody := "Pokemon API v.1.0"
	if rr.Body.String() != expectedBody {
		t.Fatalf("wrong body, expected: %s, got: %s", expectedBody, rr.Body.String())
	}
}

func TestAddPokemon(t *testing.T) {
	bd := "TestPokemon"

	// Init request
	req, err := http.NewRequest("POST", "/api/v1/pokemon", strings.NewReader(bd))
	if err != nil {
		t.Fatal(err)
	}

	// Init request recorder and handler
	rr := httptest.NewRecorder()
	hnd := handler
	hnd.InitRoutes().ServeHTTP(rr, req)

	// Check for status code
	if status := rr.Code; status != http.StatusCreated {
		t.Fatalf("wrong status code, expected: %v, got: %v", http.StatusCreated, status)
	}

	// Check for responce
	if rr.Body.String() == bd {
		t.Fatalf("wrong body, expected: %s, got: %s", bd, rr.Body.String())
	}

}
