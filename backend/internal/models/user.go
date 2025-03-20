package models

type User struct {
	ID     int    `json:"id" db:"id"`
	Name   string `json:"name" db:"name"`
	Hash   string `json:"hash" db:"hash"`
	Avatar string `json:"avatar" db:"avatar"`
}
