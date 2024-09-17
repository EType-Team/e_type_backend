package main

import (
	"api/db"
	"api/model"
	"fmt"
	"log"

	"gorm.io/gorm"
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

	if err := SeedDatabase(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
}

func SeedDatabase(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Count(&count)
	if count > 0 {
		fmt.Println("Seed data already exists, skipping seeding")
		return nil
	}

	words := []model.Word{
		{English: "achievement", Japanese: "達成", Mp3Path: "/audio/achievement.mp3"},
		{English: "allocate", Japanese: "割り当てる", Mp3Path: "/audio/allocate.mp3"},
		{English: "applicant", Japanese: "応募者", Mp3Path: "/audio/applicant.mp3"},
		{English: "appointment", Japanese: "予約", Mp3Path: "/audio/appointment.mp3"},
		{English: "assess", Japanese: "評価する", Mp3Path: "/audio/assess.mp3"},
		{English: "attend", Japanese: "出席する", Mp3Path: "/audio/attend.mp3"},
		{English: "budget", Japanese: "予算", Mp3Path: "/audio/budget.mp3"},
		{English: "candidate", Japanese: "候補者", Mp3Path: "/audio/candidate.mp3"},
		{English: "conference", Japanese: "会議", Mp3Path: "/audio/conference.mp3"},
		{English: "confirm", Japanese: "確認する", Mp3Path: "/audio/confirm.mp3"},
		{English: "consult", Japanese: "相談する", Mp3Path: "/audio/consult.mp3"},
		{English: "contribute", Japanese: "貢献する", Mp3Path: "/audio/contribute.mp3"},
		{English: "deliver", Japanese: "配達する", Mp3Path: "/audio/deliver.mp3"},
		{English: "department", Japanese: "部門", Mp3Path: "/audio/department.mp3"},
		{English: "efficient", Japanese: "効率的な", Mp3Path: "/audio/efficient.mp3"},
		{English: "estimate", Japanese: "見積もる", Mp3Path: "/audio/estimate.mp3"},
		{English: "expand", Japanese: "拡大する", Mp3Path: "/audio/expand.mp3"},
		{English: "facility", Japanese: "施設", Mp3Path: "/audio/facility.mp3"},
		{English: "frequent", Japanese: "頻繁な", Mp3Path: "/audio/frequent.mp3"},
		{English: "implement", Japanese: "実行する", Mp3Path: "/audio/implement.mp3"},
		{English: "income", Japanese: "収入", Mp3Path: "/audio/income.mp3"},
		{English: "invoice", Japanese: "請求書", Mp3Path: "/audio/invoice.mp3"},
		{English: "manage", Japanese: "管理する", Mp3Path: "/audio/manage.mp3"},
		{English: "negotiation", Japanese: "交渉", Mp3Path: "/audio/negotiation.mp3"},
		{English: "offer", Japanese: "提供する", Mp3Path: "/audio/offer.mp3"},
		{English: "participate", Japanese: "参加する", Mp3Path: "/audio/participate.mp3"},
		{English: "policy", Japanese: "方針", Mp3Path: "/audio/policy.mp3"},
		{English: "procedure", Japanese: "手続き", Mp3Path: "/audio/procedure.mp3"},
		{English: "schedule", Japanese: "予定", Mp3Path: "/audio/schedule.mp3"},
		{English: "submit", Japanese: "提出する", Mp3Path: "/audio/submit.mp3"},
	}

	if err := db.Create(&words).Error; err != nil {
		return err
	}

	lesson := model.Lesson{
		Title:       "TOEICで出る重要単語30選",
		Description: "このレッスンでは、TOEICの試験で頻出する30個の重要単語を学びます。ビジネスシーンや日常生活で役立つ単語を、効率よく覚えることができます。各単語には音声も用意されており、発音を確認しながら単語を身につけましょう。",
	}

	if err := db.Create(&lesson).Error; err != nil {
		return err
	}

	var lessonWords []model.LessonWord
	for _, word := range words {
		lw := model.LessonWord{
			LessonID: lesson.ID,
			WordID:   word.ID,
		}
		lessonWords = append(lessonWords, lw)
	}

	if err := db.Create(&lessonWords).Error; err != nil {
		return err
	}

	fmt.Println("Seed data inserted successfully")
	return nil
}
