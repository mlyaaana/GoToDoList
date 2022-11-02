package task

import (
	"todolist/domain"
	"todolist/repository/task"
)

type Service struct {
	repo task.Repository
}

func (s *Service) Create(task *domain.Task) {
	s.repo.Create(task)
}

func (s *Service) Get(id string) *domain.Task {
	return s.repo.Get(id)
}

func (s *Service) List() []*domain.Task {
	return s.repo.List()
}

func (s *Service) Delete(id string) {
	s.repo.Delete(id)
}