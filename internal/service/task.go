package service

import "github.com/salsapunk/api-rest-go/internal/repository"

type TaskService struct {
	service *repository.TaskRepository
}

func NewTaskService(TaskRepo *repository.TaskRepository) *TaskService {
	return &TaskService{
		service: TaskRepo,
	}
}
