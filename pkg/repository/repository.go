package repository

import "database/sql"

type Authorization interface{}

type AllPokemon interface{}

type Repository struct{}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{}
}
