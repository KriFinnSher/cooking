package middleware

import (
	"cooking/backend/internal/usecase"
	"github.com/labstack/echo/v4"
	"net/http"
)

func CheckChefRole(chefUseCase *usecase.ChefUseCase) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			username := c.Get("username").(string)

			chef, err := chefUseCase.GetChefInfo(c.Request().Context(), username)
			if err != nil || chef.ID == 0 {
				return c.JSON(http.StatusUnauthorized, "You must be a chef to access this resource")
			}

			return next(c)
		}
	}
}
