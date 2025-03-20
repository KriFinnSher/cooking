package appointment

import (
	"context"
	"cooking/backend/internal/models"
)

type Repo interface {
	CreateAppointment(ctx context.Context, userID, eventID int) error
	GetUserAppointments(ctx context.Context, userID int) ([]models.Appointment, error)
	GetEventAppointments(ctx context.Context, eventID int) ([]models.Appointment, error)
}
