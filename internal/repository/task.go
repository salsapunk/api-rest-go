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

	var tasks []domain.Task
	var task domain.Task

	for rows.Next() {
		err = rows.Scan(
			&task.Id,
			&task.Title,
			&task.Description,
			&task.Done,
			&task.Created_At,
			&task.Created_By,
		)

		if err != nil {
			return []domain.Task{}, err
		}

		tasks = append(tasks, task)
	}

	rows.Close()

	return tasks, nil
}

func (tR *TaskRepository) ListById(ctx context.Context, id int) (domain.Task, error) {
	row := tR.pool.QueryRow(ctx, domain.LISTBYID, id)

	var task domain.Task

	err := row.Scan(
		&task.Id,
		&task.Title,
		&task.Description,
		&task.Done,
		&task.Created_At,
		&task.Created_By,
	)
	if err != nil {
		return domain.Task{}, err
	}

	return task, nil
}

func (tR *TaskRepository) CreateTask(ctx context.Context, task *domain.Task) (int, error) {
	row := tR.pool.QueryRow(ctx, domain.CREATE, &task.Title, &task.Description)

	var id int

	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (tR *TaskRepository) UpdateTask(ctx context.Context, id int) error {
	_, err := tR.pool.Exec(ctx, domain.UPTADE, id)
	if err != nil {
		return err
	}

	return nil
}

func (tR *TaskRepository) DeleteTask(ctx context.Context, id int) error {
	_, err := tR.pool.Exec(ctx, domain.DELETE, id)
	if err != nil {
		return err
	}

	return nil
}
