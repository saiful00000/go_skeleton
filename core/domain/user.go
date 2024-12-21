package domain

import "go_clean/core/models"

type IUserRepo interface {
	CreateUser(user *models.Useer) error
	GetUser(userId uint) models.Useer
	UpdateUser(user *models.Useer) error
	DeleteUser(userId uint) error
}

type IUserService interface {
	CreateUser(user *models.Useer) error
	GetUser(userId uint) (models.Useer, error)
	UpdateUser(user *models.Useer) error
	DeleteUser(userId uint) error
}