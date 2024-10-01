package usecase

import (
	"api/model"
	"api/repository"
)

type ILessonUsecase interface {
	GetAllLesson() ([]model.LessonResponse, error)
	GetLessonById(lessonId uint) (model.LessonResponse, error)
	CreateLesson(lesson *model.Lesson, words []model.Word) error
	UpdateLesson(lessonId uint, updateReq model.LessonUpdateRequest) error
	DeleteLesson(lessonId uint) error
	AddNewWordToLesson(lessonId uint, word model.Word) error
	RemoveWordFromLesson(lessonId uint, wordId uint) error
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


func (lu *lessonUsecase) UpdateLesson(lessonId uint, updateReq model.LessonUpdateRequest) error {
	lesson := &model.Lesson{ID: lessonId, Title: updateReq.Title, Description: updateReq.Description}
	return lu.lr.UpdateLesson(lesson)
}

func (lu *lessonUsecase) DeleteLesson(lessonId uint) error {
	return lu.lr.DeleteLesson(lessonId)
}

func (lu *lessonUsecase) AddNewWordToLesson(lessonId uint, word model.Word) error {
    // 新しい単語を作成
    if err := lu.wr.CreateWord(&word); err != nil {
        return err
    }
    
    // 作成した単語をレッスンに関連付け
    return lu.lr.AddWordToLesson(lessonId, word.ID)
}

func (lu *lessonUsecase) RemoveWordFromLesson(lessonId uint, wordId uint) error {
	return lu.lr.RemoveWordFromLesson(lessonId, wordId)
}