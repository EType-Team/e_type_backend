package main

import (
	"api/controller"
	"api/db"
	"api/repository"
	"api/router"
	"api/usecase"
	"api/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	userRepository := repository.NewUserRepository(db)
	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	userController := controller.NewUserController(userUsecase)

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
