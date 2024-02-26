package main

import (
	"log"

	database "github.com/gitkoDev/pokemon-db/db"
	"github.com/gitkoDev/pokemon-db/server"
)

func main() {
	// DB connection phase
	db, err := database.ConnectToDb()
	if err != nil {
		log.Fatalln("error connecting to database:", err)
	} else {
		log.Println("connected to database")
	}

	// Routing phase
	srv := new(server.Server)
	err = srv.Run("8080", server.Router(db))
	if err != nil {
		log.Fatalln("error running server", err)
	}

}
