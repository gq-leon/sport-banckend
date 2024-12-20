package usecase

import (
	"context"
	"log/slog"
	"strconv"
	"time"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/internal/util"
	"github.com/gq-leon/sport-backend/pkg/redis"
)

type statsUseCase struct {
	contextTimeout  time.Duration
	statsRepository domain.StatsRepository
}

func NewStatsUseCase(statsRepository domain.StatsRepository, timeout time.Duration) domain.StatsUseCase {
	return &statsUseCase{
		contextTimeout:  timeout,
		statsRepository: statsRepository,
	}
}

func (su *statsUseCase) defaultStats() []domain.ProfileStats {
	return []domain.ProfileStats{
		{Label: "连续天数", Value: "5", Type: "streak"},
		{Label: "运动天数", Value: "5", Type: "days"},
		{Label: "累计时长(h)", Value: "36", Type: "hours"},
	}
}

func (su *statsUseCase) GetProfileStats(ctx context.Context, userID string) ([]domain.ProfileStats, error) {
	var data []domain.ProfileStats

	checkInData := su.getCheckInData(ctx, userID)
	if len(checkInData) == 0 {
		return su.defaultStats(), nil
	}

	data = append(data, domain.ProfileStats{Label: "运动天数", Value: strconv.Itoa(len(checkInData)), Type: "days"})
	data = append(data, su.getStreak(checkInData))
	data = append(data, domain.ProfileStats{Label: "累计时长(h)", Value: "暂无", Type: "hours"})

	return data, nil
}

func (su *statsUseCase) getCheckInData(ctx context.Context, userID string) []int {
	var (
		checkInDays []int
		now         = time.Now()
		yearDay     = now.YearDay() - 1
	)

	checkInKey := util.GenerateCheckInKey(userID, now)
	bytes, err := redis.Client.GetRange(ctx, checkInKey, 0, 366).Bytes()
	if err != nil {
		slog.ErrorContext(ctx, "未查询到 %s 的签到数据", userID)
		return checkInDays
	}

	for i := yearDay; i >= 0; i-- {
		index := i / 8
		bitIndex := i % 8

		if ok := bytes[index] >> uint(7-bitIndex) & 1; ok != 0 {
			checkInDays = append(checkInDays, i+1)
		}
	}

	return checkInDays
}

func (su *statsUseCase) getStreak(data []int) domain.ProfileStats {
	var yearDay = time.Now().YearDay()

	var streak int

	for _, datum := range data {
		if datum != yearDay {
			break
		}

		streak += 1
		yearDay -= 1
	}

	return domain.ProfileStats{Label: "连续天数", Value: strconv.Itoa(streak), Type: "streak"}

}
