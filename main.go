package main

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := connectToDb(); err != nil {
		log.Fatalln("error connecting to the database:", err)
	}

}

func connectToDb() error {
	db, err := sql.Open("pgx", "postgres://postgres:1234@localhost:5432/pokemon_db?sslmode=disable")

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
