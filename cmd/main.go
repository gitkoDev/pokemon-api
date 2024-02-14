package main

import (
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/database"
	"github.com/gitkoDev/pokemon-db/pkg/handlers"

	"github.com/go-chi/chi/v5"
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
	r := chi.NewRouter()
	r.Get("/allPokemon", handlers.GetAll(db))
	r.Get("/allPokemon/{name}", handlers.GetByName(db))
	r.Post("/allPokemon", handlers.AddPokemon(db))
	r.Delete("/allPokemon/{name}", handlers.DeletePokemon(db))

	err = http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("error starting the server", err)
	}

}
