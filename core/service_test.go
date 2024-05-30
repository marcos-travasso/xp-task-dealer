package core

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
	"xp-task-dealer/core/in_memory"
	"xp-task-dealer/core/models"
)

func TestService_GetAvailableDevelopers(t *testing.T) {
	t.Run("should return all non selected developers", func(t *testing.T) {
		store := in_memory.Init()
		s := NewService(store, nil)

		dev1 := models.NewDeveloper("dev1", "dev1")
		dev2 := models.NewDeveloper("dev2", "dev2")

		err := store.SaveDeveloper(dev1)
		require.NoError(t, err)
		err = store.SaveDeveloper(dev2)
		require.NoError(t, err)

		s.selectedDevs[dev1.ID] = struct{}{}

		devs := s.GetAvailableDevelopers()

		assert.Len(t, devs, 1)
		assert.Contains(t, devs, dev2)
	})
}

func TestService_GetAvailableTasks(t *testing.T) {
	t.Run("should return all non selected tasks", func(t *testing.T) {
		store := in_memory.Init()
		s := NewService(store, nil)

		task1 := models.NewTask("task1", "task1")
		task2 := models.NewTask("task2", "task2")

		err := store.SaveTask(task1)
		require.NoError(t, err)
		err = store.SaveTask(task2)
		require.NoError(t, err)

		s.selectedTasks[task1.ID] = struct{}{}

		tasks := s.GetAvailableTasks()

		assert.Len(t, tasks, 1)
		assert.Contains(t, tasks, task2)
	})
}

func TestService_AddBlacklist(t *testing.T) {
	t.Run("should add blacklist", func(t *testing.T) {
		store := in_memory.Init()
		s := NewService(store, nil)

		dev := models.NewDeveloper("dev", "dev")
		task := models.NewTask("task", "task")

		err := store.SaveDeveloper(dev)
		require.NoError(t, err)
		err = store.SaveTask(task)
		require.NoError(t, err)

		require.Empty(t, s.devBlacklists[dev.ID])
		require.Empty(t, s.taskBlacklists[task.ID])

		err = s.AddBlacklist(dev.ID, task.ID)
		require.NoError(t, err)

		assert.Len(t, s.devBlacklists[dev.ID], 1)
		assert.Contains(t, s.devBlacklists[dev.ID], task.ID)
		assert.Len(t, s.taskBlacklists[task.ID], 1)
		assert.Contains(t, s.taskBlacklists[task.ID], dev.ID)
	})
}

func TestService_AddSelected(t *testing.T) {
	t.Run("should add selected", func(t *testing.T) {
		store := in_memory.Init()
		s := NewService(store, nil)

		dev := models.NewDeveloper("dev", "dev")
		task := models.NewTask("task", "task")

		err := store.SaveDeveloper(dev)
		require.NoError(t, err)
		err = store.SaveTask(task)
		require.NoError(t, err)

		require.Empty(t, s.selectedDevs)
		require.Empty(t, s.selectedTasks)

		err = s.AddSelected(dev.ID, task.ID)
		require.NoError(t, err)

		assert.Len(t, s.selectedDevs, 1)
		assert.Contains(t, s.selectedDevs, dev.ID)
		assert.Len(t, s.selectedTasks, 1)
		assert.Contains(t, s.selectedTasks, task.ID)
	})
}

func TestService_filterByDeveloper(t *testing.T) {
	t.Run("should return tasks out of blacklist", func(t *testing.T) {
		store := in_memory.Init()
		s := NewService(store, nil)

		dev := models.NewDeveloper("dev", "dev")
		task1 := models.NewTask("task1", "task1")
		task2 := models.NewTask("task2", "task2")

		err := store.SaveDeveloper(dev)
		require.NoError(t, err)
		err = store.SaveTask(task1)
		require.NoError(t, err)
		err = store.SaveTask(task2)
		require.NoError(t, err)

		err = s.AddBlacklist(dev.ID, task1.ID)
		require.NoError(t, err)

		tasks := s.filterByDeveloper(dev)

		assert.Len(t, tasks, 1)
		assert.Contains(t, tasks, task2)
	})
}

func TestService_filterByTask(t *testing.T) {
	t.Run("should return developers out of blacklist", func(t *testing.T) {
		store := in_memory.Init()
		s := NewService(store, nil)

		dev1 := models.NewDeveloper("dev1", "dev1")
		dev2 := models.NewDeveloper("dev2", "dev2")
		task := models.NewTask("task", "task")

		err := store.SaveDeveloper(dev1)
		require.NoError(t, err)
		err = store.SaveDeveloper(dev2)
		require.NoError(t, err)
		err = store.SaveTask(task)
		require.NoError(t, err)

		err = s.AddBlacklist(dev1.ID, task.ID)
		require.NoError(t, err)

		developers := s.filterByTask(task)

		assert.Len(t, developers, 1)
		assert.Contains(t, developers, dev2)
	})
}
