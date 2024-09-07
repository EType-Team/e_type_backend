package validator

import (
	"api/model"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type IUserWordProgressValidator interface {
	UserWordProgressValidate(userWordProgress model.UserWordProgress) error
}

type userWordProgressValidator struct{}

func NewTaskValidator() IUserWordProgressValidator {
	return &userWordProgressValidator{}
}

func (uwpv *userWordProgressValidator) UserWordProgressValidate(userWordProgress model.UserWordProgress) error {
	return validation.ValidateStruct(&userWordProgress,
		validation.Field(
			&userWordProgress.WordID,
			validation.Required.Error("単語が存在しません"),
		),
	)
}
