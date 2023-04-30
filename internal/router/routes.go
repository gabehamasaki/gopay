package router

import (
	"github.com/gabehamasaki/gopay/internal/controllers"
	"github.com/gabehamasaki/gopay/internal/middlewares"
	"github.com/gabehamasaki/gopay/internal/services"
	"github.com/gin-gonic/gin"
)

var (
	basePath          string
	accountController *controllers.AccountController
)

func initializeRoutes(r *gin.Engine) {
	basePath = "/api/v1"

	initializeControllers()

	v1 := r.Group(basePath)
	{

		public := v1.Group("")
		{
			// Authentication routes
			public.POST("/auth/login", controllers.Login)
			public.POST("/auth/forgot-password", controllers.ForgotPassword)

			// Create Account route
			public.POST("/account", accountController.Create)
		}

		// private routes
		private := v1.Group("", middlewares.Authentication)
		{
			// Account routes
			private.GET("/account", accountController.Show)
			private.PUT("/account", accountController.Update)
			private.DELETE("/account", accountController.Delete)
		}
	}
}

func initializeControllers() {
	accountController = &controllers.AccountController{
		Service: services.NewAccountService(),
	}
}
