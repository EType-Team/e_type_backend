package usecase

import (
	"api/model"
	"api/repository"
)

type ILessonUsecase interface {
	GetAllLesson() ([]model.LessonResponse, error)
	GetLessonById(lessonId uint) (model.LessonResponse, error)
}

type lessonUsecase struct {
	lr repository.ILessonRepository
}

func NewLessonUsecase(lr repository.ILessonRepository) ILessonUsecase {
	return &lessonUsecase{lr}
}

func (lu *lessonUsecase) GetAllLesson() ([]model.LessonResponse, error) {
	lessons := []model.Lesson{}
	if err := lu.lr.GetAllLesson(&lessons); err != nil {
		return nil, err
	}
	resLessons := []model.LessonResponse{}
	for _, v := range lessons {
		l := model.LessonResponse{
			ID:          v.ID,
			Title:       v.Title,
			Description: v.Description,
		}
		resLessons = append(resLessons, l)
	}
	return resLessons, nil
}

func (lu *lessonUsecase) GetLessonById(lessonId uint) (model.LessonResponse, error) {
	lesson := model.Lesson{}
	if err := lu.lr.GetLessonById(&lesson, lessonId); err != nil {
		return model.LessonResponse{}, err
	}
	resLesson := model.LessonResponse{
		ID:          lesson.ID,
		Title:       lesson.Title,
		Description: lesson.Description,
	}
	return resLesson, nil
}
