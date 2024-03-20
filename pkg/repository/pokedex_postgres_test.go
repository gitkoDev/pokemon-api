package repository

// import (
// 	"regexp"
// 	"testing"

// 	"github.com/DATA-DOG/go-sqlmock"
// 	"github.com/gitkoDev/pokemon-api/models"
// )

// // func TestAddPokemon(t *testing.T) {
// // 	tests := []struct {
// // 		testName   string
// // 		inputBody  models.Pokemon
// // 		beforeTest func(sqlmock.Sqlmock)
// // 		want       models.Pokemon
// // 		wantErr    bool
// // 	}{
// // 		{
// // 			testName:  "should not return error",
// // 			inputBody: models.Pokemon{Name: "Pikachu", PokemonType: []string{"Electric"}, Hp: 40, Attack: 40, Defense: 40},
// // 			beforeTest: func(s sqlmock.Sqlmock) {
// // 				s.ExpectQuery(`INSERT INTO pokemon (name, type, hp, attack, defense) VALUES($1, $2, $3, $4, $5)`).
// // 					WithArgs("Pikachu", []string{"Electric"}, 40, 40, 40).
// // 					WillReturnError(errors.New("AddPokemon returns error"))
// // 			},
// // 			wantErr: true,
// // 		},
// // 	}

// // 	for _, testCase := range tests {
// // 		t.Run(testCase.testName, func(t *testing.T) {
// // 			mockDB, mock, err := sqlmock.New()
// // 			if err != nil {
// // 				t.Fatalf("error while opening mock database connection")
// // 			}
// // 			defer mockDB.Close()

// // 			// db := sqlx.NewDb(mockDB, "sqlmock")

// // 			repo := &PokedexPostgres{db: mockDB}

// // 			if testCase.beforeTest != nil {
// // 				testCase.beforeTest(mock)
// // 			}

// // 			err = repo.AddPokemon(testCase.inputBody)
// // 			if err != nil {
// // 				t.Fatal(err)
// // 			}
// // 			// if (err != nil) != testCase.wantErr {
// // 			// 	t.Errorf("AddPokemon error: %v, wantError: %v", err, testCase.wantErr)
// // 			// 	return
// // 			// }
// // 		})
// // 	}
// // }

// func TestAddPokemon(t *testing.T) {
// 	mockDB, mock, err := sqlmock.New()
// 	if err != nil {
// 		t.Fatal("error not expected while starting mock database connection")
// 	}
// 	defer mockDB.Close()

// 	tests := []struct {
// 		testName  string
// 		mock      func()
// 		inputBody models.Pokemon
// 		want      int
// 		wantError bool
// 	}{
// 		{
// 			testName:  "should not give error",
// 			inputBody: models.Pokemon{Name: "Test", PokemonType: []string{"TestPower"}, Hp: 40, Attack: 40, Defense: 40},
// 			mock: func() {
// 				mock.ExpectQuery(regexp.QuoteMeta(`SELECT id FROM pokemon WHERE name = $1`)).WithArgs("Test").WillReturnRows(mock.NewRows([]string{"id"}).AddRow(0))

// 			},
// 			wantError: false,
// 		},
// 	}

// 	for _, testCase := range tests {
// 		testCase.mock()

// 		repo := &PokedexPostgres{db: mockDB}

// 		err = repo.AddPokemon(testCase.inputBody)
// 		if err != nil {
// 			t.Fatalf("error not expected while testing AddPokemon got: %v", err)
// 		}

// 	}

// }
