package handler

import (
	"encoding/json"
	"net/http"

	"github.com/salsapunk/api-rest-go/internal/domain"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.Response{
		Message: "pong",
		Status:  200,
	})
}
