package xp_task_dealer

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"time"
)

func TestSQLiteStore(t *testing.T) {
	s := InitTestDB()

	// Must have no tasks
	tasks, err := s.GetTasks()
	assert.NoError(t, err)
	assert.Empty(t, tasks)

	task := Task{
		ID:          "task_id",
		Name:        "task name",
		Description: "task description",
		Date:        time.Now(),
	}

	// Must save a task without error
	err = s.SaveTask(task)
	assert.NoError(t, err)

	// Must return the saved task
	tasks, err = s.GetTasks()
	assert.NoError(t, err)
	assert.Len(t, tasks, 1)

	savedTask := tasks[0]
	assert.Equal(t, task.ID, savedTask.ID)
	assert.Equal(t, task.Name, savedTask.Name)
	assert.Equal(t, task.Description, savedTask.Description)
	assert.Equal(t, task.Date.Unix(), savedTask.Date.Unix())

	// Must return the task by the ID
	byIdTask, err := s.GetTaskById(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task.ID, byIdTask.ID)

	// Must not return a invalid task
	invalidTask, err := s.GetTaskById("invalid_id")
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Equal(t, invalidTask, Task{})

	// Must update a value
	task.Name = "new_task_name"
	err = s.SaveTask(task)
	assert.NoError(t, err)

	// Must update the task
	updatedTask, err := s.GetTaskById(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task.Name, updatedTask.Name)
}
