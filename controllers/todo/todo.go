package todo

import "todolist/model"

type TodoRepository interface {
	Find(activityID interface{}) ([]model.Todo, error)
	FindByID(todoID int) (model.Todo, error)
	Create(todo model.Todo) (model.Todo, error)
	Update(todo model.Todo, todoID int) (model.Todo, error)
	Delete(todoID int) error
}

type TodoUsecase interface {
	GetAll(activityID interface{}) ([]model.GetTodoResponse, error)
	GetByID(todoID int) (model.GetTodoResponse, error)
	Create(todo model.Todo) (model.GetTodoResponse, error)
	Update(todo model.Todo, todoID int) (model.GetTodoResponse, error)
	Delete(todoID int) error
}
