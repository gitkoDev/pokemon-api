package handler

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gitkoDev/pokemon-api/helpers"
)

const (
	authorizationHeader = "Authorization"
)

func (h *Handler) IdentifyUser(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authToken := r.Header.Values(authorizationHeader)

		if authToken == nil {
			helpers.RespondWithError(w, errors.New("empty authorization header"), http.StatusUnauthorized)
			return
		}

		authToken = strings.Split(authToken[0], " ")
		tokenPart := authToken[1]
		id, err := h.services.Authorization.ParseToken(tokenPart)
		if err != nil {
			helpers.RespondWithError(w, err, http.StatusInternalServerError)
			return
		}

		log.Println("success", id)

		next.ServeHTTP(w, r)
	})

}
