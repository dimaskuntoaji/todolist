package handler

import (
	"github.com/gofiber/fiber/v2"
	activityHttp "todolist/controllers/activity/http"
	activityRepository "todolist/controllers/activity/repositoryactivity"
	activityUsecase "todolist/controllers/activity/usecase"
	todoHttp "todolist/controllers/todo/http"
	todoRepository "todolist/controllers/todo/repositorytodo"
	todoUsecase "todolist/controllers/todo/usecase"
	"gorm.io/gorm"
)

func MapHandler(app *fiber.App, db *gorm.DB) {
	activityRepo := activityRepository.NewActivityRepository(db)
	todoRepo := todoRepository.NewMysqlTodoRepository(db)

	activityUcase := activityUsecase.NewActivityUsecase(activityRepo)
	todoUcase := todoUsecase.NewTodoUsecase(todoRepo)

	activityHandler := activityHttp.NewActivityHandler(activityUcase)
	todoHandler := todoHttp.NewTodoHandler(todoUcase)

	activityHttp.MapRoutes(app, *activityHandler)
	todoHttp.MapRoutes(app, *todoHandler)
}
