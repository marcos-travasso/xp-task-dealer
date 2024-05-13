package core

import "xp-task-dealer/core/models"

type Dealer interface {
	GetDeveloperForTask(task models.Task, developers []models.Developer) (models.Developer, error)
	GetTaskForDeveloper(developer models.Developer, tasks []models.Task) (models.Task, error)
	GetPairForDeveloper(mainDeveloper models.Developer, task models.Task, developers []models.Developer) (models.Developer, error)
}
