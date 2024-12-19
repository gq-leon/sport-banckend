package usecase

import (
	"context"
	"time"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/timeutil"
)

var format = "2006-01-02"

type calenderUseCase struct {
	contextTimeout     time.Duration
	calenderRepository domain.CalenderRepository
}

func NewCalenderUseCase(repository domain.CalenderRepository, timeout time.Duration) domain.CalenderUseCase {
	return &calenderUseCase{
		contextTimeout:     timeout,
		calenderRepository: repository,
	}
}

func (cu *calenderUseCase) GetDayWorkouts(ctx context.Context, userID string, date string) ([]domain.TrainPlan, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()
	return cu.calenderRepository.GetTrainPlansByDate(ctx, userID, date, date)
}

func (cu *calenderUseCase) GetMonthWorkouts(ctx context.Context, userID, date string) ([]domain.TrainPlan, error) {
	ctx, cancel := context.WithTimeout(ctx, cu.contextTimeout)
	defer cancel()

	startDate, endDate, err := timeutil.GetFirstAndLastDayOfMonth(date, format)
	if err != nil {
		return nil, err
	}

	return cu.calenderRepository.GetTrainPlansByDate(ctx, userID, startDate.Format(format), endDate.Format(format))
}
