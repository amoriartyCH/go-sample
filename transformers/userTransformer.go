package transformers

import (
	"github.com/amoriartyCH/go-sample/models/entity"
	"github.com/amoriartyCH/go-sample/models/rest"
)

type UserTransformer interface {
	ToRest(entity *entity.UserDao) *rest.UserRest
	ToEntity(rest *rest.UserRest) *entity.UserDao
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

func (u *UserTransformerImpl) ToRest(entity *entity.UserDao) *rest.UserRest {
	return &rest.UserRest{
		FirstName: entity.FirstName,
		LastName:  entity.LastName,
	}
}

func (u *UserTransformerImpl) ToEntity(rest *rest.UserRest) *entity.UserDao {
	return &entity.UserDao{
		ID:			rest.ID,
		FirstName: 	rest.FirstName,
		LastName:  	rest.LastName,
	}
}
