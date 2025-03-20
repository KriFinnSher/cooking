package postgres

import (
	"context"
	"cooking/backend/internal/models"
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
	query, args, err := squirrel.Insert("recipes").
		Columns("user_id", "title", "ingredients", "recipe_text").
		Values(userID, title, ingredients, recipeText).
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

	var recipe models.Recipe
	err = r.db.GetContext(ctx, &recipe, query, args...)
	if err != nil {
		return models.Recipe{}, err
	}

	return recipe, nil
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

	var recipes []models.Recipe
	err = r.db.SelectContext(ctx, &recipes, query, args...)
	if err != nil {
		return nil, err
	}

	return recipes, nil
}

func (r *RecipeRepo) UpdateRecipe(ctx context.Context, recipeID int, title string, ingredients map[string]int, recipeText string) error {
	query, args, err := squirrel.Update("recipes").
		Set("title", title).
		Set("ingredients", ingredients).
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
