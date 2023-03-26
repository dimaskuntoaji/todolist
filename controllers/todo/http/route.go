package http

import "github.com/gofiber/fiber/v2"

func MapRoutes(r fiber.Router, h todoHandler) {
	r.Get("/todo-items", h.GetAllTodos())
	r.Get("/todo-items/:id", h.GetTodoByID())
	r.Post("/todo-items", h.CreateTodo())
	r.Patch("/todo-items/:id", h.UpdateTodo())
	r.Delete("/todo-items/:id", h.DeleteTodo())
}
