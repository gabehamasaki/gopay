package main

import (
	"fmt"

	"github.com/gabehamasaki/gopay/config"
	"github.com/gabehamasaki/gopay/internal/router"
	"github.com/gabehamasaki/gopay/internal/services"
)

func main() {
	// Initialize Configuration
	if err := config.Init(); err != nil {
		fmt.Printf("Failed to initialize configuration: %v\n", err)
		return
	}

	// Initialize Services
	services.Initialize()

	router.Initialize()
}
