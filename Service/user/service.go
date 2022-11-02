package user

import (
	"todolist/domain"
	"todolist/repository/user"
)

type Service struct {
	repo user.Repository
}

func (s *Service) Create(user *domain.User) {
	s.repo.Create(user)
}

func (s *Service) Get(id string) *domain.User {
	return s.repo.Get(id)
}

func (s *Service) List() []*domain.User {
	return s.repo.List()
}

func (s *Service) Delete(id string) {
	s.repo.Delete(id)
}
