package service

import "github.com/gitkoDev/pokemon-db/pkg/repository"

type Authorization interface {
}

type AllPokemon interface {
}

type Service struct {
	Authorization
	AllPokemon
}

func NewService(repo *repository.Repository) *Service {
	return &Service{}
}
