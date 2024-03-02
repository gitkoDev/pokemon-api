package repository

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	pokemonTable  = "pokemon"
	trainersTable = "pokemon_trainers"
)

type Config struct {
	Host     string
	Port     string
	UserName string
	Password string
	DBName   string
	SSLMode  string
}

func NewDB(cfg Config) (*sql.DB, error) {
	db, err := sql.Open("pgx", fmt.Sprintf("user= %v password= %v host=%v port=%v database=%v sslmode=disable", cfg.UserName, cfg.Password, cfg.Host, cfg.Port, cfg.DBName))
	if err != nil {
		errMessage := fmt.Sprintf("error opening DB connection: %v", err)
		return nil, errors.New(errMessage)
	}

	if err = db.Ping(); err != nil {
		errMessage := fmt.Sprintf("Ping() DB error: %v", err)
		return nil, errors.New(errMessage)
	}

	return db, nil
}
