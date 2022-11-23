package task

import (
	"todolist/domain"
)

type MapRepository struct {
	storage map[string]*domain.Task
}

func NewMapRepository() *MapRepository {
	return &MapRepository{storage: make(map[string]*domain.Task)}
}

func (m *MapRepository) Create(task *domain.Task) {
	m.storage[task.Id] = task
}

func (m *MapRepository) Get(id string) *domain.Task {
	return m.storage[id]
}

func (m *MapRepository) List() []*domain.Task {
	result := make([]*domain.Task, 0)

	for _, v := range m.storage {
		result = append(result, v)
	}

	return result
}

func (m *MapRepository) Delete(id string) {
	delete(m.storage, id)
}

func (m *MapRepository) Done(id string) {
	m.storage[id].Done = !m.storage[id].Done
}

func (m *MapRepository) ListCompletedTasks() []*domain.Task {
	result := make([]*domain.Task, 0)
	for _, v := range m.storage {
		if v.Done {
			result = append(result, v)
		}
	}

	return result
}
