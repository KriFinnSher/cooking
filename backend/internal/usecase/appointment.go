package usecase

import (
	"cooking/backend/internal/usecase/appointment"
)

type AppointmentUseCase struct {
	AppointmentRepo appointment.Repo
}

func AppointmentInstance(repo appointment.Repo) *AppointmentUseCase {
	return &AppointmentUseCase{AppointmentRepo: repo}
}

//func (a *AppointmentUseCase) SetAppointment(ctx context.Context, userID, eventID int) error {
//	err := a.AppointmentRepo.CreateAppointment(ctx, userID, eventID)
//	if err != nil {
//		return err
//	}
//	return nil
//}
