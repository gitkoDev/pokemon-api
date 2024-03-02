package service

import (
	"github.com/gitkoDev/pokemon-db/models"
	"github.com/gitkoDev/pokemon-db/pkg/repository"
)

type PokemonService struct {
	repo repository.Pokedex
}

func NewPokemonListService(repo repository.Pokedex) *PokemonService {
	return &PokemonService{repo: repo}
}

func (s *PokemonService) AddPokemon(pokemonToAdd models.Pokemon) error {
	return s.repo.AddPokemon(pokemonToAdd)
}

func (s *PokemonService) GetAll() ([]models.Pokemon, error) {
	return s.repo.GetAll()
}

func (s *PokemonService) GetByName(pokemonName string) (models.Pokemon, error) {
	return s.repo.GetByName(pokemonName)
}

func (s *PokemonService) UpdatePokemon(pokemonToUpdate models.Pokemon, pokemonName string) error {
	return s.repo.UpdatePokemon(pokemonToUpdate, pokemonName)
}

func (s *PokemonService) DeletePokemon(pokemonToDelete string) error {
	return s.repo.DeletePokemon(pokemonToDelete)
}
