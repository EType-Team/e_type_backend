package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type IUserRepository interface {
	GetUserById(user *model.User, userId uint) error
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &userRepository{db}
}

func (ur *userRepository) GetUserById(user *model.User, userId uint) error {
	if err := ur.db.Where("id=?", userId).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).First(user).Error; err != nil {
		return err
	}
	return nil
}

func (ur *userRepository) CreateUser(user *model.User) error {
	if err := ur.db.Create(user).Error; err != nil {
		return err
	}
	return nil
}
