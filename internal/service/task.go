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
	validate   validator.Validate
}

func NewTaskService(TaskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{
		repository: TaskRepo,
		validate:   *validator.New(),
	}
}

func (tS *TaskService) ListAllTasks(ctx context.Context) ([]domain.Task, error) {
	tasks, err := tS.repository.ListAllTasks(ctx)
	if err != nil {
		return nil, err
	}

	return tasks, nil
}

func (tS *TaskService) ListById(ctx context.Context, id int) (domain.Task, error) {
	task, err := tS.repository.ListById(ctx, id)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (tS *TaskService) CreateTask(ctx context.Context, task *domain.Task) (int, error) {
	task.Created_At = time.Now()

	err := tS.validate.Struct(task)
	if err != nil {
		return 0, err
	}

	id, err := tS.repository.CreateTask(ctx, task)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tS *TaskService) UpdateTask(ctx context.Context, id int) error {
	err := tS.repository.UpdateTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}

func (tS *TaskService) DeleteTask(ctx context.Context, id int) error {
	err := tS.repository.DeleteTask(ctx, id)
	if err != nil {
		return err
	}

	return nil
}
