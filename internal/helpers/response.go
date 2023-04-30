package helpers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, err error) {
	re, _ := err.(*HttpError)

	c.Header("Content-Type", "application/json")
	c.JSON(re.StatusCode, re)
}

func SendSuccess(c *gin.Context, op string, data interface{}) {
	c.Header("Content-Type", "application/json")
	c.JSON(http.StatusOK, gin.H{
		"op":   op,
		"data": data,
	})
}

type HttpError struct {
	Message    string `json:"message"`
	StatusCode int    `json:"code"`
	Op         string `json:"op"`
}

func (h *HttpError) Error() string {
	return fmt.Sprintf("%v: %v", h.Op, h.Message)
}
