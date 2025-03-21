package usecase

import (
	"context"
	"cooking/backend/internal/auth"
	"cooking/backend/internal/models"
	"cooking/backend/internal/usecase/chef"
	"github.com/labstack/echo/v4"
)

type ChefUseCase struct {
	ChefRepo chef.Repo
}

func ChefInstance(repo chef.Repo) *ChefUseCase {
	return &ChefUseCase{ChefRepo: repo}
}

func (c *ChefUseCase) RegisterChef(ctx context.Context, name, password string) (string, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to hash password")
	}

	_, err = c.ChefRepo.GetChefByName(ctx, name)
	if err == nil {
		return "", echo.NewHTTPError(400, "chef with this name already exists")
	}

	err = c.ChefRepo.CreateChef(ctx, name, hash)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to create new chef")
	}

	token, err := auth.GenerateToken(name, true)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to generate token")
	}

	return token, nil
}

func (c *ChefUseCase) AuthenticateChef(ctx context.Context, name, password string) (string, error) {
	chefInstance, err := c.ChefRepo.GetChefByName(ctx, name)
	if err != nil {
		return "", echo.NewHTTPError(401, "chef not found")
	}

	isCorrect := auth.CheckPasswordHash(password, chefInstance.Hash)
	if !isCorrect {
		return "", echo.NewHTTPError(401, "incorrect password")
	}

	token, err := auth.GenerateToken(chefInstance.Name, true)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to generate token")
	}

	return token, nil
}

func (c *ChefUseCase) GetChefInfo(ctx context.Context, name string) (models.Chef, error) {
	return c.ChefRepo.GetChefByName(ctx, name)
}
