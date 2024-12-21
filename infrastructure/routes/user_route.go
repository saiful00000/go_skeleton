package routes

import (
	"go_clean/core/controllers"

	"github.com/gin-gonic/gin"
)

func UserRoutes(g *gin.Engine) {
	user := g.Group("/users")
	user.POST("/create", controllers.CreateUser)
	user.GET("/get", controllers.GetUser)
	user.PUT("/update", controllers.UpdateUser)
	user.DELETE("/delete", controllers.DeleteUser)
}