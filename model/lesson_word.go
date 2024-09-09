package model

import "time"

type LessonWord struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	LessonID  uint      `json:"lesson_id" gorm:"not null"`
	WordID    uint      `json:"word_id" gorm:"not null"`
	Word      Word      `json:"word" gorm:"foreignKey:WordID"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type LessonWordResponse struct {
	ID       uint `json:"id" gorm:"primaryKey"`
	LessonID uint `json:"lesson_id"`
	WordID   uint `json:"word_id"`
	Word     Word `json:"word"`
}
