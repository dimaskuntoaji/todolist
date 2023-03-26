package model

import (
	"time"
)

type Todo struct {
	ID         uint      `json:"id" gorm:"primaryKey;column:todo_id"`
	ActivityID uint      `json:"activity_group_id" gorm:"column:activity_group_id"`
	Title      string    `json:"title" gorm:"type:varchar(255)"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority" gorm:"type:varchar(255)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Activity   Activity  `json:"activity" gorm:"foreignkey:ActivityID"`
}

type GetTodoResponse struct {
	ID         uint      `json:"id"`
	ActivityID uint      `json:"activity_group_id"`
	Title      string    `json:"title"`
	IsActive   bool      `json:"is_active"`
	Priority   string    `json:"priority"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}