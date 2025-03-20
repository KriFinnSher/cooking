package recipe

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) UpdateRecipe(ctx echo.Context) error {
	recipeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid recipe ID")
	}

	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	err = h.RecipeUseCase.UpdateRecipeDetails(ctx.Request().Context(), recipeID, req.Title, req.Ingredients, req.RecipeText)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Recipe updated successfully")
}
