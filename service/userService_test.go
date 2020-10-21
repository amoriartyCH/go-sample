package service

import (
	"errors"
	. "github.com/amoriartyCH/go-sample/db/mocks"
	e "github.com/amoriartyCH/go-sample/models/entity"
	r "github.com/amoriartyCH/go-sample/models/rest"
	. "github.com/amoriartyCH/go-sample/transformers/mocks"
	"github.com/golang/mock/gomock"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

const (
	id = "id"
	firstName = "firstName"
	lastName = "lastName"
)

func TestNewUserService(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	transformer := NewMockUserTransformer(mockCtrl)
	client := NewMockUserClient(mockCtrl)

	Convey("Given I want to create a new userService and I have a correctly configured config", t, func(){

		svc := &UserServiceImpl{
			userTransformer: transformer,
			userClient: client,
		}

		Convey("Then the constructor returns me a valid userService", func(){
			So(svc, ShouldNotBeNil)
			So(svc.userTransformer, ShouldNotBeNil)
			So(svc.userClient, ShouldNotBeNil)
		})
	})
}

func TestUserServiceImpl_CreateUser(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mTransformer := NewMockUserTransformer(mockCtrl)
	mClient := NewMockUserClient(mockCtrl)

	svc := &UserServiceImpl{
		userClient: mClient,
		userTransformer: mTransformer,
	}

	// ---------- Error Path - userClient fails to save to Database ---------- \\
	Convey("Given I want to create a user and I have the correct rest resource", t, func(){

		rest := &r.UserRest{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}

		Convey("When the transformer is called, it successfully returns the dao ready entity", func() {

			entity := &e.UserDao{
				ID:        id,
				FirstName: firstName,
				LastName:  lastName,
			}

			mTransformer.EXPECT().ToEntity(rest).Return(entity)

			Convey("Then the userClient throws an error when attempting to store the user and returns errors", func() {

				ucError := errors.New("error when attempting to save new user")
				mClient.EXPECT().CreateUser(entity).Return(ucError)
				r, err := svc.CreateUser(rest)

				So(r, ShouldEqual, Error)
				So(err, ShouldEqual, ucError)
			})
		})
	})


	// ---------- Success Path - No errors ---------- \\
	Convey("Given I want to create a user and I have the correct rest resource", t, func(){

		rest := &r.UserRest{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}

		Convey("When the transformer is called, it successfully returns the dao ready entity", func() {

			entity := &e.UserDao{
				ID:        id,
				FirstName: firstName,
				LastName:  lastName,
			}

			mTransformer.EXPECT().ToEntity(rest).Return(entity)

			Convey("Then the userClient successfully stores the user and returns nil errors", func() {

				mClient.EXPECT().CreateUser(entity).Return(nil)
				r, err := svc.CreateUser(rest)

				So(r, ShouldEqual, Success)
				So(err, ShouldEqual, nil)
			})
		})
	})
}

func TestUserServiceImpl_GetUser(t *testing.T) {

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mTransformer := NewMockUserTransformer(mockCtrl)
	mClient := NewMockUserClient(mockCtrl)

	svc := &UserServiceImpl{
		userClient: mClient,
		userTransformer: mTransformer,
	}

	// ---------- Error Path - userClient returns no user by id exists ---------- \\
	Convey("Given I want to get a user but the Id doesn't exist", t, func() {

		wrongId := "wrongId"

		Convey("When I attempt to get that user", func() {

			mClient.EXPECT().GetUser(wrongId).Return(nil, nil)

			r, u, err := svc.GetUser(wrongId)

			Convey("Then response is not found, and the returned user and error are nil", func() {

				So(r, ShouldEqual, NotFound)
				So(u, ShouldEqual, nil)
				So(err, ShouldEqual, nil)
			})
		})
	})

	// ---------- Error Path - userClient fails to retrieve the user ---------- \\
	Convey("Given I want to get a user and I have a valid userId", t, func() {

		// Const Id will be used from const block above.

		Convey("When I attempt to get that user", func() {

			ucError := errors.New("error when attempting to get user with given Id")
			mClient.EXPECT().GetUser(id).Return(nil, ucError)

			r, u, err := svc.GetUser(id)

			Convey("THen the response is Error, and we have an error returned", func() {

				So(r, ShouldEqual, Error)
				So(u, ShouldEqual, nil)
				So(err, ShouldEqual, ucError)
			})
		})
	})

	// ---------- Success Path - No errors ---------- \\
	Convey("Given I want to get a user and I have a valid userId", t, func() {

		rest := &r.UserRest{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}
		entity := &e.UserDao{
			ID:        id,
			FirstName: firstName,
			LastName:  lastName,
		}

		// Const Id will be used from const block above.

		Convey("When I attempt to get that user", func() {

			mClient.EXPECT().GetUser(id).Return(entity, nil)
			mTransformer.EXPECT().ToRest(entity).Return(rest)
			r, u, err := svc.GetUser(id)

			Convey("Then the response is Success, and our user is returned", func() {

				So(r, ShouldEqual, Success)
				So(u, ShouldEqual, rest)
				So(err, ShouldEqual, nil)
			})
		})
	})}