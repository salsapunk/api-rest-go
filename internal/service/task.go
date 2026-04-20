package service

import (
	"context"
	"time"

	"github.com/go-playground/validator/v10"
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

func (tR *TaskService) CreateTask(ctx context.Context, task *domain.Task) (int, error) {
	task.Created_At = time.Now()

	validate := validator.New()
	err := validate.Struct(task)
	if err != nil {
		return 0, err
	}

	id, err := tR.repository.CreateTask(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}
