package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gitkoDev/pokemon-db/helpers"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) IdentifyUser(w http.ResponseWriter, r *http.Request) {
	context := r.Context()
	header := context.Value(authorizationHeader)

	if header == "" {
		helpers.RespondWithError(w, errors.New("empty authorization header"), http.StatusUnauthorized)
		return
	}

	headerString := fmt.Sprintln(header)
	headerParts := strings.Split(headerString, " ")
	id, err := h.services.Authorization.ParseToken(headerParts[1])
	if err != nil {
		helpers.RespondWithError(w, err, http.StatusInternalServerError)
		return
	}

	log.Println(id)
}
