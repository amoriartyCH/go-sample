package main

import (
	"fmt"
	"github.com/amoriartyCH/go-sample/config"
	//"github.com/amoriartyCH/go-sample/models/user"
	"github.com/amoriartyCH/go-sample/service"
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

	/* TESTING CREATION AND RETRIEVAL OF NEW USER */
	//svc := service.NewUserService(cfg)
	//rest := user.UserRest{
	//	FirstName: "Aaron",
	//	LastName:  "Moriarty",
	//}
	//
	//r, err := svc.CreateUser(&rest)
	//fmt.Println(r, err)
	//r, u, err := svc.GetUser("18958fda-fa77-496e-0ea1-846d1dcc8615")
	//fmt.Println(r, u, err)
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