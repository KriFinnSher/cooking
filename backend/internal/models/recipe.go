package models

type Recipe struct {
	ID          int            `json:"id" db:"id"`
	UserID      int            `json:"user_id" db:"user_id"`
	Title       string         `json:"title" db:"title"`
	Ingredients map[string]int `json:"ingredients" db:"ingredients"`
	RecipeText  string         `json:"recipe_text" db:"recipe_text"`
}
