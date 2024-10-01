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
	UpdateLesson(c echo.Context) error
	DeleteLesson(c echo.Context) error
	AddNewWordToLesson(c echo.Context) error
	RemoveWordFromLesson(c echo.Context) error
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
	var req model.LessonRequest
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


func (lc *lessonController) UpdateLesson(c echo.Context) error {
	id := c.Param("lessonId")
	lessonId, _ := strconv.Atoi(id)

	var req model.LessonUpdateRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
	}

	err := lc.lu.UpdateLesson(uint(lessonId), req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to update lesson"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Lesson updated successfully"})
}

func (lc *lessonController) DeleteLesson(c echo.Context) error {
	id := c.Param("lessonId")
	lessonId, _ := strconv.Atoi(id)

	err := lc.lu.DeleteLesson(uint(lessonId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to delete lesson"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Lesson deleted successfully"})
}

func (lc *lessonController) AddNewWordToLesson(c echo.Context) error {
    lessonId, _ := strconv.Atoi(c.Param("lessonId"))
    
    var word model.Word
    if err := c.Bind(&word); err != nil {
        return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid input"})
    }

    err := lc.lu.AddNewWordToLesson(uint(lessonId), word)
    if err != nil {
        return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to add new word to lesson"})
    }

    return c.JSON(http.StatusOK, echo.Map{"message": "New word added to lesson successfully"})
}

func (lc *lessonController) RemoveWordFromLesson(c echo.Context) error {
	lessonId, _ := strconv.Atoi(c.Param("lessonId"))
	wordId, _ := strconv.Atoi(c.Param("wordId"))

	err := lc.lu.RemoveWordFromLesson(uint(lessonId), uint(wordId))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to remove word from lesson"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "Word removed from lesson successfully"})
}