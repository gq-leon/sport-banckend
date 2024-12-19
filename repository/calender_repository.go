package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/pkg/mongo"
)

type CalenderRepository struct {
	database   mongo.Database
	collection string
}

func NewCalenderRepository(db mongo.Database, collection string) domain.CalenderRepository {
	return &CalenderRepository{
		database:   db,
		collection: collection,
	}
}

func (cr *CalenderRepository) GetTrainPlansByDate(ctx context.Context, userID, startDate, endDate string) ([]domain.TrainPlan, error) {
	var (
		result       []domain.TrainPlan
		idFromHex, _ = primitive.ObjectIDFromHex(userID)
	)

	filter := bson.M{
		"user_id": idFromHex,
		"date": bson.M{
			"$gte": startDate,
			"$lte": endDate,
		},
	}

	cursor, err := cr.database.Collection(cr.collection).Find(ctx, filter)
	if err != nil {
		return result, err
	}

	err = cursor.All(ctx, &result)
	return result, err
}
