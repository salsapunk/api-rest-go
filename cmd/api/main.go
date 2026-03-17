package main

import (
	"log"
	"net/http"

	"github.com/salsapunk/api-resto-go/handler"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping", handler.healthHandler)

	http.Handle("/", mux)

	log.Println("Server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
