package transformers

import (
	e "github.com/amoriartyCH/go-sample/models/entity"
	r "github.com/amoriartyCH/go-sample/models/rest"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	ID        = "id"
	FirstName = "firstName"
	LastName  = "lastName"
)

func TestUserTransformerImpl_ToEntity(t *testing.T) {

	transformer := NewUserTransformer() // Transformer we are going to test.

	// The first (Top level) convey will pass in a string, t and a function (action).
	Convey("Given I have a fully populated rest resource", t, func(){

		rest := &r.UserRest{
			ID:        ID,
			FirstName: FirstName,
			LastName:  LastName,
		}

		// The next levels that follow will only pass in a string and a function (action).
		Convey("When I transform the rest resource to an entity resource", func(){

			entity := transformer.ToEntity(rest)

			Convey("Then the entity fields all equal the rest fields", func(){
				So(entity.ID, ShouldEqual, rest.ID)
				So(entity.FirstName, ShouldEqual, rest.FirstName)
				So(entity.LastName, ShouldEqual, rest.LastName)
			})
		})
	})
}

func TestUserTransformerImpl_ToRest(t *testing.T) {

	transformer := NewUserTransformer()

	Convey("Given I have a fully populated entity resource", t, func(){

		entity := &e.UserDao{
			ID:        ID,
			FirstName: FirstName,
			LastName:  LastName,
		}

		Convey("When I transform the entity resource into a rest resource", func(){

			rest := transformer.ToRest(entity)

			Convey("Then the rest fields all equal the entity fields except for the Id field", func(){
				So(rest.ID, ShouldEqual, "")
				So(rest.FirstName, ShouldEqual, entity.FirstName)
				So(rest.LastName, ShouldEqual, entity.LastName)
			})
		})
	})
}

func TestNewUserTransformer(t *testing.T) {

	Convey("Given I want to create a new transformer", t, func(){
		transformer := NewUserTransformer()

		Convey("Then the constructor returns me a valid transformer", func(){
			So(transformer, ShouldNotBeNil)
		})
	})
}
