package main

import (
	"errors"
	"net/http"
	"xp-task-dealer/core"
	"xp-task-dealer/server/dto"
)

func setupSuggestionsRoutes() {
	http.HandleFunc("GET /api/v1/suggestions/developer/{id}", getDeveloperSuggestionHandler)
	http.HandleFunc("GET /api/v1/suggestions/task/{id}", getTaskSuggestionHandler)

	http.HandleFunc("POST /api/v1/suggestions/blacklist", markBlacklistHandler)
	http.HandleFunc("POST /api/v1/suggestions/selected", markSelectedHandler)
}

func getDeveloperSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	task, err := service.GetTaskForDeveloper(id)
	if errors.Is(err, core.ErrNoSuggestion) {
		respondWithError(http.StatusNotFound, err, w)
		return
	} else if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapTaskToDTO(task), w)
}

func getTaskSuggestionHandler(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	dev, err := service.GetDeveloperForTask(id)
	if errors.Is(err, core.ErrNoSuggestion) {
		respondWithError(http.StatusNotFound, err, w)
		return
	} else if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapDeveloperToDTO(dev), w)
}

func markBlacklistHandler(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("task_id")
	devId := r.URL.Query().Get("dev_id")

	err := service.AddBlacklist(devId, taskId)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(nil, w)
}

func markSelectedHandler(w http.ResponseWriter, r *http.Request) {
	taskId := r.URL.Query().Get("task_id")
	devId := r.URL.Query().Get("dev_id")

	err := service.AddSelected(devId, taskId)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(nil, w)
}
