package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/pkg/mongo"
)

type AttendanceRepository struct {
	database   mongo.Database
	collection string
}

func NewAttendanceRepository(db mongo.Database, collection string) domain.AttendanceRepository {
	return &AttendanceRepository{
		database:   db,
		collection: collection,
	}
}

func (ar *AttendanceRepository) Create(ctx context.Context, attendance *domain.Attendance) error {
	_, err := ar.database.Collection(ar.collection).InsertOne(ctx, attendance)
	return err
}

func (ar *AttendanceRepository) GetRecordByUserID(ctx context.Context, userID string) ([]domain.Attendance, error) {
	var result []domain.Attendance
	idFromHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return result, err
	}

	// 倒序查询
	findOptions := options.Find().SetSort(bson.D{{Key: "_id", Value: -1}}).SetLimit(5)
	cursor, err := ar.database.Collection(ar.collection).Find(ctx, bson.M{"user_id": idFromHex}, findOptions)
	if err != nil {
		return result, err
	}

	err = cursor.All(ctx, &result)
	return result, err
}
