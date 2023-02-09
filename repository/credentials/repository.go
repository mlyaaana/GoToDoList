package credentials

import "todolist/domain"

type Repository interface {
	Create(credentials *domain.Credentials) error
	Get(username, password string) (string, error)
}
