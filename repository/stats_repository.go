package repository

import (
	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/pkg/mongo"
)

type StatsRepository struct {
	database   mongo.Database
	collection string
}

func NewStatsRepository(db mongo.Database, collection string) domain.StatsRepository {
	return &StatsRepository{
		database:   db,
		collection: collection,
	}
}
