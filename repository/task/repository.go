package task

import "todolist/domain"

type Repository interface {
	Create(task *domain.Task)
	Get(id string) *domain.Task
	List() []*domain.Task
	Delete(id string)
	Done(id string)
}
