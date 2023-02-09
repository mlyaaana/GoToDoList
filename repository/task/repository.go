package task

import "todolist/domain"

type Repository interface {
	Create(task *domain.Task) error
	Get(id string) (*domain.Task, error)
	List() ([]*domain.Task, error)
	Delete(id string)
	Done(id string) error
	ListCompletedTasks() ([]*domain.Task, error)
}
