package main

import (
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/database"
	"github.com/gitkoDev/pokemon-db/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func main() {
	if err := database.ConnectToDb(); err != nil {
		log.Fatalln("error connecting to the database:", err)
	}

	r := chi.NewRouter()
	r.Get("/", handlers.TestFunc)

	err := http.ListenAndServe(":8080", r)

	if err != nil {
		log.Fatalln("error starting the server", err)
	}

}
