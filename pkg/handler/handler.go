package handler

import (
	"net/http"

	"github.com/gitkoDev/pokemon-db/pkg/service"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() http.Handler {
	r := chi.NewRouter()

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	r.Route("/pokemon-api", func(r chi.Router) {
		r.Get("/ping", h.ping)
		r.Route("/v1", func(r chi.Router) {
			r.Post("/allPokemon", h.addPokemon)
			r.Get("/allPokemon", h.getAll)
			r.Get("/allPokemon/{name}", h.getByName)
			r.Put("/allPokemon/{name}", h.updatePokemon)
			r.Delete("/allPokemon/{name}", h.deletePokemon)
		})
	})
	return r
}
