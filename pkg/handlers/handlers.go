package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
)

func GetAll(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("HERE", db)
	}
}
