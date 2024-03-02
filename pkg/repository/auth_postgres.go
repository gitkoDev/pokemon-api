package repository

import (
	"database/sql"

	"github.com/gitkoDev/pokemon-db/models"
)

type AuthPostgres struct {
	db *sql.DB
}

func NewAuthPostgres(db *sql.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) CreateTrainer(trainer models.Trainer) (int, error) {
	var id int

	query := `INSERT INTO pokemon_trainers (NAME, PASSWORD_HASH) VALUES ($1, $2) RETURNING id`
	if err := r.db.QueryRow(query, trainer.Name, trainer.Password).Scan(&id); err != nil {
		return 0, err
	}

	return id, nil
}
