package main

import (
	"encoding/json"
	"net/http"
)

func respondWithError(status int, err error, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(status)
	response, _ := json.Marshal(map[string]string{"error": err.Error()})
	w.Write(response)
}

func respondWithJSON(data interface{}, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)
	response, _ := json.Marshal(data)
	w.Write(response)
}
