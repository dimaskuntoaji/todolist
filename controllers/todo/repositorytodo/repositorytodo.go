package repositorytodo

import (
	"todolist/model"
	"todolist/controllers/todo"
	"gorm.io/gorm"
)

type mysqlTodoRepository struct {
	db *gorm.DB
}

func NewMysqlTodoRepository(db *gorm.DB) todo.TodoRepository {
	return &mysqlTodoRepository{db}
}

func (r *mysqlTodoRepository) Find(activityID interface{}) ([]model.Todo, error) {
	todos := []model.Todo{}

	tx := r.db.Preload("Activity")
	if _, ok := activityID.(int); ok {
		tx.Where("activity_group_id = ?", activityID)
	}
	tx.Find(&todos)
	return todos, tx.Error
}

func (r *mysqlTodoRepository) FindByID(todoID int) (model.Todo, error) {
	todo := model.Todo{}
	tx := r.db.Preload("Activity").Where("todo_id = ?", todoID).First(&todo)
	return todo, tx.Error
}

func (r *mysqlTodoRepository) Create(todo model.Todo) (model.Todo, error) {
	tx := r.db.Create(&todo)
	return todo, tx.Error
}

func (r *mysqlTodoRepository) Update(todo model.Todo, todoID int) (model.Todo, error) {
	tx := r.db.Where("todo_id = ?", todoID).Updates(&todo)
	return todo, tx.Error
}

func (r *mysqlTodoRepository) Delete(todoID int) error {
	return r.db.Where("todo_id = ?", todoID).Delete(&model.Todo{}).Error
}
