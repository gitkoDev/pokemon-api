package main

import (
	"log"

	"github.com/gitkoDev/pokemon-db/database"
	_ "github.com/jackc/pgx/v5/stdlib"
)

func main() {
	if err := database.ConnectToDb(); err != nil {
		log.Fatalln("error connecting to the database:", err)
	}

}
