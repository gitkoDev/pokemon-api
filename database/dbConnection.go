package database

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/pressly/goose/v3"
)

const (
	user     = "postgres"
	password = 1234
	host     = "localhost"
	port     = 5432
	database = "pokemon_db"
)

var dsn = fmt.Sprintf("user= %v password= %v host=%v port=%v database=%v sslmode=disable", user, password, host, port, database)

var dbConnection *sql.DB

func ConnectToDb() error {
	db, err := sql.Open("pgx", dsn)

	if err != nil {
		return err
	}

	err = db.Ping()

	if err != nil {
		return err
	} else {
		log.Println("connected")
	}

	dbConnection = db
	migrateDb()

	return nil
}

func migrateDb() {
	err := goose.Up(dbConnection, "./database/migrations")
	if err != nil {
		fmt.Println(err)
	}
}
