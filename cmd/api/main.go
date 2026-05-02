package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/salsapunk/api-rest-go/internal/handler"
	"github.com/salsapunk/api-rest-go/internal/repository"
	"github.com/salsapunk/api-rest-go/internal/service"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	pool, err := pgxpool.New(ctx, os.Getenv("DB_URL"))
	if err != nil {
		log.Fatal("Error on stablishing connection with pool")
	}

	log.Println("Successfully connected to database pool")

	TaskRepo := repository.NewTaskRepository(pool)
	TaskServ := service.NewTaskService(TaskRepo)
	TaskHand := handler.NewTaskHandler(TaskServ)

	mux := http.NewServeMux()

	mux.HandleFunc("/ping", TaskHand.HealthHandler)
	mux.HandleFunc("GET /tasks", TaskHand.ListAllTasks)
	mux.HandleFunc("GET /tasks/{id}", TaskHand.ListById)
	mux.HandleFunc("POST /tasks", TaskHand.CreateTask)
	mux.HandleFunc("UPDATE /tasks/{id}", TaskHand.UpdateTask)
	mux.HandleFunc("DELETE /tasks/{id}", TaskHand.DeleteTask)

	http.Handle("/", mux)

	log.Println("Starting API in port 8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
