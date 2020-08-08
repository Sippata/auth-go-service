package mongo

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
)

// DBInstance holds mongo client
type DBInstance struct {
	Client *mongo.Client
}

// Open mongodb connection
func (db *DBInstance) Open() error {
	var err error
	ctx := context.Background()
	connString := os.Getenv("MONGO_CONNECTION_STRING")
	log.Print("Connecting to MongoDB with uri: " + connString)
	db.Client, err = mongo.Connect(ctx, options.Client().ApplyURI(connString))
	if err != nil {
		return err
	}

	if err = db.Client.Ping(ctx, readpref.Primary()); err != nil {
		return err
	}
	log.Print("MongoDB connceted")
	return nil
}

// Close mongodb connection
func (db *DBInstance) Close() {
	log.Print("Disconnecting from MongoDB")
	if err := db.Client.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	log.Print("MongoDB disconnected")
}

// GetCollection return mongo collection
func (db *DBInstance) GetCollection(dbName string, collectionName string) *mongo.Collection {
	wcMajority := writeconcern.New(writeconcern.WMajority(), writeconcern.WTimeout(1*time.Second))
	collectionOpts := options.Collection().SetWriteConcern(wcMajority)
	return db.Client.Database(dbName).Collection(collectionName, collectionOpts)
}

// WithTransaction execute operation(s) inside a transaction
func (db *DBInstance) WithTransaction(execute func() error) error {
	session, err := db.Client.StartSession()
	if err != nil {
		return err
	}
	callback := func(sessionCtx mongo.SessionContext) (interface{}, error) {
		if err := execute(); err != nil {
			return nil, err
		}
		return nil, err
	}
	if _, err := session.WithTransaction(context.Background(), callback); err != nil {
		return err
	}
	return nil
}
