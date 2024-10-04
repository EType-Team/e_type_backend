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

func CalculateProficiency(totalTypings int, updatedAt time.Time) float64 {
	// 現在時刻と更新時刻の差を計算
	timeElapsed := time.Since(updatedAt).Hours()

	// 忘却曲線の係数（回数に基づいて変わる）
	forgettingRate := 1.0 / float64(totalTypings)

	// エビングハウスの忘却曲線を基に熟練度を計算
	proficiency := math.Exp(-forgettingRate * timeElapsed)

	// 熟練度は 0.0 から 1.0 の範囲にする
	if proficiency < 0 {
		proficiency = 0
	}

	return proficiency * 100
}
