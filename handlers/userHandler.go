package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/amoriartyCH/go-sample/models/rest"
	"github.com/amoriartyCH/go-sample/service"
	"github.com/gorilla/mux"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// --------------------------- CREATE USER --------------------------- \\
// CreateUserHandler offers a handler by which to create a user
type CreateUserHandler struct {
	service service.UserService
}

/*
	CONSTRUCTOR
	NewCreateUserHandler returns a new CreateUserHandler
*/
func NewCreateUserHandler(service service.UserService) CreateUserHandler {
	return CreateUserHandler{
		service,
	}
}

func (c CreateUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// First we get the data from the request (r) and attempt to decode it into our user object.
	var user rest.UserRest
	err := json.NewDecoder(r.Body).Decode(&user)

	// If the decoding fails, it will most likely be due to bad data being submitted by the user.
	if err != nil {
		log.Error(fmt.Sprintf("Failed to decode request body to user struct: %v", err))
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Next we call the service to attempt to create the new user.
	responseType, err := c.service.CreateUser(&user)

	// Check the response type for any bad status codes returned.
	if responseType == service.Error {
		log.Error(fmt.Sprintf("Error encountered when creating user: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if responseType == service.Conflict {
		log.Info("Attempt made to create a user that already exists")
		w.WriteHeader(http.StatusConflict)
		return
	}

	// Finally we log a success message as by now we know the request was successful and we return the new user.
	log.Info("User created successfully")
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		log.Error(fmt.Sprintf("Error writing response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// --------------------------------------------------------------------------------- \\

// --------------------------- GET USER --------------------------- \\
// GetUserHandler offers a handler by which to fetch a user
type GetUserHandler struct {
	service service.UserService
}

/*
	CONSTRUCTOR
	NewGetUserHandler returns a new GetUserHandler
*/
func NewGetUserHandler(service service.UserService) GetUserHandler {
	return GetUserHandler{
		service,
	}
}

func (g GetUserHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// First we get the id of the user we are trying to find from the url parameters.
	vars := mux.Vars(r)
	userID := vars["user_id"]
	if userID == "" {
		log.Info("No userID in url")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Now we attempt to get the user by calling the service.
	responseType, u, err := g.service.GetUser(userID)

	// Next we check for any error responses returned and handle them.
	if responseType == service.Error {
		log.Error(fmt.Sprintf("Error encountered when fetching user: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if responseType == service.NotFound {
		log.Info("User not found")
		log.Debug(fmt.Sprintf("User not found by id: %s", userID))
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// By now we know the request was successful so we log a successful message.
	log.Info("User fetched successfully")
	log.Debug(fmt.Sprintf("User found with id: %s", userID))

	// Finally we return the found user back.
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(u)
	if err != nil {
		log.Error(fmt.Sprintf("Error writing response: %v", err))
		w.WriteHeader(http.StatusInternalServerError)
	}
}

// --------------------------------------------------------------------------------- \\