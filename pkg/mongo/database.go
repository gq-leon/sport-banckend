package mongo

import "go.mongodb.org/mongo-driver/mongo"

type mongoDatabase struct {
	db *mongo.Database
}

func (md *mongoDatabase) Collection(name string) Collection {
	collection := md.db.Collection(name)
	return &mongoCollection{coll: collection}
}

func (md *mongoDatabase) Client() Client {
	client := md.db.Client()
	return &mongoClient{cl: client}
}
