package postgres

import (
	"context"
	"cooking/backend/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type ChefRepo struct {
	db *sqlx.DB
}

func NewChefRepo(db *sqlx.DB) *ChefRepo {
	return &ChefRepo{db: db}
}

func (c *ChefRepo) CreateChef(ctx context.Context, name, hash string) error {
	query, args, err := squirrel.Insert("chefs").
		Columns("name", "hash").
		Values(name, hash).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}
	_, err = c.db.ExecContext(ctx, query, args...)
	if err != nil {
		return err
	}
	return nil
}

func (c *ChefRepo) GetChef(ctx context.Context, chefID int) (models.Chef, error) {
	query, args, err := squirrel.Select("*").
		From("chefs").
		Where(squirrel.Eq{"chef_id": chefID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Chef{}, err
	}

	var chef models.Chef
	err = c.db.GetContext(ctx, &chef, query, args...)
	if err != nil {
		return models.Chef{}, err
	}
	return chef, nil
}

func (c *ChefRepo) GetChefByName(ctx context.Context, name string) (models.Chef, error) {
	query, args, err := squirrel.Select("*").
		From("chefs").
		Where(squirrel.Eq{"name": name}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Chef{}, err
	}

	var chef models.Chef
	err = c.db.GetContext(ctx, &chef, query, args...)
	if err != nil {
		return models.Chef{}, err
	}
	return chef, nil
}
