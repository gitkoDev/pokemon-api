package handler

import (
	"fmt"
	"net/http"

	"github.com/gitkoDev/pokemon-db/helpers"
	"github.com/go-chi/chi/v5"
)

func (h *Handler) ping(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pokemon API v.1.0"))
}

func (h *Handler) addPokemon(w http.ResponseWriter, r *http.Request) {
	// Decode pokemon data from json
	pokemon, err := helpers.DecodePokemonJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Run DB query
	err = h.services.Pokedex.AddPokemon(pokemon)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond if successful
	helpers.WriteJSON(w, pokemon, http.StatusCreated)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	// Run DB query
	pokemon, err := h.services.Pokedex.GetAll()
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond if succesful
	helpers.WriteJSON(w, pokemon, http.StatusOK)
}

func (h *Handler) getByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	// Run DB query
	pokemon, err := h.services.Pokedex.GetByName(name)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Respond if succesful
	helpers.WriteJSON(w, pokemon, http.StatusOK)
}

func (h *Handler) updatePokemon(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	// Decode from request
	pokemon, err := helpers.DecodePokemonJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}
	// Update value in DB
	err = h.services.Pokedex.UpdatePokemon(pokemon, name)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseString := fmt.Sprintf("%s updated", name)
	helpers.WriteJSON(w, responseString, http.StatusOK)
}

func (h *Handler) deletePokemon(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")

	err := h.services.Pokedex.DeletePokemon(name)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseString := fmt.Sprintf("%s deleted", name)
	helpers.WriteJSON(w, responseString, http.StatusOK)
}
