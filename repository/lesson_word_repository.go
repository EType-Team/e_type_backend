package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type ILessonWordRepository interface {
	GetLessonWordByLessonId(lessonWord *[]model.LessonWord, lessonId uint) error
}

type lessonWordRepository struct {
	db *gorm.DB
}

func NewLessonWordRepository(db *gorm.DB) ILessonWordRepository {
	return &lessonWordRepository{db}
}

func (lwr *lessonWordRepository) GetLessonWordByLessonId(lessonWord *[]model.LessonWord, lessonId uint) error {
	if err := lwr.db.Where("lesson_id=?", lessonId).Preload("Word").Find(lessonWord).Error; err != nil {
		return err
	}
	return nil
}
