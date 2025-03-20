package subscription

import (
	"context"
	"cooking/backend/internal/models"
)

type Repo interface {
	CreateSubscription(ctx context.Context, userID, chefID int) error
	GetUserSubscriptions(ctx context.Context, userID int) ([]models.Subscription, error)
	GetChefSubscriptions(ctx context.Context, chefID int) ([]models.Subscription, error)
}
