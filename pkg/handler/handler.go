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

	r.Get("/health", h.ping)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	r.With(h.IdentifyUser).Route("/api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Post("/pokemon", h.addPokemon)
			r.Get("/pokemon", h.getAll)
			r.Get("/pokemon/{name}", h.getByName)
			r.Put("/pokemon/{name}", h.updatePokemon)
			r.Delete("/pokemon/{name}", h.deletePokemon)
		})
	})
	return r
}
