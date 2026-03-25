package handler

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/salsapunk/api-rest-go/internal/domain"
	"github.com/salsapunk/api-rest-go/internal/service"
)

type TaskHandler struct {
	handler *service.TaskService
}

func NewTaskHandler(TaskServ *service.TaskService) *TaskHandler {
	return &TaskHandler{
		handler: TaskServ,
	}
}

func (*TaskHandler) HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(domain.Response{
		Message: "pong",
		Status:  200,
	})
}

func (tHandler *TaskHandler) ListAllTasks(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	tasks, err := tHandler.handler.ListAllTasks(ctx)
	if err != nil {
		_ = fmt.Errorf("%w", err)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)

}
