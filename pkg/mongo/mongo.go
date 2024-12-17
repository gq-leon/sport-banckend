package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
)

type mongoSession struct {
	mongo.Session
}

type mongoCursor struct {
	mc *mongo.Cursor
}

type mongoSingleResult struct {
	sr *mongo.SingleResult
}

func (msr *mongoSingleResult) Decode(v interface{}) error {
	return msr.sr.Decode(v)
}

func (mr *mongoCursor) Close(ctx context.Context) error {
	return mr.mc.Close(ctx)
}

func (mr *mongoCursor) Next(ctx context.Context) bool {
	return mr.mc.Next(ctx)
}

func (mr *mongoCursor) Decode(val interface{}) error {
	return mr.mc.Decode(val)
}

func (mr *mongoCursor) All(ctx context.Context, results interface{}) error {
	return mr.mc.All(ctx, results)
}
