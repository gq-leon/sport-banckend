package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/pkg/mongo"
)

type TrainPlanRepository struct {
	database   mongo.Database
	collection string
}

func NewTrainPlanRepository(db mongo.Database, collection string) domain.TrainPlanRepository {
	return &TrainPlanRepository{
		database:   db,
		collection: collection,
	}
}

func (tpr *TrainPlanRepository) Create(ctx context.Context, plan *domain.TrainPlan) error {
	_, err := tpr.database.Collection(tpr.collection).InsertOne(ctx, plan)
	return err
}

func (tpr *TrainPlanRepository) Update(ctx context.Context, id string, plan *domain.TrainPlan) error {
	idFromHex, _ := primitive.ObjectIDFromHex(id)
	_, err := tpr.database.Collection(tpr.collection).UpdateOne(ctx, bson.M{"_id": idFromHex}, bson.M{"$set": plan})
	return err
}

func (tpr *TrainPlanRepository) Delete(ctx context.Context, id string) error {
	idFromHex, _ := primitive.ObjectIDFromHex(id)
	_, err := tpr.database.Collection(tpr.collection).DeleteOne(ctx, bson.M{"_id": idFromHex})
	return err
}

func (tpr *TrainPlanRepository) GetByDate(ctx context.Context, userID string, date string) ([]domain.TrainPlan, error) {
	var results []domain.TrainPlan

	idFromHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return results, err
	}

	cursor, err := tpr.database.Collection(tpr.collection).Find(ctx, bson.M{"user_id": idFromHex, "date": date})
	if err != nil {
		return results, err
	}

	if err = cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

func (tpr *TrainPlanRepository) GetByID(ctx context.Context, ID string) (domain.TrainPlan, error) {
	var plan domain.TrainPlan
	idFromHex, _ := primitive.ObjectIDFromHex(ID)
	err := tpr.database.Collection(tpr.collection).FindOne(ctx, bson.M{"_id": idFromHex}).Decode(&plan)
	return plan, err
}
