package handlers

import (
	"encoding/json"
	"net/http"

	"api/response"
)

func HealthHandler(w http.ResponseWriter, r *http.Response) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(response.Response{
		Message: "pong",
		Status:  200,
	})
}
