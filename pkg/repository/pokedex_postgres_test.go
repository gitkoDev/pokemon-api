package repository

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gitkoDev/pokemon-api/models"
	"github.com/lib/pq"
	"github.com/stretchr/testify/assert"
)

func Test_AddPokemon(t *testing.T) {
	// Init mock db
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDB.Close()

	// Init repo interface with necessary method
	repo := &PokedexPostgres{db: mockDB}

	// Setup tests
	tests := []struct {
		testName  string
		mockFunc  func()
		inputBody models.Pokemon
		want      int
		wantError bool
	}{
		{
			testName:  "OK",
			inputBody: models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

				mock.ExpectQuery(`INSERT INTO pokemon`).WithArgs("Test", pq.Array([]string{"TestType"}), 40, 40, 40).WillReturnRows(rows)
			},
			want:      1,
			wantError: false,
		},
		{
			testName:  "Empty fields. Error",
			inputBody: models.Pokemon{Name: "", PokemonType: []string{"TestPower"}, Hp: 40, Attack: 40, Defense: 40},
			mockFunc: func() {
				rows := sqlmock.NewRows([]string{"id"}).AddRow(1)

				mock.ExpectQuery(`INSERT INTO pokemon`).WithArgs("", pq.Array([]string{"TestType"}), 40, 40, 40).WillReturnRows(rows)
			},
			want:      0,
			wantError: true,
		},
	}

	// Run tests
	for _, testCase := range tests {
		testCase.mockFunc()

		got, err := repo.AddPokemon(testCase.inputBody)

		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, got)
		}

	}

}

func Test_GetAll(t *testing.T) {
	// Init mock DB
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDB.Close()

	// Init repo interface with necessary method
	repo := &PokedexPostgres{db: mockDB}

	// Setup tests
	tests := []struct {
		testName  string
		mockFunc  func()
		want      []models.Pokemon
		wantError bool
	}{
		{
			testName: "OK",
			mockFunc: func() {
				rows := mock.NewRows([]string{"id", "name", "type", "hp", "attack", "defense"}).
					AddRow(1, "TestPokemon", pq.Array([]string{"TestType"}), 40, 40, 40).
					AddRow(2, "TestPokemon", pq.Array([]string{"TestType"}), 40, 40, 40).
					AddRow(3, "TestPokemon", pq.Array([]string{"TestType"}), 40, 40, 40)

				mock.ExpectQuery(`SELECT id, name, type, hp, attack, defense FROM pokemon`).WillReturnRows(rows)
			},
			want: []models.Pokemon{
				{Id: 1, Name: "TestPokemon", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
				{Id: 2, Name: "TestPokemon", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
				{Id: 3, Name: "TestPokemon", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			},
			wantError: false,
		},
	}

	// Run tests
	for _, testCase := range tests {
		testCase.mockFunc()

		got, err := repo.GetAll()

		// Check for errors
		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, got)
		}

	}
}

func Test_GetById(t *testing.T) {
	// Init mock DB
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDb.Close()

	// Init repo interface with necessary method
	repo := &PokedexPostgres{db: mockDb}

	// Setup tests
	tests := []struct {
		testName  string
		mockFunc  func()
		inputId   int
		want      models.Pokemon
		wantError bool
	}{
		{
			testName: "Ok",
			mockFunc: func() {
				rows := mock.NewRows([]string{"id", "name", "type", "hp", "attack", "defense"}).
					AddRow(1, "TestPokemon", pq.Array([]string{"TestType"}), 40, 40, 40)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, type, hp, attack, defense FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnRows(rows)
			},
			inputId:   1,
			want:      models.Pokemon{Id: 1, Name: "TestPokemon", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			wantError: false,
		},
		{
			testName: "Not found. Error",
			mockFunc: func() {

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id, name, type, hp, attack, defense FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			inputId:   1,
			want:      models.Pokemon{Id: 1, Name: "TestPokemon", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			wantError: true,
		},
	}

	// Run tests
	for _, testCase := range tests {
		testCase.mockFunc()

		got, err := repo.GetById(testCase.inputId)

		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, testCase.want, got)
		}
	}
}

func Test_UpdatePokemon(t *testing.T) {
	// Init mock DB
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDb.Close()

	// Init repo interface with necessary method
	repo := &PokedexPostgres{db: mockDb}

	// Setup tests
	tests := []struct {
		testName          string
		mockFunc          func()
		inputBody         models.Pokemon
		originalPokemonId int
		wantError         bool
	}{
		{
			testName: "OK",
			mockFunc: func() {
				row := mock.NewRows([]string{"id"}).
					AddRow(1)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnRows(row)

				mock.ExpectExec(regexp.QuoteMeta(`UPDATE pokemon
				SET name = $1, type = $2, hp = $3, attack = $4, defense = $5
				WHERE id = $6
				`)).WithArgs("TestPokemon", pq.Array([]string{"TestType"}), 40, 40, 40, 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputBody:         models.Pokemon{Name: "TestPokemon", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			originalPokemonId: 1,
			wantError:         false,
		},
		{
			testName: "Empty fields. Error",
			mockFunc: func() {
				row := mock.NewRows([]string{"id"}).
					AddRow(1)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnRows(row)

				mock.ExpectExec(regexp.QuoteMeta(`UPDATE pokemon
				SET name = $1, type = $2, hp = $3, attack = $4, defense = $5
				WHERE id = $6
				`)).WithArgs("TestPokemon", pq.Array([]string{"TestType"}), 40, 40, 40, 1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			inputBody:         models.Pokemon{Name: "", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			originalPokemonId: 1,
			wantError:         true,
		},
		{
			testName: "Not found. Error",
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			inputBody:         models.Pokemon{Name: "Test", PokemonType: []string{"TestType"}, Hp: 40, Attack: 40, Defense: 40},
			originalPokemonId: 1,
			wantError:         true,
		},
	}

	// Run tests
	for _, testCase := range tests {
		testCase.mockFunc()

		err := repo.UpdatePokemon(testCase.inputBody, testCase.originalPokemonId)

		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}

func Test_DeletePokemon(t *testing.T) {
	// Init mock DB
	mockDb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal("error not expected while starting mock database connection")
	}
	defer mockDb.Close()

	// Init repo interface with necessary method
	repo := &PokedexPostgres{db: mockDb}

	// Setup tests
	tests := []struct {
		testName          string
		mockFunc          func()
		originalPokemonId int
		wantError         bool
	}{
		{
			testName: "OK",
			mockFunc: func() {
				row := mock.NewRows([]string{"id"}).
					AddRow(1)

				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnRows(row)

				mock.ExpectExec(regexp.QuoteMeta(`DELETE FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))
			},
			originalPokemonId: 1,
			wantError:         false,
		},
		{
			testName: "Not found. Error",
			mockFunc: func() {
				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM pokemon WHERE id = $1`)).WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			originalPokemonId: 1,
			wantError:         true,
		},
	}

	// Run tests
	for _, testCase := range tests {
		testCase.mockFunc()

		err := repo.DeletePokemon(testCase.originalPokemonId)

		if testCase.wantError {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
		}
	}
}
