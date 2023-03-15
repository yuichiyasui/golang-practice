package models

import (
	"time"

	"gorm.io/gorm"
)

type Status string

const (
	todo       Status = "todo"
	inProgress Status = "in_progress"
	done       Status = "done"
)

type Task struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Status    Status         `json:"status" sql:"type:status" gorm:"default: 'todo'; not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
