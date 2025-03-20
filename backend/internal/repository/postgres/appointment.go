package postgres

import (
	"context"
	"cooking/backend/internal/models"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
)

type AppointmentRepo struct {
	db *sqlx.DB
}

func NewAppointmentRepo(db *sqlx.DB) *AppointmentRepo {
	return &AppointmentRepo{db: db}
}

func (a *AppointmentRepo) CreateAppointment(ctx context.Context, userID, eventID int) error {
	query, args, err := squirrel.Insert("appointments").
		Columns("user_id", "schedule_id").
		Values(userID, eventID).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return err
	}

	_, err = a.db.ExecContext(ctx, query, args...)
	return err
}

func (a *AppointmentRepo) GetUserAppointments(ctx context.Context, userID int) ([]models.Appointment, error) {
	query, args, err := squirrel.Select("*").
		From("appointments").
		Where(squirrel.Eq{"user_id": userID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}
	var appointments []models.Appointment
	err = a.db.SelectContext(ctx, &appointments, query, args...)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}

func (a *AppointmentRepo) GetEventAppointments(ctx context.Context, eventID int) ([]models.Appointment, error) {
	query, args, err := squirrel.Select("*").
		From("appointments").
		Where(squirrel.Eq{"schedule_id": eventID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return nil, err
	}

	var appointments []models.Appointment
	err = a.db.SelectContext(ctx, &appointments, query, args...)
	if err != nil {
		return nil, err
	}

	return appointments, nil
}
