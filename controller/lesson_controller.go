package controller

import (
	"api/usecase"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ILessonController interface {
	GetAllLesson(c echo.Context) error
	GetLessonById(c echo.Context) error
}

type lessonController struct {
	lu usecase.ILessonUsecase
}

func NewLessonController(lu usecase.ILessonUsecase) ILessonController {
	return &lessonController{lu}
}

func (lc *lessonController) GetAllLesson(c echo.Context) error {
	lessonsRes, err := lc.lu.GetAllLesson()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, lessonsRes)
}

func (lc *lessonController) GetLessonById(c echo.Context) error {
	id := c.Param("lessonId")
	lessonId, _ := strconv.Atoi(id)

	lessonRes, err := lc.lu.GetLessonById(uint(lessonId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, lessonRes)
}
