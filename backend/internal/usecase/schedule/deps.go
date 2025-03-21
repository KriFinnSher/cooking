package schedule

import (
	"context"
	"cooking/backend/internal/models"
	"time"
)

type Repo interface {
	CreateEvent(ctx context.Context, title string, date time.Time, place string, chefID int) error
	GetEvent(ctx context.Context, eventID int) (models.Schedule, error)
	GetAllEvents(ctx context.Context) ([]models.Schedule, error)
	UpdateEvent(ctx context.Context, eventID int, title, place string, date time.Time, chefID int) error
	DeleteEvent(ctx context.Context, eventID int) error
}
