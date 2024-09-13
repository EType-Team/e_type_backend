package model

import "time"

type UserWordProgress struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	UserID       uint      `json:"user_id" gorm:"not null"`
	WordID       uint      `json:"word_id" gorm:"not null"`
	Word         Word      `json:"word" gorm:"foreignKey:WordID"`
	TotalTypings int       `json:"total_typings"`
	TypingSpeed  float64   `json:"typing_speed"`
	Proficiency  float64   `json:"proficiency"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

type UserWordProgressResponse struct {
	ID           uint    `json:"id" gorm:"primaryKey"`
	UserID       uint    `json:"user_id" gorm:"not null"`
	WordID       uint    `json:"word_id" gorm:"not null"`
	Word         Word    `json:"word"`
	TotalTypings int     `json:"total_typings"`
	TypingSpeed  float64 `json:"typing_speed"`
	Proficiency  float64 `json:"proficiency"`
}
