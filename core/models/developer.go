package models

import (
	"github.com/google/uuid"
)

type Developer struct {
	ID          string
	Name        string
	Description string
}

func NewDeveloper(name, description string) Developer {
	return Developer{
		ID:          uuid.NewString(),
		Name:        name,
		Description: description,
	}
}
