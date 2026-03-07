package helper

import (
	"github.com/gin-gonic/gin"
	"strings"
	"github.com/go-playground/validator/v10"
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

func FormatValidationError(err error) map[string]string {
	
	errors := make(map[string]string)

	for _, e := range err.(validator.ValidationErrors) {
		field := strings.ToLower(e.Field())
		switch e.Tag() {

		case "required":
			errors[field] = field + " is required"
		case "gt":
			errors[field] = field + " must be greater than zero"
		default:
			errors[field] = field + " is invalid"
		}
	}
	return errors
}