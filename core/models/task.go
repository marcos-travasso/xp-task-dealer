package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          string
	Name        string
	Description string
	Date        time.Time
}

func NewTask(name, description string) Task {
	return Task{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
		Date:        time.Now(),
	}
}
