package models

import "time"

type Task struct {
	ID          string
	Name        string
	Description string
	Date        time.Time
}
