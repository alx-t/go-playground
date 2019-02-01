package controllers

import (
	"encoding/json"
	"net/http"
)

func Root(w http.ResponseWriter, r *http.Request) {
	j, err := json.Marshal(map[string]string{"message": "It's working!"})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(j)
}
