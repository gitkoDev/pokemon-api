package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/pkg/handlers"
	"github.com/go-chi/chi/v5"
)

func Route(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/allPokemon", handlers.GetAll(db))
	r.Get("/allPokemon/{name}", handlers.GetByName(db))
	r.Post("/allPokemon", handlers.AddPokemon(db))
	r.Delete("/allPokemon/{name}", handlers.DeletePokemon(db))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("error starting the server", err)
	}

	return r
}
