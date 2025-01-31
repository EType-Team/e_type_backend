package controller

import (
	"api/model"
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

type IUserWordProgressController interface {
	GetAllUserWordProgress(c echo.Context) error
	IncrementOrCreateUserWordProgress(c echo.Context) error
	GetUserWordProgressByWordId(c echo.Context) error
	GetUserWordProgressByLessonId(c echo.Context) error
	IncrementOrCreateUserWordTestProgress(c echo.Context) error
}

type userWordProgressController struct {
	uwpu usecase.IUserWordProgressUsecase
}

func NewUserWordProgressController(uwpu usecase.IUserWordProgressUsecase) IUserWordProgressController {
	return &userWordProgressController{uwpu}
}

func (uwpc *userWordProgressController) GetAllUserWordProgress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userWordProgressRes, err := uwpc.uwpu.GetAllUserWordProgress(uint(userId.(float64)))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userWordProgressRes)
}

func (uwpc *userWordProgressController) IncrementOrCreateUserWordProgress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userWordProgress := model.UserWordProgress{}
	if err := c.Bind(&userWordProgress); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userWordProgressRes, err := uwpc.uwpu.IncrementOrCreateUserWordProgress(userWordProgress, uint(userId.(float64)), userWordProgress.WordID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userWordProgressRes)
}

func (uwpc *userWordProgressController) GetUserWordProgressByWordId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	wordId, err := strconv.Atoi(c.Param("wordId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid word ID")
	}

	userWordProgress, err := uwpc.uwpu.GetUserWordProgressByWordId(uint(userId.(float64)), uint(wordId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userWordProgress)
}

func (uwpc *userWordProgressController) GetUserWordProgressByLessonId(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	lessonId, err := strconv.Atoi(c.Param("lessonId"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userWordProgress, err := uwpc.uwpu.GetUserWordProgressByLessonId(uint(userId.(float64)), uint(lessonId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, userWordProgress)
}

func (uwpc *userWordProgressController) IncrementOrCreateUserWordTestProgress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	var requestData struct {
		WordID    uint `json:"word_id"`
		IsCorrect bool `json:"is_correct"`
	}
	if err := c.Bind(&requestData); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	userWordProgressRes, err := uwpc.uwpu.IncrementOrCreateUserWordTestProgress(uint(userId.(float64)), requestData.WordID, requestData.IsCorrect)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, userWordProgressRes)
}
