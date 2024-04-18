package core

type TasksStorer interface {
	SaveTask(task Task) error
	GetTasks() ([]Task, error)
	GetTaskById(id string) (Task, error)
}