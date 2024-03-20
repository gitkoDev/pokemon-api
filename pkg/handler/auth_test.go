package handler

import (
	"encoding/json"
	"testing"

	"github.com/gitkoDev/pokemon-api/models"
	"github.com/gitkoDev/pokemon-api/pkg/service"
)

type mockAuthService struct {
	authService service.AuthService
}

func (s *mockAuthService) SignUp() (bool, error) {
	inputBody := `{"name": "Test", "password": "test"}`
	var unmarchalledInput models.Trainer

	err := json.Unmarshal([]byte(inputBody), &unmarchalledInput)
	if err != nil {
		return false, err
	}

	return true, nil

}

func Test_SignUp(t *testing.T) {
	testTable := []struct {
		testCase   string
		inputBody  string
		outputBody models.Trainer
	}{
		{
			testCase:   "Should NOT give error",
			inputBody:  `{"name": "Test", "password": "test"}`,
			outputBody: models.Trainer{Name: "Test", Password: "test"},
		},
	}

	for _, test := range testTable {
		t.Run(test.testCase, func(t *testing.T) {

		})
	}

	// mockAuthService := &mockAuthService{}
	// if _, err := mockAuthService.SignUp(); err != nil {
	// 	t.Error("error", err)
	// } else {
	// 	t.Log("not error")
	// }
}
