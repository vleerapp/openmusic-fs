package helpers

import "github.com/gin-gonic/gin"

func CreateError(message string, details *string) gin.H {
	return gin.H{
		"error":   message,
		"details": details,
	}
}
