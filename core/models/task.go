package models

import (
	"github.com/google/uuid"
	"time"
)

type Task struct {
	ID          string
	Title       string
	Description string
	Date        time.Time
}

func NewTask(name, description string) Task {
	return Task{
		ID:          uuid.NewString(),
		Title:       name,
		Description: description,
		Date:        time.Now(),
	}
}
