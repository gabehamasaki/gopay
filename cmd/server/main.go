package main

import (
	"github.com/gabehamasaki/gopay/config"
	"github.com/gabehamasaki/gopay/internal/router"
	"github.com/gabehamasaki/gopay/internal/services"
)

func main() {
	// Initialize Configuration
	config.Init()

	// Initialize Services
	services.Initialize()

	// Initialize Router
	router.Initialize()
}
