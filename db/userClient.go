package db

import (
	"context"
	"fmt"
	"github.com/amoriartyCH/go-sample/config"
	"github.com/amoriartyCH/go-sample/models/entity"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserClient provides an interface by which to interact with the user database
type UserClient interface {
	CreateUser(entity *entity.UserDao) error
	GetUser(id string) (*entity.UserDao, error)
	GetAllUsers() (*[]*entity.UserDao, error)
	Shutdown()
}

// DatabaseClient is a concrete implementation of the Client interface
type UserDatabaseClient struct {
	db MongoDatabaseInterface
}

/*
	NewUserDatabaseClient returns a new implementation of the UserClient interface.
	Allows interaction with the user collection.
*/
func NewUserDatabaseClient(cfg *config.Config) UserClient {
	return &UserDatabaseClient{
		db: getMongoDatabase(cfg.MongoDBURL, cfg.MongoDBDatabase),
	}
}

// CreateUser creates a user entity in the database
func (c *UserDatabaseClient) CreateUser(entity *entity.UserDao) error {

	collection := c.db.Collection("users")
	_, err := collection.InsertOne(context.Background(), entity)
	return err
}

// GetUser fetches a user from the db according to an id
func (c *UserDatabaseClient) GetUser(id string) (*entity.UserDao, error) {

	var entity entity.UserDao

	collection := c.db.Collection("users")
	dbResource := collection.FindOne(context.Background(), bson.M{"_id": id})

	err := dbResource.Err()
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	err = dbResource.Decode(&entity)

	if err != nil {
		return nil, err
	}

	return &entity, nil
}

// GetAllUsers returns an array of all users in the database
func (c *UserDatabaseClient) GetAllUsers() (*[]*entity.UserDao, error) {

	entities := make([]*entity.UserDao, 0)

	collection := c.db.Collection("users")
	cur, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, err
	}

	for cur.Next(context.Background()) {

		var entity entity.UserDao
		err = cur.Decode(&entity)

		if err != nil {
			return nil, err
		}

		entities = append(entities, &entity)
	}

	return &entities, nil
}

// Shutdown is a hook that can be used to clean up db resources
func (c *UserDatabaseClient) Shutdown() {
	log.Info("Attempting to close the db connection thread pool")
	if mgoClient != nil {
		err := mgoClient.Disconnect(context.Background())
		if err != nil {
			log.Error(fmt.Sprintf("Failed to disconnect from the mongodb: %s", err))
			return
		}
		log.Info("disconnected from mongodb successfully")
	}
}
