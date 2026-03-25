// service trata de um intermediário entre o repository e o handler, mas com outras coisas,
// como regras, por exemplo: títulos duplicados
package service

import (
	"context"
	"fmt"

	"github.com/salsapunk/api-rest-go/internal/domain"
	"github.com/salsapunk/api-rest-go/internal/repository"
)

type TaskService struct {
	service *repository.TaskRepository
}

func NewTaskService(TaskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{
		service: TaskRepo,
	}
}

func (tService *TaskService) ListAllTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := tService.service.ListAllTasks(ctx)
	if err != nil {
		_ = fmt.Errorf("service %w", err)
		return []domain.Task{}, err
	}

	if tasks == nil {
		return []domain.Task{}, nil
	}

	return tasks, nil
}
