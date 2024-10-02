package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type IUserWordProgressRepository interface {
	GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error
	GetUserWordProgressByWordId(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error
	CreateOrUpdateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error
	GetUserWordProgressByWordIds(userWordProgress *[]model.UserWordProgress, userId uint, wordIds []uint) error
}

type userWordProgressRepository struct {
	db *gorm.DB
}

func NewUserWordProgressRepository(db *gorm.DB) IUserWordProgressRepository {
	return &userWordProgressRepository{db}
}

func (uwpr *userWordProgressRepository) GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error {
	if err := uwpr.db.Where("user_id=?", userId).Preload("Word").Find(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) GetUserWordProgressByWordId(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error {
	if err := uwpr.db.Where("user_id=? AND word_id=?", userId, wordId).Preload("Word").First(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) CreateOrUpdateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error {
	result := uwpr.db.Where("user_id = ? AND word_id = ?", userId, wordId).First(userWordProgress)

	if result.RowsAffected > 0 {
		userWordProgress.TotalTypings += 1
		if err := uwpr.db.Save(userWordProgress).Error; err != nil {
			return err
		}
	} else {
		userWordProgress.UserID = userId
		userWordProgress.WordID = wordId
		userWordProgress.TotalTypings = 1
		if err := uwpr.db.Create(userWordProgress).Error; err != nil {
			return err
		}
	}

	return nil
}

func (uwpr *userWordProgressRepository) GetUserWordProgressByWordIds(userWordProgress *[]model.UserWordProgress, userId uint, wordIds []uint) error {
	if err := uwpr.db.Where("user_id=? AND word_id IN ?", userId, wordIds).Preload("Word").Find(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}
