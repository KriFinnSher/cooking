package user

import (
	"context"
	"cooking/backend/internal/models"
)

type Repo interface {
	CreateUser(ctx context.Context, name, hash string) error
	GetUser(ctx context.Context, userID int) (models.User, error)
	GetUserByName(ctx context.Context, name string) (models.User, error)
}
