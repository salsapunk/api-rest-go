package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/salsapunk/api-rest-go/internal/domain"
	"github.com/salsapunk/api-rest-go/internal/service"
)

type TaskHandler struct {
	service *service.TaskService
}

func NewTaskHandler(TaskServ *service.TaskService) *TaskHandler {
	return &TaskHandler{
		service: TaskServ,
	}
}

func (tH *TaskHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.Response{
		Message: "pong",
		Status:  200,
	})
}

func (tH *TaskHandler) ListAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	tasks, err := tH.service.ListAllTasks(ctx)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func (tH *TaskHandler) CreateTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var task *domain.Task
	json.NewDecoder(r.Body).Decode(&task)

	id, err := tH.service.CreateTask(ctx, task)
	if err != nil {
		log.Printf("Error creating new task: %v", err)
		return
	}

	json.NewEncoder(w).Encode(domain.Response{
		Message: fmt.Sprint("Task created with id ", id),
		Status:  201,
	})
}
