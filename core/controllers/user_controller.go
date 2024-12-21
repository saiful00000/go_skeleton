package controllers

import (
	"go_clean/core/domain"
	"go_clean/core/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

var UserService domain.IUserService

func SetUserService(uService domain.IUserService) {
	UserService = uService
}


func CreateUser(g *gin.Context) {
	g.JSON(http.StatusOK, "User created.")
	return
}

func GetUser(g *gin.Context) {
	g.JSON(http.StatusOK, models.Useer{ID: 1, UserName: "dummy", Email: "", Password: "", Name: ""})
	return
}


func UpdateUser(g *gin.Context) {
	g.JSON(http.StatusOK, "User updated.")
	return
}


func DeleteUser(g *gin.Context) {
	g.JSON(http.StatusOK, "User deleted.")
	return
}