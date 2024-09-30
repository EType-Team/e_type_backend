package controller

import (
	"api/usecase"
	"api/model"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ILessonController interface {
	GetAllLesson(c echo.Context) error
	GetLessonById(c echo.Context) error
	CreateLesson(c echo.Context) error
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

func (lc *lessonController) CreateLesson(c echo.Context) error {
	var req model.CreateLessonRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	lesson := model.Lesson{
		Title:       req.Title,
		Description: req.Description,
	}

	words := make([]model.Word, len(req.Words))
	for i, w := range req.Words {
		words[i] = model.Word{
			English:  w.English,
			Japanese: w.Japanese,
			Mp3Path:  w.Mp3Path,
		}
	}

	err := lc.lu.CreateLesson(&lesson, words)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to register lesson"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Lesson registered successfully"})
}

