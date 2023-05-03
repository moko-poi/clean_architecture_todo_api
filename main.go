package main

import (
	"github.com/moko-poi/go-todo-api/controller"
	"github.com/moko-poi/go-todo-api/db"
	"github.com/moko-poi/go-todo-api/repository"
	"github.com/moko-poi/go-todo-api/router"
	"github.com/moko-poi/go-todo-api/usecase"
	"github.com/moko-poi/go-todo-api/validator"
)

func main() {
	db := db.NewDB()

	userValidator := validator.NewUserValidator()
	taskValidator := validator.NewTaskValidator()

	userRepository := repository.NewUserRepository(db)
	taskRepository := repository.NewTaskRepository(db)

	userUsecase := usecase.NewUserUsecase(userRepository, userValidator)
	taskUsecase := usecase.NewTaskUsecase(taskRepository, taskValidator)

	userController := controller.NewUserController(userUsecase)
	taskController := controller.NewTaskController(taskUsecase)

	e := router.NewRouter(userController, taskController)
	e.Logger.Fatal(e.Start(":8080"))
}
