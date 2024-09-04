package model

import "time"

type UserLessonProgress struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	UserID     uint      `json:"user_id" gorm:"not null"`
	LessonID   uint      `json:"lesson_id" gorm:"not null"`
	Completion float64   `json:"completion"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type UserLessonProgressResponse struct {
	ID         uint    `json:"id" gorm:"primaryKey"`
	UserID     uint    `json:"user_id" gorm:"not null"`
	LessonID   uint    `json:"lesson_id" gorm:"not null"`
	Completion float64 `json:"completion"`
}
