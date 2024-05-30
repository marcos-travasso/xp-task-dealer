package in_memory

import "xp-task-dealer/core/models"

func (s *InMemoryStore) SaveDeveloper(developer models.Developer) error {
	s.developers[developer.ID] = developer

	return nil
}

func (s *InMemoryStore) GetDevelopers() ([]models.Developer, error) {
	devs := make([]models.Developer, 0, len(s.developers))

	for _, dev := range s.developers {
		devs = append(devs, dev)
	}

	return devs, nil
}

func (s *InMemoryStore) GetDeveloperById(id string) (models.Developer, error) {
	return s.developers[id], nil
}
