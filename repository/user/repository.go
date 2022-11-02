package user

import "todolist/domain"

type Repository interface {
	Create(user *domain.User)
	Get(id string) *domain.User
	List(id string) []*domain.User
	Delete(id string)
}
