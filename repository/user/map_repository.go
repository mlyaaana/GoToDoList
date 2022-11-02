package user

import "todolist/domain"

type MapRepository struct {
	storage map[string]*domain.User
}

func (m *MapRepository) Create(user *domain.User) {
	m.storage[user.Id] = user
}

func (m *MapRepository) Get(id string) *domain.User {
	return m.storage[id]
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
