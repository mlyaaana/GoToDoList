package api

import (
	"time"

	"todolist/domain"
)

type Task struct {
	Id          string    `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	Deadline    time.Time `json:"deadline"`
}

func NewTask(task *domain.Task) *Task {
	return &Task{
		Id:          task.Id,
		Name:        task.Name,
		Description: task.Description,
		Done:        task.Done,
		CreatedAt:   task.CreatedAt,
		Deadline:    task.Deadline,
	}
}

func NewTasks(tasks []*domain.Task) []*Task {
	apiTasks := make([]*Task, 0)
	for _, t := range tasks {
		apiTasks = append(apiTasks, NewTask(t))
	}

	return apiTasks
}
