package chef

import (
	"context"
	"cooking/backend/internal/models"
)

type Repo interface {
	CreateChef(ctx context.Context, name, hash string) error
	GetChef(ctx context.Context, chefID int) (models.Chef, error)
	GetChefByName(ctx context.Context, name string) (models.Chef, error)
}
