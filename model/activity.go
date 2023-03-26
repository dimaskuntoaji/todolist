package model

import (
	"time"
)

type Activity struct {
	ID        uint      `json:"id" gorm:"primaryKey;column:activity_id"`
	Title     string    `json:"title" gorm:"type:varchar(255)"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}