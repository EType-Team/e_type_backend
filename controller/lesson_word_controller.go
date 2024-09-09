package controller

import (
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ILessonWordController interface {
	GetLessonWordByLessonId(c echo.Context) error
}

type lessonWordController struct {
	lwu usecase.ILessonWordUsecase
}

func NewLessonWordController(lwu usecase.ILessonWordUsecase) ILessonWordController {
	return &lessonWordController{lwu}
}

func (lwc *lessonWordController) GetLessonWordByLessonId(c echo.Context) error {
	id := c.Param("lessonId")
	lessonId, _ := strconv.Atoi(id)
	lessonWordRes, err := lwc.lwu.GetLessonWordByLessonId(uint(lessonId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, lessonWordRes)
}
