package router

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/controllers"
	"github.com/go-chi/chi/v5"
)

func Route(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Post("/allPokemon", controllers.AddPokemon(db))
	r.Get("/allPokemon", controllers.GetAll(db))
	r.Get("/allPokemon/{name}", controllers.GetByName(db))
	r.Put("/allPokemon/{name}", controllers.UpdatePokemon(db))
	r.Delete("/allPokemon/{name}", controllers.DeletePokemon(db))

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Fatalln("error starting the server", err)
	}

	return r
}