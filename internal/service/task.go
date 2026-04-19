package service

import (
	"context"

	"github.com/salsapunk/api-rest-go/internal/domain"
	"github.com/salsapunk/api-rest-go/internal/repository"
)

type TaskService struct {
	repository *repository.TaskRepository
}

func NewTaskService(TaskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repository: TaskRepo,
	}
}

func (tR *TaskService) ListAllTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := tR.repository.ListAllTasks(ctx)
	if err != nil {
		return []domain.Task{}, err
	}

	return tasks, nil
}
