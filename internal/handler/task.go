package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

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

func (tH *TaskHandler) ListById(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := r.PathValue("id")
	taskid, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "ID inválido!", http.StatusBadRequest)
		return
	}

	task, err := tH.service.ListById(ctx, taskid)

	if err != nil {
		log.Printf("Error: %v", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(task)

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
		Status:  http.StatusCreated,
	})
}

func (tH *TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	id := r.PathValue("id")
	taskid, err := strconv.Atoi(id)

	if err != nil {
		http.Error(w, "ID inválido!", http.StatusBadRequest)
		return
	}

	// melhorar tratamento de erros
	if err := tH.service.UpdateTask(ctx, taskid); err != nil {
		log.Printf("Error updating task: %v", err)
		return
	}

	json.NewEncoder(w).Encode(domain.Response{
		Message: fmt.Sprintf("Task with id %d updated!", taskid),
		Status:  202,
	})

}

func (tH *TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	strid := r.PathValue("id")
	taskid, err := strconv.Atoi(strid)
	if err != nil {
		http.Error(w, "ID inválido!", http.StatusBadRequest)
		return
	}

	err = tH.service.DeleteTask(ctx, taskid)
	if err != nil {
		http.Error(w, fmt.Sprintf("%v", err), http.StatusBadGateway)
		return
	}

	json.NewEncoder(w).Encode(domain.Response{
		Message: "Task deleted",
		Status:  204,
	})
}
