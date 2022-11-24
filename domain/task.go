package domain

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	Id          string
	Name        string
	Description string
	Done        bool
	CreatedAt   time.Time
	Deadline    time.Time
}

func NewTask(name, description string, deadline time.Time) *Task {
	return &Task{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
		Done:        false,
		CreatedAt:   time.Now(),
		Deadline:    deadline,
	}
}
