package recipe

import (
	"context"
	"cooking/backend/internal/models"
)

type Repo interface {
	CreateRecipe(ctx context.Context, userID int, title string, ingredients map[string]int, recipeText string) error
	GetRecipe(ctx context.Context, recipeID int) (models.Recipe, error)
	GetUserRecipes(ctx context.Context, userID int) ([]models.Recipe, error)
	UpdateRecipe(ctx context.Context, recipeID int, title string, ingredients map[string]int, recipeText string) error
	DeleteRecipe(ctx context.Context, recipeID int) error
}
