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

	r.Get("/ping", h.ping)

	r.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	r.With(h.IdentifyUser).Route("/pokemon-api", func(r chi.Router) {

		r.Route("/v1", func(r chi.Router) {

			r.Post("/Pokemon", h.addPokemon)
			r.Get("/Pokemon", h.getAll)
			r.Get("/Pokemon/{name}", h.getByName)
			r.Put("/Pokemon/{name}", h.updatePokemon)
			r.Delete("/Pokemon/{name}", h.deletePokemon)
		})
	})
	return r
}
