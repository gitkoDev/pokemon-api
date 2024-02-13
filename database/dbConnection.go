package database

import (
	"database/sql"
	"fmt"
	"log"
)

const (
	user     = "postgres"
	password = 1234
	host     = "localhost"
	port     = 5432
	database = "pokemon_db"
)

var dsn = fmt.Sprintf("user= %v password= %v host=%v port=%v database=%v sslmode=disable", user, password, host, port, database)

func ConnectToDb() (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	} else {
		log.Println("connected")
	}

	return db, nil
}
