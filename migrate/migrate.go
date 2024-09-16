package main

import (
	"api/db"
	"api/model"
	"fmt"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbConn := db.NewDB()
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(
		&model.User{},
		&model.Provider{},
		&model.Lesson{},
		&model.Word{},
		&model.LessonWord{},
		&model.UserLessonProgress{},
		&model.UserWordProgress{},
	)
}
