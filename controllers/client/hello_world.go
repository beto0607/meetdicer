package controllers_client

import (
	"net/http"
)

func GetIndex(w http.ResponseWriter, r *http.Request) {
	message := "Hello World!"
	_, err := w.Write([]byte(message))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}
