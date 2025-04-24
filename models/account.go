package models

type Account struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Pin     string `json:"pin"`
	Balance int    `json:"balance"`
}