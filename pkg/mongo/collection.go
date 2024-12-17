package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoCollection struct {
	coll *mongo.Collection
}

func (mc *mongoCollection) FindOne(ctx context.Context, filter interface{}) SingleResult {
	singleResult := mc.coll.FindOne(ctx, filter)
	return &mongoSingleResult{sr: singleResult}
}

func (mc *mongoCollection) InsertOne(ctx context.Context, document interface{}) (interface{}, error) {
	res, err := mc.coll.InsertOne(ctx, document)
	return res.InsertedID, err
}

func (mc *mongoCollection) InsertMany(ctx context.Context, documents []interface{}) ([]interface{}, error) {
	res, err := mc.coll.InsertMany(ctx, documents)
	if err != nil {
		return nil, err
	}
	return res.InsertedIDs, nil
}

func (mc *mongoCollection) DeleteOne(ctx context.Context, filter interface{}) (int64, error) {
	res, err := mc.coll.DeleteOne(ctx, filter)
	if err != nil {
		return 0, err
	}
	return res.DeletedCount, nil
}

func (mc *mongoCollection) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (Cursor, error) {
	cursor, err := mc.coll.Find(ctx, filter, opts...)
	return &mongoCursor{mc: cursor}, err
}

func (mc *mongoCollection) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return mc.coll.CountDocuments(ctx, filter, opts...)
}

func (mc *mongoCollection) Aggregate(ctx context.Context, pipeline interface{}) (Cursor, error) {
	aggregate, err := mc.coll.Aggregate(ctx, pipeline)
	return &mongoCursor{mc: aggregate}, err
}

func (mc *mongoCollection) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateOne(ctx, filter, update, opts...)
}

func (mc *mongoCollection) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return mc.coll.UpdateMany(ctx, filter, update, opts...)
}
