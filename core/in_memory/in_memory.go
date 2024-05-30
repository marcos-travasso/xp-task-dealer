package in_memory

import "xp-task-dealer/core/models"

type InMemoryStore struct {
	developers map[string]models.Developer
	tasks      map[string]models.Task
}

func Init() *InMemoryStore {
	return &InMemoryStore{
		developers: make(map[string]models.Developer),
		tasks:      make(map[string]models.Task),
	}
}
