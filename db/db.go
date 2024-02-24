package db

import (
	"database/sql"
	"fmt"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	user     = "pokemon_db"
	password = 1234
	host     = "localhost"
	port     = 5432
	database = "pokemon_db"
)

var dsn = fmt.Sprintf("user= %v password= %v host=%v port=%v database=%v sslmode=disable", user, password, host, port, database)

var DB *sql.DB

func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	DB = db

	return db, nil
}
