package controllers

import (
	"encoding/json"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(map[string]string{"message": "It's working!"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}

func NotImplemented(w http.ResponseWriter, r *http.Request) {
	j, _ := json.Marshal(map[string]string{"message": "Not Implemented yet"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
