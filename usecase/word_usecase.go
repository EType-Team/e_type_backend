package usecase

import (
	"api/model"
	"api/repository"
)

type IWordUsecase interface {
	GetWordById(wordId uint) (model.WordResponse, error)
}

type wordUsecase struct {
	wr repository.IWordRepository
}

func NewWordUsecase(wr repository.IWordRepository) IWordUsecase {
	return &wordUsecase{wr}
}

func (wu *wordUsecase) GetWordById(wordId uint) (model.WordResponse, error) {
	word := model.Word{}
	if err := wu.wr.GetWordById(&word, wordId); err != nil {
		return model.WordResponse{}, err
	}
	resWord := model.WordResponse{
		ID:       word.ID,
		English:  word.English,
		Japanese: word.Japanese,
		Mp3Path:  word.Mp3Path,
	}
	return resWord, nil
}
