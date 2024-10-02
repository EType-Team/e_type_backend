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
}

type userWordProgressUsecase struct {
	uwpr repository.IUserWordProgressRepository
	uwpv validator.IUserWordProgressValidator
}

func NewUserWordProgressUsecase(
	uwpr repository.IUserWordProgressRepository,
	uwpv validator.IUserWordProgressValidator,
) IUserWordProgressUsecase {
	return &userWordProgressUsecase{uwpr, uwpv}
}

func (uwpu *userWordProgressUsecase) GetAllUserWordProgress(userId uint) ([]model.UserWordProgressResponse, error) {
	userWordProgress := []model.UserWordProgress{}
	if err := uwpu.uwpr.GetAllUserWordProgress(&userWordProgress, userId); err != nil {
		return nil, err
	}
	resUserWordProgress := []model.UserWordProgressResponse{}
	for _, v := range userWordProgress {
		proficiency := model.CalculateProficiency(v.TotalTypings, v.UpdatedAt)
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
	proficiency := model.CalculateProficiency(userWordProgress.TotalTypings, userWordProgress.UpdatedAt)
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
