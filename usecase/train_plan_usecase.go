package usecase

import (
	"context"
	"time"

	"github.com/gq-leon/sport-backend/domain"
)

type trainPlanUseCase struct {
	contextTimeout      time.Duration
	trainPlanRepository domain.TrainPlanRepository
}

func NewTrainPlanUseCase(repository domain.TrainPlanRepository, timeout time.Duration) domain.TrainPlanUseCase {
	return &trainPlanUseCase{
		contextTimeout:      timeout,
		trainPlanRepository: repository,
	}
}

func (tpu *trainPlanUseCase) Create(ctx context.Context, data domain.TrainPlan) error {
	ctx, cancel := context.WithTimeout(ctx, tpu.contextTimeout)
	defer cancel()
	return tpu.trainPlanRepository.Create(ctx, &data)
}

func (tpu *trainPlanUseCase) GetPlansByDate(ctx context.Context, userID string, date string) ([]domain.TrainPlan, error) {
	ctx, cancel := context.WithTimeout(ctx, tpu.contextTimeout)
	defer cancel()
	return tpu.trainPlanRepository.GetByDate(ctx, userID, date)
}

func (tpu *trainPlanUseCase) Delete(ctx context.Context, ID string) error {
	ctx, cancel := context.WithTimeout(ctx, tpu.contextTimeout)
	defer cancel()
	return tpu.trainPlanRepository.Delete(ctx, ID)
}

func (tpu *trainPlanUseCase) GetPlanByID(ctx context.Context, ID string) (domain.TrainPlan, error) {
	ctx, cancel := context.WithTimeout(ctx, tpu.contextTimeout)
	defer cancel()
	return tpu.trainPlanRepository.GetByID(ctx, ID)
}

func (tpu *trainPlanUseCase) UpdateByID(ctx context.Context, ID string, data *domain.TrainPlan) error {
	return tpu.trainPlanRepository.Update(ctx, ID, data)
}
