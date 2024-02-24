package server

import (
	"context"
	"database/sql"
	"log"
	"net/http"
	"time"

	"github.com/gitkoDev/pokemon-db/controllers"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func Route(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	r.Get("/hello", controllers.Hello)
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
