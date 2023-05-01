package main

import (
	"github.com/gabehamasaki/gopay/config"
	"github.com/gabehamasaki/gopay/internal/router"
	"github.com/gabehamasaki/gopay/internal/services"
)

var (
	logger *config.Logger
)

func main() {

	// Initialize logger
	logger = config.GetLogger("main")

	// Initialize Configuration
	logger.Infof("Initializing configuration...")
	if err := config.Init(); err != nil {
		logger.Errorf("Failed to initialize configuration: %v\n", err)
		return
	}
	logger.Infof("Initializing configuration successfully")

	// Initialize Services
	logger.Infof("Initializing services...")
	services.Initialize()
	logger.Infof("Initializing services successfully")

	// Initialize Router
	logger.Infof("Initializing router...")
	router.Initialize()
}
