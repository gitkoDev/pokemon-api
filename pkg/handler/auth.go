package handler

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/helpers"
)

// "github.com/go-chi/chi/v5"

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {

	var input, err = helpers.DecodeTrainerJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Password == "" {
		responseString := "please provide valid trainer name and password"
		helpers.RespondWithError(w, responseString, http.StatusBadRequest)
		return
	}

	id, err := h.services.Authorization.CreateTrainer(input)
	if err != nil {
		helpers.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	responseString := fmt.Sprintf("Trainer: %s, ID: %v. Added", input.Name, id)
	helpers.WriteJSON(w, responseString, http.StatusOK)

}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	var input, _ = helpers.DecodeTrainerJSON(r)
	log.Println(input)

}
