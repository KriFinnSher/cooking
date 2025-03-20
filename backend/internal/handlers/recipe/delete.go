package recipe

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func (h *Handler) DeleteRecipe(ctx echo.Context) error {
	recipeID, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid recipe ID")
	}

	err = h.RecipeUseCase.RemoveRecipe(ctx.Request().Context(), recipeID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusOK, "Recipe deleted successfully")
}
