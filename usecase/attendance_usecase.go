package usecase

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/util"
	"github.com/gq-leon/sport-backend/pkg/redis"
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

func (au *attendanceUseCase) CheckIn(ctx context.Context, userID string) error {
	var (
		now     = time.Now()
		yearDay = now.YearDay()
		key     = fmt.Sprintf("checkin_%d-%s", now.Year(), userID)
	)

	return redis.Client.SetBit(ctx, key, int64(yearDay-1), 1).Err()
}

func (au *attendanceUseCase) BackDateCheckIn(ctx context.Context, userID string, date []string) error {
	var errStr string
	for _, day := range date {
		parse, err := time.Parse("2006-01-02", day)
		if err != nil {
			errStr = fmt.Sprintf("%s | %s", errStr, err.Error())
			continue
		}

		var (
			yearDay = parse.YearDay()
			key     = util.GenerateCheckInKey(userID)
		)
		if err = redis.Client.SetBit(ctx, key, int64(yearDay-1), 1).Err(); err != nil {
			errStr = fmt.Sprintf("%s | %s", errStr, err.Error())
		}
	}

	if errStr != "" {
		return errors.New(errStr)
	}
	return nil
}
