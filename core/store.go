package core

import "xp-task-dealer/core/models"

type Storer interface {
	SaveTask(task models.Task) error
	GetTasks() ([]models.Task, error)
	GetTaskById(id string) (models.Task, error)

	SaveDeveloper(developer models.Developer) error
	GetDevelopers() ([]models.Developer, error)
	GetDeveloperById(id string) (models.Developer, error)
}
