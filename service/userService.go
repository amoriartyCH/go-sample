package service

import (
	"fmt"
	"github.com/amoriartyCH/go-sample/config"
	"github.com/amoriartyCH/go-sample/db"
	"github.com/amoriartyCH/go-sample/models/user"
	"github.com/amoriartyCH/go-sample/transformers"
	"github.com/hashicorp/go-uuid"
	log "github.com/sirupsen/logrus"
)

// UserService provides an interface by which to interact with a User resource
type UserService interface {
	CreateUser(rest *user.UserRest) (ResponseType, error)
	GetUser(id string) (ResponseType, *user.UserRest, error)
	Shutdown()
	//GetAllUsers() (ResponseType, *[]models.User, error)
}

// UserServiceImpl provides a concrete implementation of the UserService interface
type UserServiceImpl struct {
	userClient          db.UserClient
	userTransformer		transformers.UserTransformer
}

/*
	CONSTRUCTOR
	NewUserService returns a new UserServerImpl (a concrete implementation of the UserService interface)
 */
func NewUserService(cfg *config.Config) UserService {
	return &UserServiceImpl{
		userClient:          db.NewUserDatabaseClient(cfg),
		userTransformer: transformers.NewUserTransformer(),
	}
}

func (s *UserServiceImpl) CreateUser(rest *user.UserRest) (ResponseType, error) {

	// Set the ID of the rest object to be stored into the DB.
	id, err := uuid.GenerateUUID()
	if err != nil {
		return Error, err
	}
	rest.ID = id

	// Transform the Rest object into an Entity ready for DB storage.
	entity := s.userTransformer.ToEntity(rest)

	// Call to store the entity.
	err = s.userClient.CreateUser(entity)
	if err != nil {
		log.Errorf(fmt.Sprintf("error when attempting to create user: %s", err))
		return Error, err
	}

	// Return success if we reach this point.
	return Success, nil
}

func (s *UserServiceImpl) GetUser(id string) (ResponseType, *user.UserRest, error) {

	entity, err := s.userClient.GetUser(id)
	if err != nil {
		log.Errorf(fmt.Sprintf("error when attempting to get user: %s", err))
		return Error, &user.UserRest{}, err
	}

	rest := s.userTransformer.ToRest(entity)
	return Success, rest, err
}

func (s *UserServiceImpl) Shutdown() {
	s.userClient.Shutdown()
}

// GET ALL USERS endpoint (Not needed?)
//func (s *UserServiceImpl) GetAllUsers() (ResponseType, *[]models.UserRest, error) {
//
//	users := make([]models.User, 0)
//	return Success, &users, nil
//}