package repository

import (
	"database/sql"
	"errors"
	"fmt"

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

func (r *AuthPostgres) GetTrainer(name string, password string) (models.Trainer, error) {
	// Check if trainer exists, return ID if found, return if not existing
	id, err := checkForExistenceTrainers(r, name, password)
	if id == -1 {
		responseString := fmt.Sprintln("trainer with this name and password not found")
		return models.Trainer{}, errors.New(responseString)
	}
	if err != nil {
		return models.Trainer{}, err
	}

	trainer := models.Trainer{Id: id, Name: name, Password: password}
	return trainer, nil
}

func checkForExistenceTrainers(r *AuthPostgres, name string, password string) (int, error) {
	query := "SELECT id FROM pokemon_trainers WHERE name = $1 AND password_hash = $2"
	var id int

	err := r.db.QueryRow(query, name, password).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return -1, err
		}
		return 0, err
	}

	return int(id), nil
}
