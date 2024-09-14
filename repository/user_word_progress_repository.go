package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type IUserWordProgressRepository interface {
	GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error
	GetUserWordProgressByWordId(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error
	CreateUserWordProgress(userWordProgress *model.UserWordProgress) error
	UpdateUserWordProgress(userWordProgress *model.UserWordProgress) error
	FindOrCreateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error
}

type userWordProgressRepository struct {
	db *gorm.DB
}

func NewUserWordProgressRepository(db *gorm.DB) IUserWordProgressRepository {
	return &userWordProgressRepository{db}
}

func (uwpr *userWordProgressRepository) GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error {
	if err := uwpr.db.Where("user_id=?", userId).Find(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) GetUserWordProgressByWordId(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error {
	if err := uwpr.db.Where("user_id=? AND word_id=?", userId, wordId).First(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) CreateUserWordProgress(userWordProgress *model.UserWordProgress) error {
	if err := uwpr.db.Create(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) UpdateUserWordProgress(userWordProgress *model.UserWordProgress) error {
	if err := uwpr.db.Save(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) FindOrCreateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error {
	if err := uwpr.db.Where("user_id=? AND word_id=?", userId, wordId).FirstOrCreate(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}
