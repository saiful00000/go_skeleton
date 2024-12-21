package containers

import (
	"go_clean/infrastructure/configs"

	"github.com/gin-gonic/gin"
)

func Serve(g *gin.Engine) {
	// Load the configuration from environment variable
	configs.SetConfig()

	// Get the database

	
}