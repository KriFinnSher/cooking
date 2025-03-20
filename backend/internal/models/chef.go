package models

type Chef struct {
	ID             int    `json:"id" db:"id"`
	Name           string `json:"name" db:"name"`
	Hash           string `json:"hash" db:"hash"`
	FollowersCount int    `json:"followers_count" db:"followers_count"`
	Bio            string `json:"bio" db:"bio"`
	Avatar         string `json:"avatar" db:"avatar"`
}
