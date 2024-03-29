package task

import (
	"todolist/domain"
	"todolist/repository/task"
)

type Service struct {
	repo task.Repository
}

func NewService(repo task.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) Create(task *domain.Task) error {
	return s.repo.Create(task)
}

func (s *Service) Get(id string) (*domain.Task, error) {
	return s.repo.Get(id)
}

func (s *Service) List() ([]*domain.Task, error) {
	return s.repo.List()
}

func (s *Service) Delete(id string) {
	s.repo.Delete(id)
}

func (s *Service) Done(id string) error {
	return s.repo.Done(id)
}

func (s *Service) ListCompletedTasks() ([]*domain.Task, error) {
	return s.repo.ListCompletedTasks()
}
