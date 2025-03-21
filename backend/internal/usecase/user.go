package usecase

import (
	"context"
	"cooking/backend/internal/auth"
	"cooking/backend/internal/models"
	"cooking/backend/internal/usecase/user"
	"fmt"
	"github.com/labstack/echo/v4"
)

type UserUseCase struct {
	UserRepo user.Repo
}

func UserInstance(repo user.Repo) *UserUseCase {
	return &UserUseCase{UserRepo: repo}
}

func (u *UserUseCase) RegisterUser(ctx context.Context, name, password string) (string, error) {
	hash, err := auth.HashPassword(password)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to hash password")
	}

	_, err = u.UserRepo.GetUserByName(ctx, name)
	if err == nil {
		return "", echo.NewHTTPError(400, "user with this name already exists")
	}
	fmt.Println(err)

	err = u.UserRepo.CreateUser(ctx, name, hash)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to create new user")
	}

	token, err := auth.GenerateToken(name, false)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to generate token")
	}

	return token, nil
}

func (u *UserUseCase) AuthenticateUser(ctx context.Context, name, password string) (string, error) {
	userInstance, err := u.UserRepo.GetUserByName(ctx, name)
	if err != nil {
		return "", echo.NewHTTPError(401, "user not found")
	}

	isCorrect := auth.CheckPasswordHash(password, userInstance.Hash)
	if !isCorrect {
		return "", echo.NewHTTPError(401, "incorrect password")
	}

	token, err := auth.GenerateToken(userInstance.Name, false)
	if err != nil {
		return "", echo.NewHTTPError(500, "failed to generate token")
	}

	return token, nil
}

func (u *UserUseCase) GetUserInfo(ctx context.Context, name string) (models.User, error) {
	return u.UserRepo.GetUserByName(ctx, name)
}
