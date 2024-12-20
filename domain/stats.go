package domain

import "context"

type ProfileStats struct {
	Label string `json:"label"`
	Value string `json:"value"`
	Type  string `json:"type"`
}

type StatsRepository interface {
}

type StatsUseCase interface {
	GetProfileStats(ctx context.Context, userID string) ([]ProfileStats, error)
}
