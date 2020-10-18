package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
	"time"
)

var mgoClient *mongo.Client

// MongoDatabaseInterface is an interface that describes the mongodb driver
type MongoDatabaseInterface interface {
	Collection(name string, opts ...*options.CollectionOptions) *mongo.Collection
}

func getMongoClient(mongoDBURL string) *mongo.Client {

	ctx := context.Background()

	clientOptions := options.Client().ApplyURI(mongoDBURL)
	client, err := mongo.Connect(ctx, clientOptions)

	// the program must bail out here if failing to establish a connection to the db, as this will run on application start-up
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	// cache mongo client here, in preparation for disconnect on application shutdown
	mgoClient = client

	// check we can connect to the mongodb instance - again, bail out on failure
	pingContext, cancel := context.WithDeadline(ctx, time.Now().Add(5*time.Second))
	defer cancel()
	err = client.Ping(pingContext, nil)
	if err != nil {
		log.Error("ping to mongodb timed out. please check the connection to mongodb and that it is running")
		os.Exit(1)
	}

	log.Info("connected to mongodb successfully")

	return client
}

func getMongoDatabase(mongoDBURL, databaseName string) MongoDatabaseInterface {
	return getMongoClient(mongoDBURL).Database(databaseName)
}
