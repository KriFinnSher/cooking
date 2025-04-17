package models

type Schedule struct {
	ID        int    `json:"id" db:"id"`
	EventName string `json:"event_name" db:"event_name"`
	EventDate int    `json:"event_date" db:"event_date"`
	Location  string `json:"location" db:"location"`
	ChefID    int    `json:"chef_id" db:"chef_id"`
}
