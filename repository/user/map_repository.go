package user

import (
	"errors"
	"fmt"
	"todolist/domain"
)

type MapRepository struct {
	storage map[string]*domain.User
}

func NewMapRepository() *MapRepository {
	return &MapRepository{storage: make(map[string]*domain.User)}
}

func (m *MapRepository) Create(user *domain.User) error {
	for _, v := range m.storage {
		if v.Name == user.Name {
			return errors.New("имя занято")
		}
	}
	m.storage[user.Id] = user

	return nil
}

func (m *MapRepository) Get(id string) (*domain.User, error) {
	user, ok := m.storage[id]
	if !ok {
		return nil, fmt.Errorf("user not found")
	}

	return user, nil
}

func (m *MapRepository) List() []*domain.User {
	result := make([]*domain.User, 0)

	for _, v := range m.storage {
		result = append(result, v)
	}

	return result
}

func (m *MapRepository) Delete(id string) {
	delete(m.storage, id)
}
