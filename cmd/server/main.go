package main

import (
	"github.com/gabehamasaki/gopay/config"
	"github.com/gabehamasaki/gopay/internal/router"
)

func main() {
	// Initialize Configuration
	config.Init()

	// Initialize Router
	router.Initialize()
}
