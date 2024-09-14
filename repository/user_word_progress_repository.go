package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type IUserWordProgressRepository interface {
	GetAllUserWordProgress(userWordProgress *[]model.UserWordProgress, userId uint) error
	GetUserWordProgressByWordId(userWordProgress *model.UserWordProgress, userId uint, wordId uint) error
	CreateUserWordProgress(userWordProgress *model.UserWordProgress) error
	UpdateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, userWordProgressId uint) error
	FindOrCreateUserWordProgress(userId uint, wordId uint) (*model.UserWordProgress, error)
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

func (uwpr *userWordProgressRepository) UpdateUserWordProgress(userWordProgress *model.UserWordProgress, userId uint, userWordProgressId uint) error {
	existingUserWordProgress := model.UserWordProgress{}
	if err := uwpr.db.Where("id=? AND user_id=?", userWordProgressId, userId).First(&existingUserWordProgress).Error; err != nil {
		return err
	}

	existingUserWordProgress.TotalTypings += 1

	result := uwpr.db.Model(&existingUserWordProgress).Clauses(clause.Returning{}).Where("id=? AND user_id=?", userWordProgressId, userId).Updates(map[string]interface{}{
		"total_typings": existingUserWordProgress.TotalTypings,
		"typing_speed":  userWordProgress.TypingSpeed,
		"proficiency":   0.0,
	})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected < 1 {
		return fmt.Errorf("object dose not exist")
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
