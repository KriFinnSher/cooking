package recipe

import (
	"cooking/backend/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

type Request struct {
	Title       string         `json:"title,omitempty"`
	Ingredients map[string]int `json:"ingredients,omitempty"`
	RecipeText  string         `json:"recipe_text,omitempty"`
}

type Response struct {
	ID          int            `json:"id"`
	Title       string         `json:"title"`
	Ingredients map[string]int `json:"ingredients"`
	RecipeText  string         `json:"recipe_text"`
}

type Handler struct {
	RecipeUseCase *usecase.RecipeUseCase
	UserUseCase   *usecase.UserUseCase
}

func NewRecipeHandler(recipeUseCase *usecase.RecipeUseCase, userUseCase *usecase.UserUseCase) *Handler {
	return &Handler{RecipeUseCase: recipeUseCase, UserUseCase: userUseCase}
}

func (h *Handler) CreateRecipe(ctx echo.Context) error {
	var req Request
	if err := ctx.Bind(&req); err != nil {
		return ctx.JSON(http.StatusBadRequest, "invalid request data")
	}

	username := ctx.Get("username").(string)
	user, err := h.UserUseCase.GetUserInfo(ctx.Request().Context(), username)
	if err != nil {
		return ctx.JSON(http.StatusUnauthorized, "user not found")
	}

	err = h.RecipeUseCase.SaveRecipe(ctx.Request().Context(), user.ID, req.Title, req.Ingredients, req.RecipeText)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, err.Error())
	}

	return ctx.JSON(http.StatusCreated, "Recipe created successfully")
}
