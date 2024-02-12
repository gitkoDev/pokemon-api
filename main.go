package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

const (
	user     = "postgres"
	password = 1234
	host     = "localhost"
	port     = 5432
	database = "pokemon_db"
)

var dsn = fmt.Sprintf("user= %v password= %v host=%v port=%v database=%v sslmode=disable", user, password, host, port, database)

func main() {
	if err := connectToDb(); err != nil {
		log.Fatalln("error connecting to the database:", err)
	}

}

func connectToDb() error {
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

	return nil
}
