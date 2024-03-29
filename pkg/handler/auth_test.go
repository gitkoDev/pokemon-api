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

func Test_SignUp(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		inputString     string
		inputTrainer    models.Trainer
		mock            func(s *service.MockAuthorization, trainer models.Trainer)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName:     "OK",
			inputString:  `{"name": "TestTrainer", "password": "TestPassword"}`,
			inputTrainer: models.Trainer{Name: "TestTrainer", Password: "TestPassword"},

			mock: func(s *service.MockAuthorization, trainer models.Trainer) {
				s.EXPECT().CreateTrainer(models.Trainer{Name: "TestTrainer", Password: "TestPassword"}).Return(1, nil)
			},
			wantStatusCode:  http.StatusCreated,
			wantResposeBody: `{"id":1}`,
		},
		{
			testName:     "Server error. Error",
			inputString:  `{"name": "TestTrainer", "password": "TestPassword"}`,
			inputTrainer: models.Trainer{Name: "TestTrainer", Password: "TestPassword"},

			mock: func(s *service.MockAuthorization, trainer models.Trainer) {
				s.EXPECT().CreateTrainer(models.Trainer{Name: "TestTrainer", Password: "TestPassword"}).Return(0, errors.New("internal server error"))
			},
			wantStatusCode:  http.StatusInternalServerError,
			wantResposeBody: `{"error":"internal server error"}`,
		},
		{
			testName:     "Wrong Input. Error",
			inputString:  `{"name": "TestTrainer"}`,
			inputTrainer: models.Trainer{},

			mock: func(s *service.MockAuthorization, trainer models.Trainer) {

			},
			wantStatusCode:  http.StatusBadRequest,
			wantResposeBody: `{"error":"please provide valid trainer name and password"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		t.Run(testCase.testName, func(t *testing.T) {
			// Init mock auth
			auth := service.NewMockAuthorization(ctrl)

			// Prepare mock behavior
			testCase.mock(auth, testCase.inputTrainer)

			// Init service struct with moth auth interface + handler
			service := &service.Service{Authorization: auth}
			handler := NewHandler(service)

			// Init server
			r := chi.NewRouter()
			r.Post("/auth/sign-up", handler.signUp)

			// Init recorder and request
			w := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodPost, "/auth/sign-up", strings.NewReader(testCase.inputString))

			// Perform request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.wantStatusCode, w.Code)
			assert.Equal(t, testCase.wantResposeBody, strings.TrimSpace(w.Body.String()))
		})
	}
}

func Test_SignIn(t *testing.T) {
	// Init mock controller
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Setup tests
	tests := []struct {
		testName        string
		inputName       string
		inputPassword   string
		inputString     string
		mock            func(s *service.MockAuthorization, name string, password string)
		wantStatusCode  int
		wantResposeBody string
	}{
		{
			testName:      "OK",
			inputName:     "TestTrainer",
			inputPassword: "TestPassword",
			inputString:   `{"name":"TestTrainer", "password": "TestPassword"}`,
			mock: func(s *service.MockAuthorization, name, password string) {
				s.EXPECT().GetTrainer("TestTrainer", "TestPassword").Return(models.Trainer{Name: "TestTrainer", Password: "TestPassword"}, nil)
				s.EXPECT().GenerateToken("TestTrainer", "TestPassword").Return("TestToken", nil)
			},
			wantStatusCode:  200,
			wantResposeBody: `{"token":"TestToken"}`,
		},
		{
			testName:        "Wrong input. Error",
			inputName:       "TestTrainer",
			inputPassword:   "",
			inputString:     `{"name":"TestTrainer"}`,
			mock:            func(s *service.MockAuthorization, name, password string) {},
			wantStatusCode:  400,
			wantResposeBody: `{"error":"please provide valid trainer name and password"}`,
		},
		{
			testName:      "Server error. Error",
			inputName:     "TestTrainer",
			inputPassword: "TestPassword",
			inputString:   `{"name":"TestTrainer", "password": "TestPassword"}`,
			mock: func(s *service.MockAuthorization, name, password string) {
				s.EXPECT().GetTrainer("TestTrainer", "TestPassword").Return(models.Trainer{}, errors.New("internal server error"))
				// Server error, return from function before generating token
			},
			wantStatusCode:  400,
			wantResposeBody: `{"error":"internal server error"}`,
		},
	}

	// Run tests
	for _, testCase := range tests {
		// Init mock auth
		auth := service.NewMockAuthorization(ctrl)

		// Prepare mock behavior
		testCase.mock(auth, testCase.inputName, testCase.inputPassword)

		// Init service struct with moth auth interface + handler
		service := &service.Service{Authorization: auth}
		handler := NewHandler(service)

		// Init server
		r := chi.NewRouter()
		r.Get("/auth/sign-in", handler.signIn)

		// Init recorder and request
		w := httptest.NewRecorder()
		req := httptest.NewRequest(http.MethodGet, "/auth/sign-in", strings.NewReader(testCase.inputString))

		// Perform request
		r.ServeHTTP(w, req)

		assert.Equal(t, testCase.wantStatusCode, w.Code)
		assert.Equal(t, testCase.wantResposeBody, strings.TrimSpace(w.Body.String()))

	}
}
