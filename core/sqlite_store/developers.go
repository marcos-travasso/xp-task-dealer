package sqlite_store

import (
	"xp-task-dealer/core/models"
)

func (s *SQLiteStore) SaveDeveloper(developer models.Developer) error {
	result := s.conn.Save(&developer)

	return result.Error
}

func (s *SQLiteStore) GetDevelopers() ([]models.Developer, error) {
	var developers []models.Developer

	result := s.conn.Find(&developers)
	if result.Error != nil {
		return nil, result.Error
	}

	return developers, nil
}

func (s *SQLiteStore) GetDeveloperById(id string) (models.Developer, error) {
	var developer models.Developer
	err := s.conn.First(&developer, "id = ?", id).Error

	return developer, err
}

func (s *SQLiteStore) DeleteDeveloper(id string) error {
	result := s.conn.Delete(&models.Developer{}, "id = ?", id)
	return result.Error
}
