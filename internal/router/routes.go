package router

import (
	"github.com/gabehamasaki/gopay/internal/controllers"
	"github.com/gabehamasaki/gopay/internal/middlewares"
	"github.com/gin-gonic/gin"
)

var (
	basePath string
)

func initializeRoutes(r *gin.Engine) {
	basePath = "/api/v1"

	v1 := r.Group(basePath)
	{

		public := v1.Group("")
		{
			// Authentication routes
			public.POST("/auth/login", controllers.Login)
			public.POST("/auth/forgot-password", controllers.ForgotPassword)

			// Create Account route
			public.POST("/account", controllers.CreateAccount)
		}

		// private routes
		private := v1.Group("", middlewares.Authentication)
		{
			// Account routes
			private.GET("/account", controllers.ShowAccount)
			private.PUT("/account", controllers.UpdateAccount)
			private.DELETE("/account", controllers.DeleteAccount)
		}
	}
}
