package in_memory

import "xp-task-dealer/core/models"

func (s *InMemoryStore) SaveTask(task models.Task) error {
	s.tasks[task.ID] = task

	return nil
}

func (s *InMemoryStore) GetTasks() ([]models.Task, error) {
	tasks := make([]models.Task, 0, len(s.tasks))

	for _, task := range s.tasks {
		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (s *InMemoryStore) GetTaskById(id string) (models.Task, error) {
	return s.tasks[id], nil
}
