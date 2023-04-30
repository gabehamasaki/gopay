package router

import (
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

func Initialize() {
	router := gin.Default()

	// Inicialize application routes
	initializeRoutes(router)

	run(router)
}

func run(r *gin.Engine) {
	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	r.Run(fmt.Sprintf(":%s", PORT))
}
