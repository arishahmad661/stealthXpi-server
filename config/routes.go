package config

import (
	"github.com/arishahmad661/stealth_x_pi/internal/auth"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()
	r.POST("/auth/signup", auth.SignUpHandler)
	return r
}
