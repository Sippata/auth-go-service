package mongo

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// MongoDB holds mongo client
type MongoDB struct {
	Client *mongo.Client
}

// Open mongodb connection
func (db *MongoDB) Open() error {
	var err error
	ctx := context.TODO()
	connString := os.Getenv("MONGO_CONNECTION_STRING")
	db.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		return err
	}
	err = db.Client.Ping(ctx, readpref.Primary())
	if err != nil {
		return err
	}
	return nil
}

// Close mongodb connection
func (db *MongoDB) Close() {
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
}
