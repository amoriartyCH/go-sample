package main

import (
	"fmt"
	"github.com/amoriartyCH/go-sample/config"
	"github.com/amoriartyCH/go-sample/handlers"
	"github.com/amoriartyCH/go-sample/service"
	"github.com/gorilla/mux"
	"net/http"

	log "github.com/sirupsen/logrus"
	"os"
)

func main() {

	cfg, err := config.Get()
	if err != nil {
		log.Errorf(fmt.Sprintf("error configuring service: %s. Exiting", err))
		os.Exit(1)
	}

	// Export LOG_LEVEL to change log level.
	setLogLevel(cfg)

	// Create our router used to handle our application routes.
	mainRouter := mux.NewRouter()

	// Create our user service.
	svc := service.NewUserService(cfg)

	// Feed our router and service into our register function.
	handlers.RegisterHandlers(mainRouter, svc)

	// Finally start our application.
	server := &http.Server{
		Addr:    ":3000",
		Handler: mainRouter,
	}

	server.ListenAndServe()
}

func setLogLevel(cfg *config.Config) {

	if cfg.LogLevel != "" {
		log.Info(fmt.Sprintf("Log level set in environment, attempting to set log level to: %s", cfg.LogLevel))
		lvl, err := log.ParseLevel(cfg.LogLevel)
		if err != nil {
			log.Error(fmt.Sprintf("failed to set log level: %s. Exiting", err))
			os.Exit(1)
		}
		log.SetLevel(lvl)
		log.Info("Log level set successfully")
	}
}