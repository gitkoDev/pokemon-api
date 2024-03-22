package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gitkoDev/pokemon-api/models"
	"github.com/stretchr/testify/assert"
)

func Test_CreateTrainer(t *testing.T) {
	// Init mock DB
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDb.Close()

	// Init repo interface with necessary method
	repo := &AuthPostgres{db: mockDb}

	// Setup tests
	tests := []struct {
		testName  string
		inputBody models.Trainer
		want      int
		wantError bool
		mock      func()
	}{
		{
			testName:  "OK",
			inputBody: models.Trainer{Name: "TestTrainer", Password: "TestHash"},
			want:      1,
			wantError: false,
			mock: func() {
				mock.NewRows([]string{"id", "name", "password_hash"})

				// User with this name doesn't exist, we can insert
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM pokemon_trainers WHERE name = $1")).WithArgs("TestTrainer").WillReturnError(sql.ErrNoRows)

				// Insert new user
				idRow := mock.NewRows([]string{"id"}).AddRow("1")
				mock.ExpectQuery(regexp.QuoteMeta(`INSERT INTO pokemon_trainers (NAME, PASSWORD_HASH) VALUES ($1, $2) RETURNING id`)).WithArgs("TestTrainer", "TestHash").WillReturnRows(idRow)
			},
		},
		{
			testName:  "Trainer with the same name exists. Error",
			inputBody: models.Trainer{Name: "TestTrainer", Password: "TestHash"},
			want:      0,
			wantError: true,
			mock: func() {
				row := mock.NewRows([]string{"id", "name", "password_hash"}).
					AddRow("1", "TestTrainer", "TestHash")

				// User with this name exists, we exit with error
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM pokemon_trainers WHERE name = $1")).WithArgs("TestTrainer").WillReturnRows(row)

			},
		},
	}

	for _, testCase := range tests {
		testCase.mock()

		got, err := repo.CreateTrainer(testCase.inputBody)

		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, got)
		}
	}
}

func Test_GetTrainer(t *testing.T) {
	// Init mock DB
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDb.Close()

	// Init repo interface with necessary method
	repo := &AuthPostgres{db: mockDb}

	// Setup tests
	tests := []struct {
		testName  string
		inputName string
		inputHash string
		want      models.Trainer
		wantError bool
		mock      func()
	}{
		{
			testName:  "OK",
			inputName: "TestTrainer",
			inputHash: "TestHash",
			want:      models.Trainer{Id: 1, Name: "TestTrainer", Password: "TestHash"},
			wantError: false,
			mock: func() {
				row := mock.NewRows([]string{"id"}).
					AddRow("1")

				// Trainer found, return id
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM pokemon_trainers WHERE name = $1 AND password_hash = $2")).WithArgs("TestTrainer", "TestHash").WillReturnRows(row)

			},
		},
		{
			testName:  "Trainer not found. Error",
			inputName: "TestTrainer",
			inputHash: "TestHash",
			wantError: true,
			mock: func() {
				// Trainer not found, exit
				mock.ExpectQuery(regexp.QuoteMeta("SELECT id FROM pokemon_trainers WHERE name = $1 AND password_hash = $2")).WithArgs("TestTrainer", "TestHash").WillReturnError(sql.ErrNoRows)
			},
		},
	}

	for _, testCase := range tests {
		testCase.mock()

		got, err := repo.GetTrainer(testCase.inputName, testCase.inputHash)

		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, got)
		}
	}
}
