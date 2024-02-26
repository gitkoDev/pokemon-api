package handler

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/helpers"
	"github.com/gitkoDev/pokemon-db/models"
	"github.com/gitkoDev/pokemon-db/pkg/service"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"

	"github.com/lib/pq"
)

type Handler struct {
	services *service.Service
}

func NewHanler(services *service.Service) *Handler {
	return &Handler{}
}

func Ping(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Pokemon API v.1.0"))

}

func AddPokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode pokemon data from json

		pokemon, err := helpers.DecodeJSON(r)
		if err != nil {
			logrus.Println("decodeJson() error:", err)
			return
		}

		isExisting := checkForExistence(db, pokemon.Name)
		if isExisting {
			responseString := fmt.Sprintf("%s already exists", pokemon.Name)
			helpers.WriteJSON(w, responseString, http.StatusBadRequest)
			return
		}

		// Insert pokemon into database
		query := `INSERT INTO pokemon (name, type, hp, attack, defense) VALUES($1, $2, $3, $4, $5)`
		_, err = db.Exec(query, pokemon.Name, pokemon.PokemonType, pokemon.Hp, pokemon.Attack, pokemon.Defense)
		if err != nil {
			logrus.Println("InsertPokemon() error", err)
			return
		}

		responseString := fmt.Sprintf("%s added", pokemon.Name)
		helpers.WriteJSON(w, responseString, http.StatusCreated)

	}
}

func GetAll(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get pokemon from database
		query := `SELECT name, type, hp, attack, defense FROM pokemon`

		res, err := db.Query(query)
		if err != nil {
			logrus.Println("GetAll() error:", err)
			return
		}

		defer res.Close()

		selectedPokemon := []models.Pokemon{}

		for res.Next() {
			pokemon := models.Pokemon{}

			err := res.Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
			if err != nil {
				logrus.Println("GetAll() error scanning row:", err)
				return
			}

			selectedPokemon = append(selectedPokemon, pokemon)
		}

		// Respond with json
		helpers.WriteJSON(w, selectedPokemon, http.StatusOK)
	}
}

func GetByName(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		pokemon := models.Pokemon{}

		name := chi.URLParam(r, "name")

		query := `SELECT name, type, hp, attack, defense FROM pokemon WHERE name = $1`

		err := db.QueryRow(query, name).Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
		if err != nil {
			if err == sql.ErrNoRows {
				responseString := fmt.Sprintf("%s not found", name)
				helpers.WriteJSON(w, responseString, http.StatusBadRequest)
				return
			}
			fmt.Println("GetByName() error:", err)
			return
		}

		// Print pokemon data if found
		err = helpers.WriteJSON(w, pokemon, http.StatusOK)
		if err != nil {
			logrus.Println("encodeJson() error:", err)
			return
		}

	}
}

func UpdatePokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		name := chi.URLParam(r, "name")

		// Check for existence
		isExisting := checkForExistence(db, name)
		if !isExisting {
			responseString := fmt.Sprintf("%s not found", name)
			helpers.WriteJSON(w, responseString, http.StatusBadRequest)
			return
		}

		// Decode from request
		dataToInsert, err := helpers.DecodeJSON(r)
		if err != nil {
			fmt.Println(err)
			return
		}

		// Update value in DB
		query := `UPDATE pokemon
		SET name = $1, type = $2, hp = $3, attack = $4, defense = $5
		WHERE name = $6
		`
		_, err = db.Exec(query, dataToInsert.Name, dataToInsert.PokemonType, dataToInsert.Hp, dataToInsert.Attack, dataToInsert.Defense, name)
		if err != nil {
			log.Println("UpdatePokemon() error:", err)
			return
		}

		responseString := fmt.Sprintf("%s updated", name)
		helpers.WriteJSON(w, responseString, http.StatusOK)

	}
}

func DeletePokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		name := chi.URLParam(r, "name")

		isExisting := checkForExistence(db, name)
		if !isExisting {
			responseString := fmt.Sprintf("%s not found", name)
			helpers.WriteJSON(w, responseString, http.StatusBadRequest)
			return
		}

		query := `DELETE FROM pokemon WHERE name = $1`

		_, err := db.Exec(query, name)

		if err != nil {
			if err == sql.ErrNoRows {
				log.Printf("%s not found\n", name)
				return
			}
			log.Println(err)
			return
		}

		responseString := fmt.Sprintf("%s deleted", name)
		helpers.WriteJSON(w, responseString, http.StatusOK)
	}
}

func checkForExistence(db *sql.DB, name string) bool {
	query := `SELECT id FROM pokemon WHERE name = $1`
	var id uint

	err := db.QueryRow(query, name).Scan(&id)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		fmt.Println("checkForExistence() error:", err)
	}

	return true
}
