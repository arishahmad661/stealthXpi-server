package utils

import (
	"github.com/gin-gonic/gin"
)

func SuccessResponse(c *gin.Context, status int, message string, data interface{}) {
	c.JSON(status, gin.H{
		"message": message,
		"data":    data,
	})
}

func FailResponse(c *gin.Context, status int, userMsg string, err error) {
	c.JSON(status, gin.H{
		"error": userMsg,
		"debug": err.Error(), // Remove in prod
	})
}
