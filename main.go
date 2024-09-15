package main

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/router"
	"api/usecase"
	"api/validator"
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	oauthConfig := &oauth2.Config{
		RedirectURL:  os.Getenv("GOOGLE_REDIRECT_URI"),
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
			"openid",
		},
		Endpoint: google.Endpoint,
	}

	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase, oauthConfig)

	lessonRepository := repository.NewLessonRepository(db)
	lessonUsecase := usecase.NewLessonUsecase(lessonRepository)
	lessonController := controller.NewLessonController(lessonUsecase)

	lessonWordRepository := repository.NewLessonWordRepository(db)
	lessonWordUsecase := usecase.NewLessonWordUsecase(lessonWordRepository)
	lessonWordController := controller.NewLessonWordController(lessonWordUsecase)

	userWordProgressValidator := validator.NewUserWordProgressValidator()
	userWordProgressRepository := repository.NewUserWordProgressRepository(db)
	userWordProgressUsecase := usecase.NewUserWordProgressUsecase(userWordProgressRepository, userWordProgressValidator)
	userWordProgressController := controller.NewUserWordProgressController(userWordProgressUsecase)

	e := router.NewRouter(
		userController,
		lessonController,
		lessonWordController,
		userWordProgressController,
	)
	e.Logger.Fatal(e.Start(":8080"))
}
