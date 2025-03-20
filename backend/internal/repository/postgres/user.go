package postgres

import (
	"context"
	"cooking/backend/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) CreateUser(ctx context.Context, name, hash string) error {
	query, args, err := squirrel.Insert("users").
		Columns("name", "hash").
		Values(name, hash).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *UserRepo) GetUser(ctx context.Context, userID int) (models.User, error) {
	query, args, err := squirrel.Select("id", "name", "hash", "avatar").
		From("users").
		Where(squirrel.Eq{"id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

func (r *UserRepo) GetUserByName(ctx context.Context, name string) (models.User, error) {
	query, args, err := squirrel.Select("id", "name", "hash", "avatar").
		From("users").
		Where(squirrel.Eq{"name": name}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.User{}, err
	}

	var user models.User
	err = r.db.GetContext(ctx, &user, query, args...)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
