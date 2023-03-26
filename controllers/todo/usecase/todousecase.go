package usecase

import (
	"fmt"

	"todolist/model"
	"todolist/controllers/todo"
)

type todoUsecase struct {
	todoRepository todo.TodoRepository
}

func NewTodoUsecase(todoRepository todo.TodoRepository) todo.TodoUsecase {
	return &todoUsecase{todoRepository}
}

func (u *todoUsecase) GetAll(activityID interface{}) ([]model.GetTodoResponse, error) {
	todos, err := u.todoRepository.Find(activityID)
	todosResponse := []model.GetTodoResponse{}
	if err != nil {
		return todosResponse, err
	}

	for _, todo := range todos {
		t := model.GetTodoResponse{
			ID:         todo.ID,
			ActivityID: todo.ActivityID,
			Title:      todo.Title,
			IsActive:   todo.IsActive,
			Priority:   todo.Priority,
			CreatedAt:  todo.CreatedAt,
			UpdatedAt:  todo.UpdatedAt,
		}

		todosResponse = append(todosResponse, t)
	}

	return todosResponse, nil
}

func (u *todoUsecase) GetByID(todoID int) (model.GetTodoResponse, error) {
	todo, err := u.todoRepository.FindByID(todoID)
	if err != nil {
		return model.GetTodoResponse{}, err
	}

	return model.GetTodoResponse{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

func (u *todoUsecase) Create(todo model.Todo) (model.GetTodoResponse, error) {
	if (todo == model.Todo{} || todo.Title == "") {
		return model.GetTodoResponse{}, fmt.Errorf("null struct")
	}

	if todo.ActivityID == 0 {
		return model.GetTodoResponse{}, fmt.Errorf("null activity id")
	}

	todo, err := u.todoRepository.Create(todo)
	if err != nil {
		return model.GetTodoResponse{}, err
	}

	return model.GetTodoResponse{
		ID:         todo.ID,
		ActivityID: todo.ActivityID,
		Title:      todo.Title,
		IsActive:   todo.IsActive,
		Priority:   todo.Priority,
		CreatedAt:  todo.CreatedAt,
		UpdatedAt:  todo.UpdatedAt,
	}, nil
}

func (u *todoUsecase) Update(todo model.Todo, todoID int) (model.GetTodoResponse, error) {
	foundTodo, err := u.todoRepository.FindByID(todoID)
	if err != nil {
		return model.GetTodoResponse{}, err
	}

	if todo.Title != "" {
		foundTodo.Title = todo.Title
	}

	if todo.Priority != "" {
		foundTodo.Priority = todo.Priority
	}

	if todo.IsActive != foundTodo.IsActive {
		foundTodo.IsActive = todo.IsActive
	}

	updatedTodo, err := u.todoRepository.Update(foundTodo, todoID)
	if err != nil {
		return model.GetTodoResponse{}, err
	}

	return model.GetTodoResponse{
		ID:         updatedTodo.ID,
		ActivityID: updatedTodo.ActivityID,
		Title:      updatedTodo.Title,
		IsActive:   updatedTodo.IsActive,
		Priority:   updatedTodo.Priority,
		CreatedAt:  updatedTodo.CreatedAt,
		UpdatedAt:  foundTodo.UpdatedAt,
	}, nil
}

func (u *todoUsecase) Delete(todoID int) error {
	if _, err := u.todoRepository.FindByID(todoID); err != nil {
		return err
	}

	return u.todoRepository.Delete(todoID)
}
