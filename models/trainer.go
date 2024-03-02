package models

type Trainer struct {
	Id       int    `json:"-"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
