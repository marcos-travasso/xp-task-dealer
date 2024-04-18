package sqlite_store

import (
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"testing"
	"xp-task-dealer/core/models"
)

func TestSQLiteStore_Tasks(t *testing.T) {
	s := InitTest()

	// Must have no tasks
	tasks, err := s.GetTasks()
	assert.NoError(t, err)
	assert.Empty(t, tasks)

	task := models.NewTask("test task name", "test task description")

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
	assert.Equal(t, invalidTask, models.Task{})

	// Must update a value
	task.Name = "new_task_name"
	err = s.SaveTask(task)
	assert.NoError(t, err)

	// Must update the task
	updatedTask, err := s.GetTaskById(task.ID)
	assert.NoError(t, err)
	assert.Equal(t, task.Name, updatedTask.Name)
}

func TestSQLiteStore_Developers(t *testing.T) {
	s := InitTest()

	// Must have no tasks
	developers, err := s.GetDevelopers()
	assert.NoError(t, err)
	assert.Empty(t, developers)

	developer := models.NewDeveloper("test developer name", "test developer description")

	// Must save a developer without error
	err = s.SaveDeveloper(developer)
	assert.NoError(t, err)

	// Must return the saved developer
	developers, err = s.GetDevelopers()
	assert.NoError(t, err)
	assert.Len(t, developers, 1)

	savedDeveloper := developers[0]
	assert.Equal(t, developer.ID, savedDeveloper.ID)
	assert.Equal(t, developer.Name, savedDeveloper.Name)
	assert.Equal(t, developer.Description, savedDeveloper.Description)

	// Must return the developer by the ID
	byIdDeveloper, err := s.GetDeveloperById(developer.ID)
	assert.NoError(t, err)
	assert.Equal(t, developer.ID, byIdDeveloper.ID)

	// Must not return a invalid developer
	invalidDeveloper, err := s.GetDeveloperById("invalid_id")
	assert.ErrorIs(t, err, gorm.ErrRecordNotFound)
	assert.Equal(t, invalidDeveloper, models.Developer{})

	// Must update a value
	developer.Name = "new_developer_name"
	err = s.SaveDeveloper(developer)
	assert.NoError(t, err)

	// Must update the developer
	updatedDeveloper, err := s.GetDeveloperById(developer.ID)
	assert.NoError(t, err)
	assert.Equal(t, developer.Name, updatedDeveloper.Name)
}
