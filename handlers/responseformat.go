package handlers

import (
	"encoding/json"
	"net/http"

	Schema "github.com/sridharu346/loginPractice/schema"
)

func ResponseFormat(w http.ResponseWriter, Message string, StatusCode int, Data interface{}) {
	response := Schema.Response{
		StatusCode: StatusCode,
		Message:    Message,
		Data:       Data,
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
