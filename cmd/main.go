package main

import (
	"log"

	database "github.com/gitkoDev/pokemon-db/db"
	"github.com/gitkoDev/pokemon-db/router"
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
	router.Route(db)

}
