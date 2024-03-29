package handler

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gitkoDev/pokemon-api/models"
	"github.com/gitkoDev/pokemon-api/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

func Test_AddPokemon(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		inputString     string
		mock            func(s *service.MockPokedex, pokemon string)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName:    "OK",
			inputString: `{"name":"Test", "type":["TestType"],"hp":40,"attack":40,"defense":40}`,
			mock: func(s *service.MockPokedex, pokemon string) {
				s.EXPECT().AddPokemon(models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40}).Return(1, nil)
			},
			wantStatusCode:  201,
			wantResposeBody: `{"id":1,"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}`,
		},
		{
			testName:    "Server error. Error",
			inputString: `{"name":"Test", "type":["TestType"],"hp":40,"attack":40,"defense":40}`,
			mock: func(s *service.MockPokedex, pokemon string) {
				s.EXPECT().AddPokemon(models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40}).Return(0, errors.New("internal service error"))
			},
			wantStatusCode:  500,
			wantResposeBody: `{"error":"internal service error"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock pokedex
			pokedex := service.NewMockPokedex(ctrl)

			// Prepare mock behavior
			testCase.mock(pokedex, testCase.inputString)

			// Init service struct with moth auth interface + handler
			service := &service.Service{Pokedex: pokedex}
			handler := NewHandler(service)

			// Init server
			r := chi.NewRouter()
			r.Post("/api/v1/pokemon", handler.addPokemon)

			// Init recorder and request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/pokemon", strings.NewReader(testCase.inputString))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.wantStatusCode, w.Code)
			assert.Equal(t, testCase.wantResposeBody, strings.Trim(w.Body.String(), "\n"))
		})
	}
}

func Test_GetAll(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		mock            func(s *service.MockPokedex)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName: "OK",
			mock: func(s *service.MockPokedex) {
				s.EXPECT().GetAll().Return([]models.Pokemon{
					{Id: 1, Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
					{Id: 2, Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
				}, nil)
			},
			wantStatusCode:  200,
			wantResposeBody: `[{"id":1,"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40},{"id":2,"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}]`,
		},
		{
			testName: "Server error. Error",
			mock: func(s *service.MockPokedex) {
				s.EXPECT().GetAll().Return([]models.Pokemon{}, errors.New("internal service error"))
			},
			wantStatusCode:  500,
			wantResposeBody: `{"error":"internal service error"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock pokedex
			pokedex := service.NewMockPokedex(ctrl)

			// Prepare mock behavior
			testCase.mock(pokedex)

			// Init service struct with moth auth interface + handler
			service := &service.Service{Pokedex: pokedex}
			handler := NewHandler(service)

			// Init server
			r := chi.NewRouter()
			r.Post("/api/v1/pokemon", handler.getAll)

			// Init recorder and request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/pokemon", nil)

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.wantStatusCode, w.Code)
			assert.Equal(t, testCase.wantResposeBody, strings.TrimSpace(w.Body.String()))
		})
	}

}

func Test_GetById(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		inputId         int
		mock            func(s *service.MockPokedex)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName: "OK",
			mock: func(s *service.MockPokedex) {
				s.EXPECT().GetById(1).Return(models.Pokemon{
					Id: 1, Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40,
				}, nil)
			},
			wantStatusCode:  200,
			wantResposeBody: `{"id":1,"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}`,
		},
		{
			testName: "Server error. Error",
			mock: func(s *service.MockPokedex) {
				s.EXPECT().GetById(1).Return(models.Pokemon{}, errors.New("internal service error"))
			},
			wantStatusCode:  500,
			wantResposeBody: `{"error":"internal service error"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock pokedex
			pokedex := service.NewMockPokedex(ctrl)

			// Prepare mock behavior
			testCase.mock(pokedex)

			// Init service struct with moth auth interface + handler
			service := &service.Service{Pokedex: pokedex}
			handler := NewHandler(service)

			// Init server
			r := chi.NewRouter()
			r.Post("/api/v1/pokemon/{id}", handler.getById)

			// Init recorder and request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/api/v1/pokemon/1", nil)

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.wantStatusCode, w.Code)
			assert.Equal(t, testCase.wantResposeBody, strings.TrimSpace(w.Body.String()))
		})
	}

}

func Test_UpdatePokemon(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		inputId         int
		inputBody       string
		mock            func(s *service.MockPokedex)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName:  "OK",
			inputBody: `{"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}`,
			mock: func(s *service.MockPokedex) {
				s.EXPECT().UpdatePokemon(models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40}, 1).
					Return(nil)
			},
			wantStatusCode:  200,
			wantResposeBody: `{"id":1,"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}`,
		},
		{
			testName:  "Server error. Error",
			inputBody: `{"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}`,
			mock: func(s *service.MockPokedex) {
				s.EXPECT().UpdatePokemon(models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40}, 1).
					Return(errors.New("internal server error"))
			},
			wantStatusCode:  500,
			wantResposeBody: `{"error":"internal server error"}`,
		},
		{
			testName:  "Server error. Error",
			inputBody: `{"name":"Test","type":["TestType"],"hp":40,"attack":40,"defense":40}`,
			mock: func(s *service.MockPokedex) {
				s.EXPECT().UpdatePokemon(models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40}, 1).
					Return(errors.New("internal server error"))
			},
			wantStatusCode:  500,
			wantResposeBody: `{"error":"internal server error"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock pokedex
			pokedex := service.NewMockPokedex(ctrl)

			// Prepare mock behavior
			testCase.mock(pokedex)

			// Init service struct with moth auth interface + handler
			service := &service.Service{Pokedex: pokedex}
			handler := NewHandler(service)

			// Init server
			r := chi.NewRouter()
			r.Put("/api/v1/pokemon/{id}", handler.updatePokemon)

			// Init recorder and request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPut, "/api/v1/pokemon/1", strings.NewReader(testCase.inputBody))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.wantStatusCode, w.Code)
			assert.Equal(t, testCase.wantResposeBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func Test_DeletePokemon(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		inputId         int
		mock            func(s *service.MockPokedex)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName: "OK",
			inputId:  1,
			mock: func(s *service.MockPokedex) {
				s.EXPECT().DeletePokemon(1).Return(nil)
			},
			wantStatusCode:  200,
			wantResposeBody: `{"message":"1 deleted"}`,
		}, {
			testName: "Server error. Error",
			inputId:  1,
			mock: func(s *service.MockPokedex) {
				s.EXPECT().DeletePokemon(1).Return(errors.New("internal server error"))
			},
			wantStatusCode:  500,
			wantResposeBody: `{"error":"internal server error"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock pokedex
			pokedex := service.NewMockPokedex(ctrl)

			// Prepare mock behavior
			testCase.mock(pokedex)

			// Init service struct with moth auth interface + handler
			service := &service.Service{Pokedex: pokedex}
			handler := NewHandler(service)

			// Init server
			r := chi.NewRouter()
			r.Delete("/api/v1/pokemon/{id}", handler.deletePokemon)

			// Init recorder and request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodDelete, "/api/v1/pokemon/1", nil)

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.wantStatusCode, w.Code)
			assert.Equal(t, testCase.wantResposeBody, strings.TrimSpace(w.Body.String()))
		})
	}
}
