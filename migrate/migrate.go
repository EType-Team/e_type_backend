package main

import (
	"api/db"
	"api/model"
	"api/seed"
	"fmt"
	"log"
)

func main() {
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

	if err := seed.SeedDatabase(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	if err := seed.SeedLesson2(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	if err := seed.SeedLesson3(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	if err := seed.SeedLesson4(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
}
