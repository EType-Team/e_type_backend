package model

import "time"

type Word struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	English   string    `json:"english" gorm:"unique not null"`
	Japanese  string    `json:"japanese" gorm:"not null"`
	Mp3Path   string    `json:"mp3_path" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type WordResponse struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	English  string `json:"english" gorm:"not null"`
	Japanese string `json:"japanese" gorm:"not null"`
	Mp3Path  string `json:"mp3_path" gorm:"not null"`
}
