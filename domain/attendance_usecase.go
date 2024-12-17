package domain

import "context"

type AttendanceUseCase interface {
	Create(ctx context.Context, attendance *Attendance) error
	Fetch(ctx context.Context, userID string) ([]Attendance, error)
}
