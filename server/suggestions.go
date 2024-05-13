package main

import (
	"net/http"
	"xp-task-dealer/server/dto"
)

func setupSuggestionsRoutes() {
	http.HandleFunc("GET /api/v1/suggestions/developer/{id}", getDeveloperSuggestionHandler)
	http.HandleFunc("GET /api/v1/suggestions/task/{id}", getTaskSuggestionHandler)
}

func getDeveloperSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	dev, err := dbStore.GetDeveloperById(id)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	tasks, err := dbStore.GetTasks()
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	task, err := dealer.GetTaskForDeveloper(dev, tasks)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapTaskToDTO(task), w)
}

func getTaskSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	task, err := dbStore.GetTaskById(id)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	devs, err := dbStore.GetDevelopers()
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	dev, err := dealer.GetDeveloperForTask(task, devs)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapDeveloperToDTO(dev), w)
}
