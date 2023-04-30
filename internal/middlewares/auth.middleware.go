package middlewares

import (
	"net/http"

	"github.com/gabehamasaki/gopay/internal/helpers"
	"github.com/gin-gonic/gin"
)

func Authentication(c *gin.Context) {
	token := c.Request.Header.Get("Authorization")
	if token == "" || token == "Bearer undefined" {
		helpers.SendError(c, &helpers.HttpError{
			Message:    "No Authorization header provided",
			StatusCode: http.StatusForbidden,
			Op:         "authentication-middleware",
		})
		c.Abort()
		return
	}
	c.Next()
}
