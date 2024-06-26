package sqlite_store

import "xp-task-dealer/core/models"

func (s *SQLiteStore) SaveTask(task models.Task) error {
	result := s.conn.Save(&task)

	return result.Error
}

func (s *SQLiteStore) GetTasks() ([]models.Task, error) {
	var tasks []models.Task

	result := s.conn.Find(&tasks)
	if result.Error != nil {
		return nil, result.Error
	}

	return tasks, nil
}

func (s *SQLiteStore) GetTaskById(id string) (models.Task, error) {
	var task models.Task
	err := s.conn.First(&task, "id = ?", id).Error

	return task, err
}

func (s *SQLiteStore) DeleteTask(id string) error {
	result := s.conn.Delete(&models.Task{}, "id = ?", id)
	return result.Error
}
