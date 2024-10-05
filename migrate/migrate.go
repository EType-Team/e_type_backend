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

	if err := DropColumns(dbConn); err != nil {
		log.Fatalf("Failed to drop columns: %v", err)
	}

	if err := SeedDatabase(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}

	if err := SeedLesson2(dbConn); err != nil {
		log.Fatalf("Failed to seed database: %v", err)
	}
}

func DropColumns(db *gorm.DB) error {
	err := db.Migrator().DropColumn(&model.UserWordProgress{}, "typing_speed")
	if err != nil {
		return fmt.Errorf("failed to drop column 'typing_speed' from 'user_word_progress' table: %w", err)
	}

	err = db.Migrator().DropColumn(&model.UserWordProgress{}, "proficiency")
	if err != nil {
		return fmt.Errorf("failed to drop column 'proficiency' from 'user_word_progress' table: %w", err)
	}

	fmt.Println("Columns dropped successfully")
	return nil
}

func SeedDatabase(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Count(&count)
	if count > 0 {
		fmt.Println("Seed data already exists, skipping seeding")
		return nil
	}

	words1 := []model.Word{
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

	if err := db.Create(&words1).Error; err != nil {
		return err
	}

	lesson1 := model.Lesson{
		Title:       "TOEICで出る重要単語30選",
		Description: "このレッスンでは、TOEICの試験で頻出する30個の重要単語を学びます。ビジネスシーンや日常生活で役立つ単語を、効率よく覚えることができます。各単語には音声も用意されており、発音を確認しながら単語を身につけましょう。",
	}

	if err := db.Create(&lesson1).Error; err != nil {
		return err
	}

	var lessonWords1 []model.LessonWord
	for _, word := range words1 {
		lw := model.LessonWord{
			LessonID: lesson1.ID,
			WordID:   word.ID,
		}
		lessonWords1 = append(lessonWords1, lw)
	}

	if err := db.Create(&lessonWords1).Error; err != nil {
		return err
	}

	fmt.Println("Seed data inserted successfully")
	return nil
}

func SeedLesson2(db *gorm.DB) error {
	var count int64
	db.Model(&model.Word{}).Where("english IN (?)", []string{"stand", "sit", "hold", "look at", "walk"}).Count(&count)
	if count > 0 {
		fmt.Println("Lesson 2 seed data already exists, skipping seeding")
		return nil
	}

	words2 := []model.Word{
		{English: "stand", Japanese: "立っている", Mp3Path: "/audio/stand.mp3"},
		{English: "sit", Japanese: "座っている", Mp3Path: "/audio/sit.mp3"},
		{English: "hold", Japanese: "持っている", Mp3Path: "/audio/hold.mp3"},
		{English: "look at", Japanese: "見ている", Mp3Path: "/audio/look_at.mp3"},
		{English: "walk", Japanese: "歩いている", Mp3Path: "/audio/walk.mp3"},
		{English: "carry", Japanese: "運んでいる", Mp3Path: "/audio/carry.mp3"},
		{English: "write", Japanese: "書いている", Mp3Path: "/audio/write.mp3"},
		{English: "read", Japanese: "読んでいる", Mp3Path: "/audio/read.mp3"},
		{English: "point", Japanese: "指を指している", Mp3Path: "/audio/point.mp3"},
		{English: "open", Japanese: "開ける", Mp3Path: "/audio/open.mp3"},
		{English: "close", Japanese: "閉じる", Mp3Path: "/audio/close.mp3"},
		{English: "wear", Japanese: "着ている", Mp3Path: "/audio/wear.mp3"},
		{English: "put on", Japanese: "着るところ", Mp3Path: "/audio/put_on.mp3"},
		{English: "take off", Japanese: "脱ぐ", Mp3Path: "/audio/take_off.mp3"},
		{English: "talk", Japanese: "話している", Mp3Path: "/audio/talk.mp3"},
		{English: "listen to", Japanese: "聞いている", Mp3Path: "/audio/listen_to.mp3"},
		{English: "use", Japanese: "使っている", Mp3Path: "/audio/use.mp3"},
		{English: "clean", Japanese: "掃除する", Mp3Path: "/audio/clean.mp3"},
		{English: "cook", Japanese: "料理する", Mp3Path: "/audio/cook.mp3"},
		{English: "serve", Japanese: "提供する", Mp3Path: "/audio/serve.mp3"},
		{English: "push", Japanese: "押す", Mp3Path: "/audio/push.mp3"},
		{English: "pull", Japanese: "引く", Mp3Path: "/audio/pull.mp3"},
		{English: "fill", Japanese: "満たす", Mp3Path: "/audio/fill.mp3"},
		{English: "empty", Japanese: "空にする", Mp3Path: "/audio/empty.mp3"},
		{English: "lift", Japanese: "持ち上げる", Mp3Path: "/audio/lift.mp3"},
		{English: "lower", Japanese: "下げる", Mp3Path: "/audio/lower.mp3"},
		{English: "turn on", Japanese: "つける", Mp3Path: "/audio/turn_on.mp3"},
		{English: "turn off", Japanese: "消す", Mp3Path: "/audio/turn_off.mp3"},
		{English: "repair", Japanese: "修理する", Mp3Path: "/audio/repair.mp3"},
		{English: "paint", Japanese: "塗る", Mp3Path: "/audio/paint.mp3"},
	}

	if err := db.Create(&words2).Error; err != nil {
		return err
	}

	lesson2 := model.Lesson{
		Title:       "TOEIC Part 1 で出る重要動詞30選",
		Description: "このレッスンでは、TOEIC Part 1（写真描写問題）で頻出する30個の重要な動詞を学びます。日常生活やビジネスの場面で使われる動作を表す単語を、写真をイメージしながら効率よく覚えることができます。各単語には発音を確認できる音声も付いており、リスニングと発音の練習にも最適です。",
	}

	if err := db.Create(&lesson2).Error; err != nil {
		return err
	}

	var lessonWords2 []model.LessonWord
	for _, word := range words2 {
		lw := model.LessonWord{
			LessonID: lesson2.ID,
			WordID:   word.ID,
		}
		lessonWords2 = append(lessonWords2, lw)
	}

	if err := db.Create(&lessonWords2).Error; err != nil {
		return err
	}

	fmt.Println("Lesson 2 seed data inserted successfully")
	return nil
}
