package postgres

import (
	"context"
	"cooking/backend/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type SubscriptionRepo struct {
	db *sqlx.DB
}

func NewSubscriptionRepo(db *sqlx.DB) *SubscriptionRepo {
	return &SubscriptionRepo{db: db}
}

func (r *SubscriptionRepo) CreateSubscription(ctx context.Context, userID, chefID int) error {
	query, args, err := squirrel.Insert("subscriptions").
		Columns("user_id", "chef_id").
		Values(userID, chefID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *SubscriptionRepo) GetUserSubscriptions(ctx context.Context, userID int) ([]models.Subscription, error) {
	query, args, err := squirrel.Select("user_id", "chef_id").
		From("subscriptions").
		Where(squirrel.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var subscriptions []models.Subscription
	err = r.db.SelectContext(ctx, &subscriptions, query, args...)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}

func (r *SubscriptionRepo) GetChefSubscriptions(ctx context.Context, chefID int) ([]models.Subscription, error) {
	query, args, err := squirrel.Select("user_id", "chef_id").
		From("subscriptions").
		Where(squirrel.Eq{"chef_id": chefID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var subscriptions []models.Subscription
	err = r.db.SelectContext(ctx, &subscriptions, query, args...)
	if err != nil {
		return nil, err
	}

	return subscriptions, nil
}
