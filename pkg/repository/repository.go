package repository

import (
	"database/sql"

	"github.com/gitkoDev/pokemon-api/models"
)

type Authorization interface {
	CreateTrainer(trainer models.Trainer) (int, error)
	GetTrainer(name string, password string) (models.Trainer, error)
}

type Pokedex interface {
	GetAll() ([]models.Pokemon, error)
	GetById(pokemonId int) (models.Pokemon, error)
	AddPokemon(pokemon models.Pokemon) (int, error)
	UpdatePokemon(pokemon models.Pokemon, pokemonId int) error
	DeletePokemon(pokemonId int) error
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
