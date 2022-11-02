package domain

import (
	"github.com/google/uuid"
)

type Task struct {
	Id          string
	Name        string
	Description string
}

func NewTask(name, description string) *Task {
	return &Task{
		Id:          uuid.New().String(),
		Name:        name,
		Description: description,
	}
}
