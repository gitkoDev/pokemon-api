package repository

import (
	"database/sql"

	"github.com/gitkoDev/pokemon-api/models"
	"github.com/lib/pq"
)

type PokedexPostgres struct {
	db *sql.DB
}

func NewPokedexPostgres(db *sql.DB) *PokedexPostgres {
	return &PokedexPostgres{db: db}
}

func (r *PokedexPostgres) AddPokemon(pokemonToAdd models.Pokemon) (int, error) {
	var pokemonId int

	query := `INSERT INTO pokemon (name, type, hp, attack, defense) VALUES($1, $2, $3, $4, $5) RETURNING id`
	row := r.db.QueryRow(query, pokemonToAdd.Name, pokemonToAdd.PokemonType, pokemonToAdd.Hp, pokemonToAdd.Attack, pokemonToAdd.Defense)

	err := row.Scan(&pokemonId)
	if err != nil {
		return 0, err
	}

	return pokemonId, nil
}

func (r *PokedexPostgres) GetAll() ([]models.Pokemon, error) {
	// Get pokemon from database
	query := `SELECT id, name, type, hp, attack, defense FROM pokemon`

	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	selectedPokemon := []models.Pokemon{}

	for res.Next() {
		pokemon := models.Pokemon{}

		err := res.Scan(&pokemon.Id, &pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
		if err != nil {
			return nil, err
		}

		selectedPokemon = append(selectedPokemon, pokemon)
	}

	return selectedPokemon, nil
}

func (r *PokedexPostgres) GetById(pokemonId int) (models.Pokemon, error) {
	pokemon := models.Pokemon{}

	query := `SELECT id, name, type, hp, attack, defense FROM pokemon WHERE id = $1`

	err := r.db.QueryRow(query, pokemonId).Scan(&pokemon.Id, &pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
	if err != nil {
		return models.Pokemon{}, err
	}

	// Print pokemon data if found
	return pokemon, nil
}

func (r *PokedexPostgres) UpdatePokemon(newPokemon models.Pokemon, pokemonId int) error {
	// Check for pokemon existence
	var scannedId int
	query := `SELECT id FROM pokemon WHERE id = $1`
	err := r.db.QueryRow(query, pokemonId).Scan(&scannedId)
	if err != nil {
		return err
	}

	// Update pokemon if exists
	query = `UPDATE pokemon
		SET name = $1, type = $2, hp = $3, attack = $4, defense = $5
		WHERE id = $6
		`
	_, err = r.db.Exec(query, newPokemon.Name, newPokemon.PokemonType, newPokemon.Hp, newPokemon.Attack, newPokemon.Defense, pokemonId)
	if err != nil {
		return err
	}

	return nil
}

func (r *PokedexPostgres) DeletePokemon(pokemonId int) error {
	// Check for pokemon existence
	var scannedId int
	query := `SELECT id FROM pokemon WHERE id = $1`
	err := r.db.QueryRow(query, pokemonId).Scan(&scannedId)
	if err != nil {
		return err
	}

	// Delete pokemon if exists
	query = `DELETE FROM pokemon WHERE id = $1`

	_, err = r.db.Exec(query, pokemonId)
	if err != nil {
		return err
	}

	return nil
}

func checkForExistencePokedex(r *PokedexPostgres, pokemonName string) (bool, error) {
	query := `SELECT id FROM pokemon WHERE id = $1`
	var id uint

	err := r.db.QueryRow(query, pokemonName).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}

	return true, nil
}
