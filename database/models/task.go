package models

import "time"

type Task struct {
	Id          string    `gorm:"id"`
	Name        string    `gorm:"name"`
	Description string    `gorm:"description"`
	Done        bool      `gorm:"done"`
	CreatedAt   time.Time `gorm:"created_at"`
	Deadline    time.Time `gorm:"deadline"`
}

func (Task) TableName() string {
	return "tasks"
}
