package containers

import (
	"go_clean/core/controllers"
	"go_clean/core/repository"
	"go_clean/core/services"
	"go_clean/infrastructure/configs"
	"go_clean/infrastructure/connections"
	"go_clean/infrastructure/routes"

	"github.com/gin-gonic/gin"
)

func Serve(g *gin.Engine) {
	// Load the configuration from environment variable
	configs.SetConfig()

	// Get the database
	db := connections.GetDB()

	// user
	userRepo := repository.UserRepoInstance(db)
	userService := services.UserServiceInstance(userRepo)
	controllers.SetUserService(userService)

	// setup routes
	routes.UserRoutes(g)

	// Run the server
	g.Run()
}