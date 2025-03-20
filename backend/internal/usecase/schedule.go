package usecase

import (
	"context"
	"cooking/backend/internal/models"
	"cooking/backend/internal/usecase/schedule"
	"time"
)

type ScheduleUseCase struct {
	ScheduleRepo schedule.Repo
}

func ScheduleInstance(repo schedule.Repo) *ScheduleUseCase {
	return &ScheduleUseCase{ScheduleRepo: repo}
}

func (s *ScheduleUseCase) CreateEvent(ctx context.Context, title string, date time.Time, place string, chefID int) error {
	return s.ScheduleRepo.CreateEvent(ctx, title, date, place, chefID)
}

func (s *ScheduleUseCase) GetEventDetails(ctx context.Context, eventID int) (models.Schedule, error) {
	event, err := s.ScheduleRepo.GetEvent(ctx, eventID)
	if err != nil {
		return models.Schedule{}, err
	}
	return event, nil
}

func (s *ScheduleUseCase) GetChefEvents(ctx context.Context, chefID int) ([]models.Schedule, error) {
	events, err := s.ScheduleRepo.GetChefEvents(ctx, chefID)
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (s *ScheduleUseCase) UpdateEventDetails(ctx context.Context, eventID int, title string, date time.Time, place string, chefID int) error {
	err := s.ScheduleRepo.UpdateEvent(ctx, eventID, title, place, date, chefID)
	if err != nil {
		return err
	}
	return nil
}

func (s *ScheduleUseCase) RemoveEvent(ctx context.Context, eventID int) error {
	err := s.ScheduleRepo.DeleteEvent(ctx, eventID)
	if err != nil {
		return err
	}
	return nil
}
