package server

import (
	"context"
	"database/sql"
	"net/http"
	"time"

	"github.com/gitkoDev/pokemon-db/pkg/handler"
	"github.com/go-chi/chi/v5"
)

type Server struct {
	httpServer *http.Server
}

func (s *Server) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      handler,
	}

	return s.httpServer.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}

func Router(db *sql.DB) http.Handler {
	r := chi.NewRouter()

	r.Route("/pokemon-api", func(r chi.Router) {
		r.Get("/ping", handler.Ping)
		r.Route("/v1", func(r chi.Router) {
			r.Post("/allPokemon", handler.AddPokemon(db))
			r.Get("/allPokemon", handler.GetAll(db))
			r.Get("/allPokemon/{name}", handler.GetByName(db))
			r.Put("/allPokemon/{name}", handler.UpdatePokemon(db))
			r.Delete("/allPokemon/{name}", handler.DeletePokemon(db))
		})
	})
	return r
}
