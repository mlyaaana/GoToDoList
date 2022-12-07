package credentials

import (
	"errors"
	"todolist/domain"
)

type MapRepository struct {
	storage map[string]*domain.Credentials
}

func NewMapRepository() *MapRepository {
	return &MapRepository{storage: make(map[string]*domain.Credentials)}
}

func (m *MapRepository) Create(credentials *domain.Credentials) {
	m.storage[credentials.Username] = credentials
}

func (m *MapRepository) Get(username, password string) (string, error) {
	if creds, ok := m.storage[username]; ok && creds.Password == password {
		return creds.UserId, nil
	}

	return "", errors.New("неверный логин или пароль")
}
