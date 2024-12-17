package bootstrap

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/gq-leon/sport-backend/pkg/mongo"
)

func NewMongoDatabase(env *Env) mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongodbURI := fmt.Sprintf("mongodb://%s:%s@%s:%s", env.DBUser, env.DBPass, env.DBHost, env.DBPort)
	if env.DBUser == "" || env.DBPass == "" {
		mongodbURI = fmt.Sprintf("mongodb://%s:%s", env.DBHost, env.DBPort)
	}

	client, err := mongo.NewClient(mongodbURI)
	if err != nil {
		log.Fatal(err)
	}

	if err = client.Connect(ctx); err != nil {
		log.Fatal(err)
	}

	if err = client.Ping(ctx); err != nil {
		log.Fatal(err)
	}

	return client
}

func CloseMongoDBConnection(client mongo.Client) {
	if client == nil {
		return
	}

	if err := client.Disconnect(context.TODO()); err != nil {
		log.Fatal(err)
	}

	log.Println("Connection to MongoDB closed.")
}
