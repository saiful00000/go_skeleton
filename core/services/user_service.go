package services

import (
	"go_clean/core/domain"
	"go_clean/core/models"
)

type UserService struct {
	repo domain.IUserRepo
}

func UserServiceInstance(r domain.IUserRepo) domain.IUserService {
	return UserService{
		repo: r,
	}
}

func (u UserService) CreateUser(user *models.Useer) error {
	panic("unimplemented")
}

func (u UserService) DeleteUser(userId uint) error {
	panic("unimplemented")
}

func (u UserService) GetUser(userId uint) (models.Useer, error) {
	panic("unimplemented")
}

func (u UserService) UpdateUser(user *models.Useer) error {
	panic("unimplemented")
}
