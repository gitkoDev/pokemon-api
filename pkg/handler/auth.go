package handler

import (
	"errors"
	"net/http"

	"github.com/gitkoDev/pokemon-api/helpers"
)

func (h *Handler) signUp(w http.ResponseWriter, r *http.Request) {

	var input, err = helpers.DecodeTrainerJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	if input.Name == "" || input.Password == "" {
		responseString := "please provide valid trainer name and password"
		helpers.RespondWithError(w, errors.New(responseString), http.StatusBadRequest)
		return
	}

	id, err := h.services.Authorization.CreateTrainer(input)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	helpers.WriteJSON(w, map[string]int{"id": id}, http.StatusCreated)
}

func (h *Handler) signIn(w http.ResponseWriter, r *http.Request) {
	// Decode user data from request
	input, err := helpers.DecodeTrainerJSON(r)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// Validate user data
	if input.Name == "" || input.Password == "" {
		responseString := "please provide valid trainer name and password"
		helpers.RespondWithError(w, errors.New(responseString), http.StatusBadRequest)
		return
	}

	// Check for user's existence in DB
	_, err = h.services.Authorization.GetTrainer(input.Name, input.Password)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusBadRequest)
		return
	}

	// If user found, generate token for him
	token, err := h.services.Authorization.GenerateToken(input.Name, input.Password)
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	helpers.WriteJSON(w, map[string]string{
		"token": token,
	}, http.StatusOK)
}
