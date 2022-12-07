package user

import "todolist/domain"

type Repository interface {
	Create(user *domain.User) error
	Get(id string) *domain.User
	List() []*domain.User
	Delete(id string)
}
