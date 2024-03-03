package repository

import (
	"database/sql"

	"github.com/gitkoDev/pokemon-db/models"
)

type Authorization interface {
	CreateTrainer(trainer models.Trainer) (int, error)
	GetTrainer(name string, password string) (models.Trainer, error)
}

type Pokedex interface {
	GetAll() ([]models.Pokemon, error)
	GetByName(pokemonName string) (models.Pokemon, error)
	AddPokemon(pokemon models.Pokemon) error
	UpdatePokemon(pokemon models.Pokemon, pokemonName string) error
	DeletePokemon(pokemonName string) error
}

type Repository struct {
	Authorization
	Pokedex
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Pokedex:       NewPokedexPostgres(db),
	}
}
