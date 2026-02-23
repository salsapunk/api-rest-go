package main

import (
	"log"
	"net/http"

	"api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	mux := mux.NewRouter()

	mux.HandleFunc("/ping", handlers.HealthHandler)

	http.Handle("/", mux)

	log.Println("Server starting at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
