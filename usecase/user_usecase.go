package usecase

import (
	"api/model"
	"api/repository"
	"api/validator"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type IUserUsecase interface {
	GetUserById(userId uint) (model.UserResponse, error)
	GetUserByEmail(user *model.User, email string) error
	CreateUser(user *model.User) error
	UpdateUser(user model.User, userId uint) (model.UserResponse, error)
	GenerateJWT(userID uint) (string, error)
}

type userUsecase struct {
	ur repository.IUserRepository
	uv validator.IUserValidator
}

func NewUserUsecase(ur repository.IUserRepository, uv validator.IUserValidator) IUserUsecase {
	return &userUsecase{ur, uv}
}

func (uu *userUsecase) GetUserById(userId uint) (model.UserResponse, error) {
	user := model.User{}
	if err := uu.ur.GetUserById(&user, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}
	return resUser, nil
}

func (uu *userUsecase) GetUserByEmail(user *model.User, email string) error {
	return uu.ur.GetUserByEmail(user, email)
}

func (uu *userUsecase) CreateUser(user *model.User) error {
	return uu.ur.CreateUser(user)
}

func (uu *userUsecase) UpdateUser(user model.User, userId uint) (model.UserResponse, error) {
	if err := uu.ur.UpdateUser(&user, userId); err != nil {
		return model.UserResponse{}, err
	}
	resUser := model.UserResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
		Image: user.Image,
	}
	return resUser, nil
}

func (uu *userUsecase) GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 12).Unix(),
	})
	return token.SignedString([]byte(os.Getenv("SECRET")))
}
