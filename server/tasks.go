package main

import (
	"encoding/json"
	"net/http"
	"xp-task-dealer/core/models"
	"xp-task-dealer/server/dto"
)

func setupTasksRoutes() {
	http.HandleFunc("GET /api/v1/tasks", getTasksHandler)
	http.HandleFunc("POST /api/v1/tasks", createTaskHandler)
}

func getTasksHandler(w http.ResponseWriter, r *http.Request) {
	tasks, err := dbStore.GetTasks()
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapTasksToDTO(tasks), w)
}

func createTaskHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var taskDTO dto.TaskDTO
	err := decoder.Decode(&taskDTO)
	if err != nil {
		respondWithError(http.StatusBadRequest, err, w)
		return
	}

	task := models.NewTask(taskDTO.Title, taskDTO.Description)
	err = dbStore.SaveTask(task)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapTaskToDTO(task), w)
}
