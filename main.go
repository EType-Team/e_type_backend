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

	userWordProgressValidator := validator.NewUserWordProgressValidator()
	userWordProgressRepository := repository.NewUserWordProgressRepository(db)
	userWordProgressUsecase := usecase.NewUserWordProgressUsecase(userWordProgressRepository, userWordProgressValidator)
	userWordProgressController := controller.NewUserWordProgressController(userWordProgressUsecase)

	e := router.NewRouter(userController, userWordProgressController)
	e.Logger.Fatal(e.Start(":8080"))
}
