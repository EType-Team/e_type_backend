package repository

import (
	"api/model"
	"fmt"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type IUserWordProgressRepository interface {
	GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error
	GetUserWordProgressById(userWordProgress *model.UserWordProgress, userId uint, userWordProgressId uint) error
	CreateUserWordProgress(userWordProgress *model.UserWordProgress) error
	UpdateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, userWordProgressId uint) error
}

type userWordProgressRepository struct {
	db *gorm.DB
}

func NewUserWordProgressRepository(db *gorm.DB) IUserWordProgressRepository {
	return &userWordProgressRepository{db}
}

func (uwpr *userWordProgressRepository) GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error {
	if err := uwpr.db.Where("user_id=?", userId).Order("updated_at").Find(userWordProgress).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) GetUserWordProgressById(userWordProgress *model.UserWordProgress, userId uint, userWordProgressId uint) error {
	if err := uwpr.db.Where("user_id=?", userId).First(userWordProgress, userWordProgressId).Error; err != nil {
		return err
	}
	return nil
}

func (uwpr *userWordProgressRepository) CreateUserWordProgress(userWordProgress *model.UserWordProgress) error {
	if err := uwpr.db.Create(userWordProgress).Error; err != nil {
		return nil
	}
	return nil
}

func (uwpr *userWordProgressRepository) UpdateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, userWordProgressId uint) error {
	result := uwpr.db.Model(userWordProgress).Clauses(clause.Returning{}).Where("id=? AND user_id=?", userWordProgressId, userId).Update("total_typings", userWordProgress)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object dose not exist")
	}
	return nil
}
