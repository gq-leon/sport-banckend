package usecase

import (
	"context"
	"time"

	"github.com/gq-leon/sport-backend/domain"
)

type attendanceUseCase struct {
	contextTimeout       time.Duration
	attendanceRepository domain.AttendanceRepository
}

func NewAttendanceUseCase(repository domain.AttendanceRepository, timeout time.Duration) domain.AttendanceUseCase {
	return &attendanceUseCase{
		contextTimeout:       timeout,
		attendanceRepository: repository,
	}
}

func (au *attendanceUseCase) Create(ctx context.Context, attendance *domain.Attendance) error {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	return au.attendanceRepository.Create(ctx, attendance)
}

func (au *attendanceUseCase) Fetch(ctx context.Context, userID string) ([]domain.Attendance, error) {
	ctx, cancel := context.WithTimeout(ctx, au.contextTimeout)
	defer cancel()
	return au.attendanceRepository.GetRecordByUserID(ctx, userID)
}
