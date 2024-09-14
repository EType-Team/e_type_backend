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
	FindOrCreateUserWordProgress(userId uint, wordId uint) (*model.UserWordProgress, error)
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
	}
	return nil
}

func (uwpr *userWordProgressRepository) FindOrCreateUserWordProgress(userId uint, wordId uint) (*model.UserWordProgress, error) {
	userWordProgress := &model.UserWordProgress{}
	err := uwpr.db.Where("user_id=? AND word_id=?", userId, wordId).First(userWordProgress).Error
	if err == gorm.ErrRecordNotFound {
		userWordProgress.UserID = userId
		userWordProgress.WordID = wordId
		userWordProgress.TotalTypings = 1
		if err := uwpr.db.Create(userWordProgress).Error; err != nil {
			return nil, err
		}
	} else if err == nil {
		userWordProgress.TotalTypings += 1
	} else {
		return nil, err
	}
	return userWordProgress, nil
}
