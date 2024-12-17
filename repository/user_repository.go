package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/gq-leon/sport-backend/domain"
	"github.com/gq-leon/sport-backend/pkg/mongo"
)

type UserRepository struct {
	database   mongo.Database
	collection string
}

func NewUserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &UserRepository{database: db, collection: collection}
}

func (ur *UserRepository) Create(ctx context.Context, user *domain.User) error {
	_, err := ur.database.Collection(ur.collection).InsertOne(ctx, user)
	return err
}

func (ur *UserRepository) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var user domain.User
	err := ur.database.Collection(ur.collection).FindOne(ctx, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *UserRepository) GetByID(c context.Context, id string) (domain.User, error) {
	var user domain.User
	primitiveID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}

	err = ur.database.Collection(ur.collection).FindOne(c, bson.M{"_id": primitiveID}).Decode(&user)
	return user, err
}
