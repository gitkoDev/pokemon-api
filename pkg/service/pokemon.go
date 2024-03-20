package service

import (
	"github.com/gitkoDev/pokemon-api/models"
	"github.com/gitkoDev/pokemon-api/pkg/repository"
)

type PokemonService struct {
	repo repository.Pokedex
}

func NewPokemonListService(repo repository.Pokedex) *PokemonService {
	return &PokemonService{repo: repo}
}

func (s *PokemonService) AddPokemon(pokemonToAdd models.Pokemon) (int, error) {
	return s.repo.AddPokemon(pokemonToAdd)
}

func (s *PokemonService) GetAll() ([]models.Pokemon, error) {
	return s.repo.GetAll()
}

func (s *PokemonService) GetById(pokemonId int) (models.Pokemon, error) {
	return s.repo.GetById(pokemonId)
}

func (s *PokemonService) UpdatePokemon(pokemonToUpdate models.Pokemon, pokemonId int) error {
	return s.repo.UpdatePokemon(pokemonToUpdate, pokemonId)
}

func (s *PokemonService) DeletePokemon(pokemonId int) error {
	return s.repo.DeletePokemon(pokemonId)
}
