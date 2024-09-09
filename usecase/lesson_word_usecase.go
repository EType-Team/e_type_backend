package usecase

import (
	"api/model"
	"api/repository"
)

type ILessonWordUsecase interface {
	GetLessonWordByLessonId(lessonId uint) ([]model.LessonWordResponse, error)
}

type lessonWordUsecase struct {
	lwr repository.ILessonWordRepository
}

func NewLessonWordUsecase(lwr repository.ILessonWordRepository) ILessonWordUsecase {
	return &lessonWordUsecase{lwr}
}

func (lwu *lessonWordUsecase) GetLessonWordByLessonId(lessonId uint) ([]model.LessonWordResponse, error) {
	lessonWords := []model.LessonWord{}
	if err := lwu.lwr.GetLessonWordByLessonId(&lessonWords, lessonId); err != nil {
		return nil, err
	}
	resLessonWord := []model.LessonWordResponse{}
	for _, v := range lessonWords {
		lw := model.LessonWordResponse{
			ID:       v.ID,
			LessonID: v.LessonID,
			WordID:   v.WordID,
			Word:     v.Word,
		}
		resLessonWord = append(resLessonWord, lw)
	}
	return resLessonWord, nil
}
