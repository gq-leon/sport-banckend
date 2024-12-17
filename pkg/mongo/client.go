package mongo

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type mongoClient struct {
	uri string
	cl  *mongo.Client
}

func NewClient(uri string) (Client, error) {
	time.Local = time.UTC
	return &mongoClient{uri: uri}, nil
}

func (mc *mongoClient) Database(name string) Database {
	db := mc.cl.Database(name)
	return &mongoDatabase{db: db}
}

func (mc *mongoClient) Connect(ctx context.Context) (err error) {
	mc.cl, err = mongo.Connect(ctx, options.Client().ApplyURI(mc.uri))
	return
}

func (mc *mongoClient) Disconnect(ctx context.Context) error {
	return mc.cl.Disconnect(ctx)
}

func (mc *mongoClient) StartSession() (mongo.Session, error) {
	session, err := mc.cl.StartSession()
	return &mongoSession{session}, err
}

func (mc *mongoClient) UseSession(ctx context.Context, fn func(mongo.SessionContext) error) error {
	return mc.cl.UseSession(ctx, fn)
}

func (mc *mongoClient) Ping(ctx context.Context) error {
	return mc.cl.Ping(ctx, readpref.Primary())
}
