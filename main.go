package main

import (
	"github.com/fogtz/goVacations.git/config"
	"github.com/fogtz/goVacations.git/router"
)

var (
	logger config.Logger
)

func main() {
	logger := config.GetLogger("main")
	// Init Configs
	err := config.Init()
	if err != nil {
		logger.Errorf("Error initializing configs: %v", err)
		return
	}

	// Init Router
	router.Init()
}
