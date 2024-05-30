package core

import (
	"errors"
	"xp-task-dealer/core/models"
)

var ErrNoSuggestion = errors.New("no suggestion found")

type Service struct {
	s Storer
	d Dealer

	devBlacklists  map[string]map[string]struct{}
	taskBlacklists map[string]map[string]struct{}

	selectedDevs  map[string]string
	selectedTasks map[string]string
}

func NewService(s Storer, d Dealer) *Service {
	return &Service{
		s:              s,
		d:              d,
		devBlacklists:  make(map[string]map[string]struct{}),
		taskBlacklists: make(map[string]map[string]struct{}),

		selectedDevs:  make(map[string]string),
		selectedTasks: make(map[string]string),
	}
}

func (s *Service) GetTaskForDeveloper(id string) (models.Task, error) {
	task, exists := s.selectedDevs[id]
	if exists {
		return s.s.GetTaskById(task)
	}

	dev, err := s.s.GetDeveloperById(id)
	if err != nil {
		return models.Task{}, err
	}

	filteredTasks := s.filterByDeveloper(dev)
	if len(filteredTasks) == 0 {
		return models.Task{}, ErrNoSuggestion
	}

	return s.d.GetTaskForDeveloper(dev, filteredTasks)
}

func (s *Service) GetDeveloperForTask(id string) (models.Developer, error) {
	dev, exists := s.selectedTasks[id]
	if exists {
		return s.s.GetDeveloperById(dev)
	}

	task, err := s.s.GetTaskById(id)
	if err != nil {
		return models.Developer{}, err
	}

	filteredDevs := s.filterByTask(task)
	if len(filteredDevs) == 0 {
		return models.Developer{}, ErrNoSuggestion
	}

	return s.d.GetDeveloperForTask(task, filteredDevs)
}

func (s *Service) GetPairForDeveloper(mainDeveloper models.Developer, task models.Task) (models.Developer, error) {
	devs, err := s.s.GetDevelopers()
	if err != nil {
		return models.Developer{}, err
	}

	if len(devs) == 0 {
		return models.Developer{}, ErrNoSuggestion
	}

	return s.d.GetPairForDeveloper(mainDeveloper, task, devs)
}

func (s *Service) GetAvailableDevelopers() []models.Developer {
	devs, _ := s.s.GetDevelopers()

	var availableDevs []models.Developer

	for _, dev := range devs {
		if _, exists := s.selectedDevs[dev.ID]; !exists {
			availableDevs = append(availableDevs, dev)
		}
	}

	return availableDevs
}

func (s *Service) GetAvailableTasks() []models.Task {
	tasks, _ := s.s.GetTasks()

	var availableTasks []models.Task

	for _, task := range tasks {
		if _, exists := s.selectedTasks[task.ID]; !exists {
			availableTasks = append(availableTasks, task)
		}
	}

	return availableTasks
}

func (s *Service) AddBlacklist(devId, taskId string) error {
	dev, err := s.s.GetDeveloperById(devId)
	if err != nil {
		return err
	}
	task, err := s.s.GetTaskById(taskId)
	if err != nil {
		return err
	}

	devMap, exists := s.devBlacklists[dev.ID]
	if !exists {
		devMap = make(map[string]struct{})
		s.devBlacklists[dev.ID] = devMap
	}
	devMap[task.ID] = struct{}{}

	taskMap, exists := s.taskBlacklists[task.ID]
	if !exists {
		taskMap = make(map[string]struct{})
		s.taskBlacklists[task.ID] = taskMap
	}
	taskMap[dev.ID] = struct{}{}

	return nil
}

func (s *Service) AddSelected(devId, taskId string) error {
	s.selectedDevs[devId] = taskId
	s.selectedTasks[taskId] = devId

	return nil
}

func (s *Service) filterByDeveloper(developer models.Developer) []models.Task {
	tasks := s.GetAvailableTasks()

	devMap, exists := s.devBlacklists[developer.ID]
	if !exists {
		devMap = make(map[string]struct{})
		s.devBlacklists[developer.ID] = devMap
		return tasks
	}

	var filteredTasks []models.Task
	for i, task := range tasks {
		_, exists = devMap[task.ID]
		if !exists {
			filteredTasks = append(filteredTasks, tasks[i])
		}
	}

	return filteredTasks
}

func (s *Service) filterByTask(task models.Task) []models.Developer {
	devs := s.GetAvailableDevelopers()

	taskMap, exists := s.taskBlacklists[task.ID]
	if !exists {
		taskMap = make(map[string]struct{})
		s.taskBlacklists[task.ID] = taskMap
		return devs
	}

	var filteredDevs []models.Developer
	for i, dev := range devs {
		_, exists = taskMap[dev.ID]
		if !exists {
			filteredDevs = append(filteredDevs, devs[i])
		}
	}

	return filteredDevs
}
