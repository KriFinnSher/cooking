package middleware

import (
	"cooking/backend/internal/auth"
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		authHeader := c.Request().Header.Get("Authorization")
		if authHeader == "" {
			return echo.NewHTTPError(http.StatusUnauthorized, "missing token")
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token format")
		}

		claims, err := auth.ParseToken(tokenString)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid or expired token")
		}

		if username, ok := claims["username"].(string); ok {
			c.Set("username", username)
		} else {
			return echo.NewHTTPError(http.StatusUnauthorized, "invalid token payload")
		}
		return next(c)
	}
}
