package models

type Trainer struct {
	Id       int    `json:"-" db:"id"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
