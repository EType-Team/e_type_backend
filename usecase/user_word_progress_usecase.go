package usecase

import (
	"api/model"
	"api/repository"
	"api/validator"
)

type IUserWordProgressUsecase interface {
	GetAllUserWordProgress(userId uint) ([]model.UserWordProgressResponse, error)
	GetUserWordProgressById(userId uint, userWordProgressId uint) (model.UserWordProgressResponse, error)
	CreateUserWordProgress(userWordProgress model.UserWordProgress) (model.UserWordProgressResponse, error)
	UpdateUserWordProgress(userWordProgress model.UserWordProgress, userId uint, userWordProgressId uint) (model.UserWordProgressResponse, error)
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
		uwp := model.UserWordProgressResponse{
			ID:           v.ID,
			UserID:       v.UserID,
			WordID:       v.WordID,
			TotalTypings: int(v.TotalTypings),
			TypingSpeed:  v.TypingSpeed,
			Proficiency:  v.Proficiency,
		}
		resUserWordProgress = append(resUserWordProgress, uwp)
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) GetUserWordProgressById(userId uint, userWordProgressId uint) (model.UserWordProgressResponse, error) {
	userWordProgress := model.UserWordProgress{}
	if err := uwpu.uwpr.GetUserWordProgressById(&userWordProgress, userId, userWordProgressId); err != nil {
		return model.UserWordProgressResponse{}, err
	}
	resUserWordProgress := model.UserWordProgressResponse{
		ID:           userWordProgress.ID,
		UserID:       userWordProgress.UserID,
		WordID:       userWordProgress.WordID,
		TotalTypings: userWordProgress.TotalTypings,
		TypingSpeed:  userWordProgress.TypingSpeed,
		Proficiency:  userWordProgress.Proficiency,
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) CreateUserWordProgress(userWordProgress model.UserWordProgress) (model.UserWordProgressResponse, error) {
	if err := uwpu.uwpv.UserWordProgressValidate(userWordProgress); err != nil {
		return model.UserWordProgressResponse{}, err
	}

	if err := uwpu.uwpr.CreateUserWordProgress(&userWordProgress); err != nil {
		return model.UserWordProgressResponse{}, err
	}
	resUserWordProgress := model.UserWordProgressResponse{
		ID:           userWordProgress.ID,
		UserID:       userWordProgress.UserID,
		WordID:       userWordProgress.WordID,
		TotalTypings: userWordProgress.TotalTypings,
		TypingSpeed:  userWordProgress.TypingSpeed,
		Proficiency:  userWordProgress.Proficiency,
	}
	return resUserWordProgress, nil
}

func (uwpu *userWordProgressUsecase) UpdateUserWordProgress(userWordProgress model.UserWordProgress, userId uint, userWordProgressId uint) (model.UserWordProgressResponse, error) {
	if err := uwpu.uwpv.UserWordProgressValidate(userWordProgress); err != nil {
		return model.UserWordProgressResponse{}, err
	}

	if err := uwpu.uwpr.UpdateUserWordProgress(&userWordProgress, userId, userWordProgressId); err != nil {
		return model.UserWordProgressResponse{}, err
	}

	updateUserWordProgress := model.UserWordProgress{}
	if err := uwpu.uwpr.GetUserWordProgressById(&updateUserWordProgress, userId, userWordProgressId); err != nil {
		return model.UserWordProgressResponse{}, err
	}

	resUserWordProgress := model.UserWordProgressResponse{
		ID:           updateUserWordProgress.ID,
		UserID:       updateUserWordProgress.UserID,
		WordID:       updateUserWordProgress.WordID,
		TotalTypings: updateUserWordProgress.TotalTypings,
		TypingSpeed:  updateUserWordProgress.TypingSpeed,
		Proficiency:  updateUserWordProgress.Proficiency,
	}
	return resUserWordProgress, nil
}
