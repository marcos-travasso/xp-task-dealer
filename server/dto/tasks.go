package dto

import (
	"time"
	"xp-task-dealer/core/models"
)

type TaskDTO struct {
	ID          string `json:"id,omitempty"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Date        string `json:"date,omitempty"`
}

func MapTaskToDTO(task models.Task) TaskDTO {
	return TaskDTO{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Date:        task.Date.Format(time.DateTime),
	}
}

func MapTasksToDTO(tasks []models.Task) []TaskDTO {
	tasksDTO := make([]TaskDTO, 0)
	for _, task := range tasks {
		tasksDTO = append(tasksDTO, MapTaskToDTO(task))
	}
	return tasksDTO
}
