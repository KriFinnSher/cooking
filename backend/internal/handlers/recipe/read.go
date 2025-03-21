package recipe

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) GetRecipe(ctx echo.Context) error {
	recipeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid recipe ID")
	}

	recipeDetails, err := h.RecipeUseCase.GetRecipeDetails(ctx.Request().Context(), recipeID)
	if err != nil {
		return ctx.JSON(http.StatusNotFound, "recipe not found")
	}

	response := Response{
		ID:          recipeDetails.ID,
		Title:       recipeDetails.Title,
		Ingredients: recipeDetails.Ingredients,
		RecipeText:  recipeDetails.RecipeText,
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllUserRecipes(ctx echo.Context) error {
	username := ctx.Get("username").(string)
	user, err := h.UserUseCase.GetUserInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "user not found")
	}

	recipes, err := h.RecipeUseCase.GetAllUserRecipes(ctx.Request().Context(), user.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	var response []Response
	for _, recipe := range recipes {
		response = append(response, Response{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Ingredients: recipe.Ingredients,
			RecipeText:  recipe.RecipeText,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}

func (h *Handler) GetAllRecipes(ctx echo.Context) error {

	recipes, err := h.RecipeUseCase.GetAllRecipes(ctx.Request().Context())
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	var response []Response
	for _, recipe := range recipes {
		response = append(response, Response{
			ID:          recipe.ID,
			Title:       recipe.Title,
			Ingredients: recipe.Ingredients,
			RecipeText:  recipe.RecipeText,
		})
	}

	return ctx.JSON(http.StatusOK, response)
}
