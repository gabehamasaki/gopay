package router

import (
	"fmt"
	"os"

	"github.com/gabehamasaki/gopay/config"
	"github.com/gin-gonic/gin"
)

var (
	logger *config.Logger
)

func Initialize() {

	// Initialize Logger
	logger = config.GetLogger("router")

	router := gin.Default()

	// Mapping application routes
	logger.Infof("Mapping application routes...")
	initializeRoutes(router)
	logger.Infof("Mapping application routes successfully")

	run(router)
}

func run(r *gin.Engine) {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	logger.Infof("Starting server on port %s", PORT)
	r.Run(fmt.Sprintf(":%s", PORT))
}
