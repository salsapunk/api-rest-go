package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/salsapunk/api-rest-go/internal/domain"
)

type TaskRepository struct {
	pool *pgxpool.Pool
}

func NewTaskRepository(pool *pgxpool.Pool) *TaskRepository {
	return &TaskRepository{
		pool: pool,
	}
}

func (tR *TaskRepository) ListAllTasks(ctx context.Context) ([]domain.Task, error) {
	rows, err := tR.pool.Query(ctx, domain.LISTALL)
	if err != nil {
		return []domain.Task{}, err
	}

	var task domain.Task
	tasks := []domain.Task{}

	for rows.Next() {
		rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Done,
			&task.Created_At,
			&task.Created_By,
		)
		tasks = append(tasks, task)
	}

	return []domain.Task{}, nil
}
