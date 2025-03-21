package postgres

import (
	"context"
	"cooking/backend/internal/models"
	"encoding/json"
	"errors"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type RecipeRepo struct {
	db *sqlx.DB
}

func NewRecipeRepo(db *sqlx.DB) *RecipeRepo {
	return &RecipeRepo{db: db}
}

func (r *RecipeRepo) CreateRecipe(ctx context.Context, userID int, title string, ingredients map[string]int, recipeText string) error {
	jsonIngredients, err := json.Marshal(ingredients)
	if err != nil {
		return errors.New("failed to serialize ingredients")
	}

	query, args, err := squirrel.Insert("recipes").
		Columns("user_id", "title", "ingredients", "recipe_text").
		Values(userID, title, jsonIngredients, recipeText).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *RecipeRepo) GetRecipe(ctx context.Context, recipeID int) (models.Recipe, error) {
	query, args, err := squirrel.Select("id", "user_id", "title", "ingredients", "recipe_text").
		From("recipes").
		Where(squirrel.Eq{"id": recipeID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Recipe{}, err
	}

	var recipe struct {
		ID          int    `db:"id"`
		UserID      int    `db:"user_id"`
		Title       string `db:"title"`
		Ingredients []byte `db:"ingredients"`
		RecipeText  string `db:"recipe_text"`
	}

	err = r.db.GetContext(ctx, &recipe, query, args...)
	if err != nil {
		return models.Recipe{}, err
	}

	var ingredients map[string]int
	if err := json.Unmarshal(recipe.Ingredients, &ingredients); err != nil {
		return models.Recipe{}, errors.New("failed to parse ingredients JSON")
	}

	return models.Recipe{
		ID:          recipe.ID,
		UserID:      recipe.UserID,
		Title:       recipe.Title,
		Ingredients: ingredients,
		RecipeText:  recipe.RecipeText,
	}, nil
}

func (r *RecipeRepo) GetUserRecipes(ctx context.Context, userID int) ([]models.Recipe, error) {
	query, args, err := squirrel.Select("id", "user_id", "title", "ingredients", "recipe_text").
		From("recipes").
		Where(squirrel.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var rawRecipes []struct {
		ID          int    `db:"id"`
		UserID      int    `db:"user_id"`
		Title       string `db:"title"`
		Ingredients []byte `db:"ingredients"`
		RecipeText  string `db:"recipe_text"`
	}

	err = r.db.SelectContext(ctx, &rawRecipes, query, args...)
	if err != nil {
		return nil, err
	}

	recipes := make([]models.Recipe, 0, len(rawRecipes))
	for _, raw := range rawRecipes {
		var ingredients map[string]int
		if err := json.Unmarshal(raw.Ingredients, &ingredients); err != nil {
			return nil, errors.New("failed to parse ingredients JSON")
		}

		recipes = append(recipes, models.Recipe{
			ID:          raw.ID,
			UserID:      raw.UserID,
			Title:       raw.Title,
			Ingredients: ingredients,
			RecipeText:  raw.RecipeText,
		})
	}

	return recipes, nil
}

func (r *RecipeRepo) GetAll(ctx context.Context) ([]models.Recipe, error) {
	query, args, err := squirrel.Select("id", "user_id", "title", "ingredients", "recipe_text").
		From("recipes").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var rawRecipes []struct {
		ID          int    `db:"id"`
		UserID      int    `db:"user_id"`
		Title       string `db:"title"`
		Ingredients []byte `db:"ingredients"`
		RecipeText  string `db:"recipe_text"`
	}

	err = r.db.SelectContext(ctx, &rawRecipes, query, args...)
	if err != nil {
		return nil, err
	}

	recipes := make([]models.Recipe, 0, len(rawRecipes))
	for _, raw := range rawRecipes {
		var ingredients map[string]int
		if err := json.Unmarshal(raw.Ingredients, &ingredients); err != nil {
			return nil, errors.New("failed to parse ingredients JSON")
		}

		recipes = append(recipes, models.Recipe{
			ID:          raw.ID,
			UserID:      raw.UserID,
			Title:       raw.Title,
			Ingredients: ingredients,
			RecipeText:  raw.RecipeText,
		})
	}

	return recipes, nil
}

func (r *RecipeRepo) UpdateRecipe(ctx context.Context, recipeID int, title string, ingredients map[string]int, recipeText string) error {
	ingredientsJSON, err := json.Marshal(ingredients)
	if err != nil {
		return errors.New("failed to encode ingredients JSON")
	}

	query, args, err := squirrel.Update("recipes").
		Set("title", title).
		Set("ingredients", ingredientsJSON).
		Set("recipe_text", recipeText).
		Where(squirrel.Eq{"id": recipeID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *RecipeRepo) DeleteRecipe(ctx context.Context, recipeID int) error {
	query, args, err := squirrel.Delete("recipes").
		Where(squirrel.Eq{"id": recipeID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}
