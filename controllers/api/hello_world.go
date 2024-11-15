package controllers_api

import (
	"encoding/json"
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	message := map[string]string{
		"hello": "world",
	}
	err := json.NewEncoder(w).Encode(&message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
