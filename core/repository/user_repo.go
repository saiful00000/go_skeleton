package repository

import (
	"go_clean/core/domain"
	"go_clean/core/models"

	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

// / bind interface with the BookRepo struct
func UserRepoInstance(d *gorm.DB) domain.IUserRepo {
	return &UserRepo{
		db: d,
	}
}

func (b *UserRepo) CreateUser(user *models.Useer) error {
	panic("unimplemented")
}

func (b *UserRepo) DeleteUser(userId uint) error {
	panic("unimplemented")
}

func (b *UserRepo) GetUser(userId uint) models.Useer {
	panic("unimplemented")
}

func (b *UserRepo) UpdateUser(user *models.Useer) error {
	panic("unimplemented")
}
