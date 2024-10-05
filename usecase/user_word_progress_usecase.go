package usecase

import (
	"api/model"
	"api/repository"
	"api/validator"
)

type IUserWordProgressUsecase interface {
	GetAllUserWordProgress(userId uint) ([]model.UserWordProgressResponse, error)
	IncrementOrCreateUserWordProgress(userWordProgress model.UserWordProgress, userId uint, wordId uint) (model.UserWordProgressResponse, error)
	GetUserWordProgressByWordId(userId uint, wordId uint) (model.UserWordProgressResponse, error)
	GetUserWordProgressByLessonId(userId uint, lessonId uint) ([]model.UserWordProgressResponse, error)
	IncrementOrCreateUserWordTestProgress(userId uint, wordId uint, isCorrect bool) (model.UserWordProgressResponse, error)
}

type userWordProgressUsecase struct {
	uwpr repository.IUserWordProgressRepository
	uwpv validator.IUserWordProgressValidator
	lwr  repository.ILessonWordRepository
}

func NewUserWordProgressUsecase(
	uwpr repository.IUserWordProgressRepository,
	uwpv validator.IUserWordProgressValidator,
	lwr repository.ILessonWordRepository,
) IUserWordProgressUsecase {
	return &userWordProgressUsecase{uwpr, uwpv, lwr}
}

func (uwpu *userWordProgressUsecase) GetAllUserWordProgress(userId uint) ([]model.UserWordProgressResponse, error) {
	userWordProgress := []model.UserWordProgress{}
	if err := uwpu.uwpr.GetAllUserWordProgress(&userWordProgress, userId); err != nil {
		return nil, err
	}
	resUserWordProgress := []model.UserWordProgressResponse{}
	for _, v := range userWordProgress {
		proficiency := model.CalculateProficiency(v.TotalTypings, v.CorrectTests, v.TotalTests, v.UpdatedAt)
		uwp := model.UserWordProgressResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			WordID:       v.WordID,
			Word:         v.Word,
			TotalTypings: int(v.TotalTypings),
			Proficiency:  proficiency,
		}
		resUserWordProgress = append(resUserWordProgress, uwp)
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) GetUserWordProgressByWordId(userId uint, wordId uint) (model.UserWordProgressResponse, error) {
	userWordProgress := model.UserWordProgress{}
	if err := uwpu.uwpr.GetUserWordProgressByWordId(&userWordProgress, userId, wordId); err != nil {
		return model.UserWordProgressResponse{}, err
	}
	proficiency := model.CalculateProficiency(userWordProgress.TotalTypings, userWordProgress.CorrectTests, userWordProgress.TotalTests, userWordProgress.UpdatedAt)
	resUserWordProgress := model.UserWordProgressResponse{
		ID:           userWordProgress.ID,
		UserID:       userWordProgress.UserID,
		WordID:       userWordProgress.WordID,
		Word:         userWordProgress.Word,
		TotalTypings: userWordProgress.TotalTypings,
		Proficiency:  proficiency,
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) IncrementOrCreateUserWordProgress(userWordProgress model.UserWordProgress, userId uint, wordId uint) (model.UserWordProgressResponse, error) {
	if err := uwpu.uwpr.CreateOrUpdateUserWordProgress(&userWordProgress, userId, wordId); err != nil {
		return model.UserWordProgressResponse{}, err
	}
	resUserWordProgress := model.UserWordProgressResponse{
		ID:           userWordProgress.ID,
		UserID:       userWordProgress.UserID,
		WordID:       userWordProgress.WordID,
		TotalTypings: userWordProgress.TotalTypings,
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) GetUserWordProgressByLessonId(userId uint, lessonId uint) ([]model.UserWordProgressResponse, error) {
	lessonWords := []model.LessonWord{}
	if err := uwpu.lwr.GetLessonWordByLessonId(&lessonWords, lessonId); err != nil {
		return nil, err
	}

	wordIds := []uint{}
	for _, lw := range lessonWords {
		wordIds = append(wordIds, lw.WordID)
	}

	userWordProgress := []model.UserWordProgress{}
	if err := uwpu.uwpr.GetUserWordProgressByWordIds(&userWordProgress, userId, wordIds); err != nil {
		return nil, err
	}

	resUserWordProgress := []model.UserWordProgressResponse{}
	for _, uwp := range userWordProgress {
		proficiency := model.CalculateProficiency(uwp.TotalTypings, uwp.CorrectTests, uwp.TotalTests, uwp.UpdatedAt)
		resUserWordProgress = append(resUserWordProgress, model.UserWordProgressResponse{
			ID:           uwp.ID,
			UserID:       uwp.UserID,
			WordID:       uwp.WordID,
			Word:         uwp.Word,
			LessonID:     lessonId,
			TotalTypings: uwp.TotalTypings,
			Proficiency:  proficiency,
		})
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) IncrementOrCreateUserWordTestProgress(userId uint, wordId uint, isCorrect bool) (model.UserWordProgressResponse, error) {
	userWordProgress := model.UserWordProgress{}
	if err := uwpu.uwpr.CreateOrUpdateUserWordTestProgress(&userWordProgress, userId, wordId, isCorrect); err != nil {
		return model.UserWordProgressResponse{}, err
	}
	resUserWordProgress := model.UserWordProgressResponse{
		ID:           userWordProgress.ID,
		UserID:       userWordProgress.UserID,
		WordID:       userWordProgress.WordID,
		TotalTypings: userWordProgress.TotalTypings,
	}
	return resUserWordProgress, nil
}
