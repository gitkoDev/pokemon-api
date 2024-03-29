package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gitkoDev/pokemon-api/helpers"
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
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// Run DB query
	id, err := h.services.Pokedex.AddPokemon(pokemon)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// By default "id" is a zero value, but we need to show the user what the sql "id" of updated pokemon is
	pokemon.Id = id
	helpers.WriteJSON(w, pokemon, http.StatusCreated)
}

func (h *Handler) getAll(w http.ResponseWriter, r *http.Request) {
	// Run DB query
	pokemon, err := h.services.Pokedex.GetAll()
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Respond if succesful
	helpers.WriteJSON(w, pokemon, http.StatusOK)
}

func (h *Handler) getById(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// Run DB query
	pokemon, err := h.services.Pokedex.GetById(intId)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// Respond if succesful
	helpers.WriteJSON(w, pokemon, http.StatusOK)
}

func (h *Handler) updatePokemon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// Decode from request
	pokemon, err := helpers.DecodePokemonJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}
	// Update value in DB
	err = h.services.Pokedex.UpdatePokemon(pokemon, intId)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	// By default "id" is a zero value, but we need to show the user what the sql "id" of updated pokemon is
	pokemon.Id = intId
	helpers.WriteJSON(w, pokemon, http.StatusOK)
}

func (h *Handler) deletePokemon(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	err = h.services.Pokedex.DeletePokemon(intId)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	responseString := fmt.Sprintf("%d deleted", intId)
	helpers.RespondWithMessage(w, responseString, 200)
}
