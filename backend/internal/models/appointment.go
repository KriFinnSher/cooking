package models

type Appointment struct {
	UserID     int `json:"user_id" db:"user_id"`
	ScheduleID int `json:"schedule_id" db:"schedule_id"`
}
