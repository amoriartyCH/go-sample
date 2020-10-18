package transformers

import (
	"github.com/amoriartyCH/go-sample/models/user"
)

type UserTransformer interface {
	ToRest(entity *user.UserDao) *user.UserRest
	ToEntity(rest *user.UserRest) *user.UserDao
}

// A concrete implementation of the UserTransformer interface.
type UserTransformerImpl struct {}

/*
	CONSTRUCTOR
	NewUserService returns a new UserTransformerImpl (a concrete implementation of the UserTransformer interface)
*/
func NewUserTransformer() UserTransformer {
	return &UserTransformerImpl{}
}

func (u *UserTransformerImpl) ToRest(entity *user.UserDao) *user.UserRest {
	return &user.UserRest{
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

func (u *UserTransformerImpl) ToEntity(rest *user.UserRest) *user.UserDao {
	return &user.UserDao{
		ID:			rest.ID,
		FirstName: 	rest.FirstName,
		LastName:  	rest.LastName,
	}
}
