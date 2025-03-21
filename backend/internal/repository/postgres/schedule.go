package postgres

import (
	"context"
	"cooking/backend/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"time"
)

type ScheduleRepo struct {
	db *sqlx.DB
}

func NewScheduleRepo(db *sqlx.DB) *ScheduleRepo {
	return &ScheduleRepo{db: db}
}

func (r *ScheduleRepo) CreateEvent(ctx context.Context, title string, date time.Time, place string, chefID int) error {
	query, args, err := squirrel.Insert("schedules").
		Columns("event_name", "event_date", "location", "chef_id").
		Values(title, date, place, chefID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *ScheduleRepo) GetEvent(ctx context.Context, eventID int) (models.Schedule, error) {
	query, args, err := squirrel.Select("id", "event_name", "event_date", "location", "chef_id").
		From("schedules").
		Where(squirrel.Eq{"id": eventID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return models.Schedule{}, err
	}

	var schedule models.Schedule
	err = r.db.GetContext(ctx, &schedule, query, args...)
	if err != nil {
		return models.Schedule{}, err
	}

	return schedule, nil
}

func (r *ScheduleRepo) GetAllEvents(ctx context.Context) ([]models.Schedule, error) {
	query, args, err := squirrel.Select("id", "event_name", "event_date", "location", "chef_id").
		From("schedules").
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var schedules []models.Schedule
	err = r.db.SelectContext(ctx, &schedules, query, args...)
	if err != nil {
		return nil, err
	}

	return schedules, nil
}

func (r *ScheduleRepo) UpdateEvent(ctx context.Context, eventID int, title, place string, date time.Time, chefID int) error {
	query, args, err := squirrel.Update("schedules").
		Set("event_name", title).
		Set("event_date", date).
		Set("location", place).
		Set("chef_id", chefID).
		Where(squirrel.Eq{"id": eventID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}

func (r *ScheduleRepo) DeleteEvent(ctx context.Context, eventID int) error {
	query, args, err := squirrel.Delete("schedules").
		Where(squirrel.Eq{"id": eventID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = r.db.ExecContext(ctx, query, args...)
	return err
}
