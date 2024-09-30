package usecase

import (
	"api/model"
	"api/repository"
)

type ILessonUsecase interface {
	GetAllLesson() ([]model.LessonResponse, error)
	GetLessonById(lessonId uint) (model.LessonResponse, error)
	CreateLesson(lesson *model.Lesson, words []model.Word) error
}

type lessonUsecase struct {
	lr repository.ILessonRepository
	wr repository.IWordRepository
}

func NewLessonUsecase(lr repository.ILessonRepository, wr repository.IWordRepository) ILessonUsecase {
	return &lessonUsecase{lr, wr}
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


func (lu *lessonUsecase) CreateLesson(lesson *model.Lesson, words []model.Word) error {
    if err := lu.lr.CreateLesson(lesson); err != nil {
        return err
    }

    for _, word := range words {
        if err := lu.wr.CreateWord(&word); err != nil {
            return err
        }
        lessonWord := model.LessonWord{
            LessonID: lesson.ID,
            WordID:   word.ID,
        }

        if err := lu.lr.CreateLessonWord(&lessonWord); err != nil {
            return err
        }
    }

    return nil
}
