package handlers

import (
	"encoding/json"
	"net/http"
)

func ResponseFormat(w http.ResponseWriter, Message string, StatusCode int, Data interface{}) {
	response := Schema.response{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data
	}
	jsonResponse, err := json.Marshal(response)
	if err != nil {
		ResponseFormat(w, err.Error(), http.StatusInternalServerError, nil)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(response.StatusCode)
	w.Write(jsonResponse)
}
