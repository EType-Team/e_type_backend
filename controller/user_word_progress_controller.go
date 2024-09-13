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
	GetUserWordProgressById(c echo.Context) error
	CreateUserWordProgress(c echo.Context) error
	UpdateUserWordProgress(c echo.Context) error
	IncrementUserWordProgress(c echo.Context) error
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

func (uwpc *userWordProgressController) GetUserWordProgressById(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("userWordProgressId")
	userWordProgressId, _ := strconv.Atoi(id)

	userWordProgressRes, err := uwpc.uwpu.GetUserWordProgressById(uint(userId.(float64)), uint(userWordProgressId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userWordProgressRes)
}

func (uwpc *userWordProgressController) CreateUserWordProgress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userWordProgress := model.UserWordProgress{}
	if err := c.Bind(&userWordProgress); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	userWordProgress.UserID = uint(userId.(float64))
	userWordProgressRes, err := uwpc.uwpu.CreateUserWordProgress(userWordProgress)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userWordProgressRes)
}

func (uwpc *userWordProgressController) UpdateUserWordProgress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]
	id := c.Param("userWordProgressId")
	userWordProgressId, _ := strconv.Atoi(id)

	userWordProgress := model.UserWordProgress{}
	if err := c.Bind(&userWordProgress); err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}

	userWordProgressRes, err := uwpc.uwpu.UpdateUserWordProgress(userWordProgress, uint(userId.(float64)), uint(userWordProgressId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, userWordProgressRes)
}

func (uwpc *userWordProgressController) IncrementUserWordProgress(c echo.Context) error {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	userId := claims["user_id"]

	userWordProgress := model.UserWordProgress{}
	if err := c.Bind(&userWordProgress); err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	userWordProgressRes, err := uwpc.uwpu.IncrementTotalTypings(uint(userId.(float64)), userWordProgress.WordID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, userWordProgressRes)
}
