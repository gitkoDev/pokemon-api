package controllers

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gitkoDev/pokemon-db/helpers"
	"github.com/gitkoDev/pokemon-db/models"
	"github.com/go-chi/chi/v5"

	"github.com/lib/pq"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello"))
}

func AddPokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Decode pokemon data from json

		pokemon, err := helpers.DecodeJson(r)
		if err != nil {
			log.Println("decodeJson() error:", err)
			return
		}

		isExisting := checkForExistence(db, pokemon.Name)
		if isExisting {
			responseString := fmt.Sprintf("%s already exists", pokemon.Name)
			w.Write([]byte(responseString))
			return
		}

		// Insert pokemon into database
		query := `INSERT INTO pokemon (name, type, hp, attack, defense) VALUES($1, $2, $3, $4, $5)`
		_, err = db.Exec(query, pokemon.Name, pokemon.PokemonType, pokemon.Hp, pokemon.Attack, pokemon.Defense)
		if err != nil {
			log.Println("InsertPokemon() error", err)
			return
		}

		responseString := fmt.Sprintf("%s added", pokemon.Name)
		w.Write([]byte(responseString))

	}
}

func GetAll(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get pokemon from database
		query := `SELECT name, type, hp, attack, defense FROM pokemon`

		res, err := db.Query(query)
		if err != nil {
			log.Println("GetAll() error:", err)
			return
		}

		defer res.Close()

		selectedPokemon := []models.Pokemon{}

		for res.Next() {
			pokemon := models.Pokemon{}

			err := res.Scan(&pokemon.Name, (*pq.StringArray)(&pokemon.PokemonType), &pokemon.Hp, &pokemon.Attack, &pokemon.Defense)
			if err != nil {
				log.Println("GetAll() error scanning row:", err)
				return
			}

			selectedPokemon = append(selectedPokemon, pokemon)
		}

		// Respond with json
		helpers.EncodeJson(w, selectedPokemon)
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
				w.Write([]byte(responseString))
				return
			}
			fmt.Println("GetByName() error:", err)
			return
		}

		// Print pokemon data if found
		err = helpers.EncodeJson(w, pokemon)
		if err != nil {
			log.Println("encodeJson() error:", err)
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
			w.Write([]byte(responseString))
			return
		}

		// Decode from request
		dataToInsert, err := helpers.DecodeJson(r)
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
		w.Write([]byte(responseString))

	}
}

func DeletePokemon(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		name := chi.URLParam(r, "name")

		isExisting := checkForExistence(db, name)
		if !isExisting {
			responseString := fmt.Sprintf("%s not found", name)
			w.Write([]byte(responseString))
			return
		}

		query := `DELETE FROM pokemon WHERE name = $1`

		_, err := db.Exec(query, name)

		if err != nil {
			if err == sql.ErrNoRows {
				fmt.Printf("%s not found\n", name)
				return
			}
			fmt.Println(err)
			return
		}

		responseString := fmt.Sprintf("%s deleted", name)
		w.Write([]byte(responseString))
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
