package model

import (
	"math"
	"time"
)

type UserWordProgress struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	WordID       uint      `json:"word_id" gorm:"not null"`
	Word         Word      `json:"word" gorm:"foreignKey:WordID"`
	TotalTypings int       `json:"total_typings"`
	CorrectTests int       `json:"correct_tests"`
	TotalTests   int       `json:"total_tests"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserWordProgressResponse struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	UserID       uint    `json:"user_id" gorm:"not null"`
	WordID       uint    `json:"word_id" gorm:"not null"`
	Word         Word    `json:"word"`
	LessonID     uint    `json:"lessonId"`
	TotalTypings int     `json:"total_typings"`
	Proficiency  float64 `json:"proficiency"`
}

func CalculateProficiency(totalTypings int, correctTests int, totalTests int, updatedAt time.Time) float64 {
	if totalTypings == 0 {
		return 0.0
	}

	timeElapsed := time.Since(updatedAt).Hours()

	var testAccuracy float64
	if totalTests > 0 {
		testAccuracy = float64(correctTests) / float64(totalTests)
	} else {
		testAccuracy = 0
	}

	forgettingRate := 1.0 / (float64(totalTypings) + (testAccuracy * float64(totalTests)))

	proficiency := math.Exp(-forgettingRate * timeElapsed)

	if proficiency < 0 {
		proficiency = 0
	}

	return proficiency * 100
}
