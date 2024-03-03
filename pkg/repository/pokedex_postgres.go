package repository

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/gitkoDev/pokemon-db/models"
	"github.com/lib/pq"
)

type PokedexPostgres struct {
	db *sql.DB
}

func NewPokedexPostgres(db *sql.DB) *PokedexPostgres {
	return &PokedexPostgres{db: db}
}

func (r *PokedexPostgres) AddPokemon(pokemonToAdd models.Pokemon) error {
	isExisting, err := checkForExistencePokedex(r, pokemonToAdd.Name)
	if err != nil {
		return err
	}
	if isExisting {
		errMsg := fmt.Sprintf("%s is already in Pokedex", pokemonToAdd.Name)
		return errors.New(errMsg)
	}

	query := `INSERT INTO pokemon (name, type, hp, attack, defense) VALUES($1, $2, $3, $4, $5)`
	_, err = r.db.Exec(query, pokemonToAdd.Name, pokemonToAdd.PokemonType, pokemonToAdd.Hp, pokemonToAdd.Attack, pokemonToAdd.Defense)
	if err != nil {
		return err
	}

	return nil
}

func (r *PokedexPostgres) GetAll() ([]models.Pokemon, error) {
	// Get pokemon from database
	query := `SELECT name, type, hp, attack, defense FROM pokemon`

	res, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}

	defer res.Close()

	selectedPokemon := []models.Pokemon{}

	for res.Next() {
		pokemon := models.Pokemon{}

		err := res.Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
		if err != nil {
			return nil, err
		}

		selectedPokemon = append(selectedPokemon, pokemon)
	}

	return selectedPokemon, nil
}

func (r *PokedexPostgres) GetByName(pokemonName string) (models.Pokemon, error) {
	pokemon := models.Pokemon{}

	query := `SELECT name, type, hp, attack, defense FROM pokemon WHERE name = $1`

	err := r.db.QueryRow(query, pokemonName).Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
	if err != nil {
		return models.Pokemon{}, err
	}

	// Print pokemon data if found
	return pokemon, nil
}

func (r *PokedexPostgres) UpdatePokemon(newPokemon models.Pokemon, originalName string) error {
	isExisting, err := checkForExistencePokedex(r, originalName)
	if err != nil {
		if err == sql.ErrNoRows {
			responseString := fmt.Sprintf("%s not found in Pokedex", originalName)
			return errors.New(responseString)
		}
		return err
	}
	if !isExisting {
		errMsg := fmt.Sprintf("%s is not in Pokedex", originalName)
		return errors.New(errMsg)
	}

	query := `UPDATE pokemon
		SET name = $1, type = $2, hp = $3, attack = $4, defense = $5
		WHERE name = $6
		`
	_, err = r.db.Exec(query, newPokemon.Name, newPokemon.PokemonType, newPokemon.Hp, newPokemon.Attack, newPokemon.Defense, originalName)
	if err != nil {
		return err
	}

	return nil
}

func (r *PokedexPostgres) DeletePokemon(pokemonName string) error {
	isExisting, err := checkForExistencePokedex(r, pokemonName)
	if err != nil {
		return err
	}
	if !isExisting {
		errMsg := fmt.Sprintf("%s is not in Pokedex", pokemonName)
		return errors.New(errMsg)
	}

	query := `DELETE FROM pokemon WHERE name = $1`

	_, err = r.db.Exec(query, pokemonName)
	if err != nil {
		return err
	}

	return nil
}

func checkForExistencePokedex(r *PokedexPostgres, pokemonName string) (bool, error) {
	query := `SELECT id FROM pokemon WHERE name = $1`
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
