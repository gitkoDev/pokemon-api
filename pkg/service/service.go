package service

import (
	"github.com/gitkoDev/pokemon-db/models"
	"github.com/gitkoDev/pokemon-db/pkg/repository"
)

type Authorization interface {
	CreateTrainer(trainer models.Trainer) (int, error)
	GetTrainer(name string, password string) (models.Trainer, error)
	GenerateToken(name string, password string) (string, error)
}

type Pokedex interface {
	GetAll() ([]models.Pokemon, error)
	GetByName(pokemonName string) (models.Pokemon, error)
	AddPokemon(pokemon models.Pokemon) error
	UpdatePokemon(pokemon models.Pokemon, pokemonName string) error
	DeletePokemon(pokemonName string) error
}

type Service struct {
	Authorization
	Pokedex
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repo.Authorization),
		Pokedex:       NewPokemonListService(repo.Pokedex),
	}
}
