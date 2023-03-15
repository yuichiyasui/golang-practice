package models

import (
	"time"

	"gorm.io/gorm"
)

type status string

const (
	todo       status = "todo"
	inProgress status = "in_progress"
	done       status = "done"
)

type Task struct {
	ID        uint           `json:"id" gorm:"primaryKey"`
	Title     string         `json:"title" gorm:"not null"`
	Status    status         `json:"status" sql:"type:status; default: 'todo'; not null"`
	CreatedAt time.Time      `json:"createdAt"`
	UpdatedAt time.Time      `json:"updatedAt"`
	DeletedAt gorm.DeletedAt `json:"deletedAt" gorm:"index"`
}
