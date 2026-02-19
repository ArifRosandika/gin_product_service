package helper

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, message string, data interface{}) {
	c.JSON(200, gin.H{
		"status": "success",
		"message": message,
		"data": data,
	})
}

func ErrorResponse(c *gin.Context, status int, payload gin.H) {
	c.JSON(status, payload)
}