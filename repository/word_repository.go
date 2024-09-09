package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type IWordRepository interface {
	GetWordById(word *model.Word, wordId uint) error
}

type wordRepository struct {
	db *gorm.DB
}

func NewWordRepository(db *gorm.DB) IWordRepository {
	return &wordRepository{db}
}

func (wr *wordRepository) GetWordById(word *model.Word, wordId uint) error {
	if err := wr.db.Where("id=?", wordId).First(word).Error; err != nil {
		return err
	}
	return nil
}
