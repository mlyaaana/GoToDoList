package auth

import (
	"todolist/domain"
	"todolist/repository/credentials"
	"todolist/repository/user"
)

type Service struct {
	credentialsRepo credentials.Repository
	userRepo        user.Repository
}

func NewService(credentialsRepo credentials.Repository, userRepo user.Repository) *Service {
	return &Service{
		credentialsRepo: credentialsRepo,
		userRepo:        userRepo,
	}
}

func (s *Service) Register(username, password string) (string, error) {
	u := domain.NewUser(username)
	c := domain.NewCredentials(username, password, u.Id)
	if err := s.userRepo.Create(u); err != nil {
		return "", err
	}
	s.credentialsRepo.Create(c)
	return u.Id, nil
}

func (s *Service) Auth(username, password string) (string, error) {
	return s.credentialsRepo.Get(username, password)
}
