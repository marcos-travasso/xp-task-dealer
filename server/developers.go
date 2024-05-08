package main

import (
	"encoding/json"
	"net/http"
	"xp-task-dealer/core/models"
	"xp-task-dealer/server/dto"
)

func setupDevelopersRoutes() {
	http.HandleFunc("GET /api/v1/developers", getDevelopersHandler)
	http.HandleFunc("POST /api/v1/developers", createDeveloperHandler)
}

func getDevelopersHandler(w http.ResponseWriter, r *http.Request) {
	devs, err := dbStore.GetDevelopers()
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapDevelopersToDTO(devs), w)
}

func createDeveloperHandler(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var devDTO dto.DeveloperDTO
	err := decoder.Decode(&devDTO)
	if err != nil {
		respondWithError(http.StatusBadRequest, err, w)
		return
	}

	dev := models.NewDeveloper(devDTO.Name, devDTO.Description)
	err = dbStore.SaveDeveloper(dev)
	if err != nil {
		respondWithError(http.StatusInternalServerError, err, w)
		return
	}

	respondWithJSON(dto.MapDeveloperToDTO(dev), w)
}
