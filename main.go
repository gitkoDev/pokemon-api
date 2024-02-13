package main

import (
	"fmt"
	"log"

	"github.com/gitkoDev/pokemon-db/database"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
)

func main() {
	db, err := database.ConnectToDb()
	if err != nil {
		log.Fatalln("error connecting to the database:", err)
	}

	err = goose.Up(db, "./database/sql/schema")
	if err != nil {
		fmt.Println(err)
	}
}
