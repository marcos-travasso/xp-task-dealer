package dto

import (
	"xp-task-dealer/core/models"
)

type DeveloperDTO struct {
	ID          string `json:"id,omitempty"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

func MapDeveloperToDTO(dev models.Developer) DeveloperDTO {
	return DeveloperDTO{
		ID:          dev.ID,
		Name:        dev.Name,
		Description: dev.Description,
	}
}

func MapDevelopersToDTO(developers []models.Developer) []DeveloperDTO {
	devs := make([]DeveloperDTO, 0)
	for _, dev := range developers {
		devs = append(devs, MapDeveloperToDTO(dev))
	}
	return devs
}
