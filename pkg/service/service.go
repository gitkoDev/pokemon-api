package service

import (
	"github.com/gitkoDev/pokemon-api/models"
	"github.com/gitkoDev/pokemon-api/pkg/repository"
)

type Authorization interface {
	CreateTrainer(trainer models.Trainer) (int, error)
	GetTrainer(name string, password string) (models.Trainer, error)
	GenerateToken(name string, password string) (string, error)
	ParseToken(token string) (int, error)
}

type Pokedex interface {
	GetAll() ([]models.Pokemon, error)
	GetById(pokemonId int) (models.Pokemon, error)
	AddPokemon(pokemon models.Pokemon) (int, error)
	UpdatePokemon(pokemon models.Pokemon, pokemonId int) error
	DeletePokemon(pokemonId int) error
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
