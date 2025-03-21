package usecase

import (
	"context"
	"cooking/backend/internal/models"
	"cooking/backend/internal/usecase/recipe"
)

type RecipeUseCase struct {
	RecipeRepo recipe.Repo
}

func RecipeInstance(repo recipe.Repo) *RecipeUseCase {
	return &RecipeUseCase{RecipeRepo: repo}
}

func (r *RecipeUseCase) SaveRecipe(ctx context.Context, userID int, title string, ingredients map[string]int, recipeText string) error {
	return r.RecipeRepo.CreateRecipe(ctx, userID, title, ingredients, recipeText)
}

func (r *RecipeUseCase) GetRecipeDetails(ctx context.Context, recipeID int) (models.Recipe, error) {
	recipeInstance, err := r.RecipeRepo.GetRecipe(ctx, recipeID)
	if err != nil {
		return models.Recipe{}, err
	}
	return recipeInstance, nil
}

func (r *RecipeUseCase) GetAllUserRecipes(ctx context.Context, userID int) ([]models.Recipe, error) {
	recipes, err := r.RecipeRepo.GetUserRecipes(ctx, userID)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *RecipeUseCase) GetAllRecipes(ctx context.Context) ([]models.Recipe, error) {
	recipes, err := r.RecipeRepo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return recipes, nil
}

func (r *RecipeUseCase) UpdateRecipeDetails(ctx context.Context, recipeID int, title string, ingredients map[string]int, recipeText string) error {
	err := r.RecipeRepo.UpdateRecipe(ctx, recipeID, title, ingredients, recipeText)
	if err != nil {
		return err
	}
	return nil
}

func (r *RecipeUseCase) RemoveRecipe(ctx context.Context, recipeID int) error {
	err := r.RecipeRepo.DeleteRecipe(ctx, recipeID)
	if err != nil {
		return err
	}
	return nil
}
