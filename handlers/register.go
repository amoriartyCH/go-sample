package handlers

import (
	"github.com/amoriartyCH/go-sample/service"
	"github.com/gorilla/mux"
	"net/http"
)

// This function will be called from main.go to register our routes on the server.
func RegisterHandlers(router *mux.Router, userService service.UserService) {
	router.Handle("/users", NewCreateUserHandler(userService)).Methods(http.MethodPost)
	router.Handle("/users/{user_id}", NewGetUserHandler(userService)).Methods(http.MethodGet)
}
