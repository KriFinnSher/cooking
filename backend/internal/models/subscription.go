package models

type Subscription struct {
	UserID int `json:"user_id" db:"user_id"`
	ChefID int `json:"chef_id" db:"chef_id"`
}
