package repository

import (
	"api/model"

	"gorm.io/gorm"
)

type ILessonRepository interface {
	GetAllLesson(lesson *[]model.Lesson) error
	GetLessonById(lesson *model.Lesson, lessonId uint) error
}

type lessonRepository struct {
	db *gorm.DB
}

func NewLessonRepository(db *gorm.DB) ILessonRepository {
	return &lessonRepository{db}
}

func (lr *lessonRepository) GetAllLesson(lesson *[]model.Lesson) error {
	if err := lr.db.Find(lesson).Error; err != nil {
		return err
	}
	return nil
}

func (lr *lessonRepository) GetLessonById(lesson *model.Lesson, lessonId uint) error {
	if err := lr.db.Where("id=?", lessonId).First(lesson).Error; err != nil {
		return err
	}
	return nil
}
