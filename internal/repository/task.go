package repository

import (
	"context"
	"log"

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

// ListAllTasks é a função de listagem dos dados, que vai usar o "SELECT x FROM table" do sql
// ela retorna um slice de tasks, como uma lista das tasks que foram recebidas pelo SELECT
// e um erro, que, se bem sucedida, será nil
func (tRepository *TaskRepository) ListAllTasks(ctx context.Context) ([]domain.Task, error) {
	// rows armazenará todas as rows retornadas do comando SELECT, dado pelo Query()
	rows, err := tRepository.pool.Query(ctx, domain.LIST_TASKS)
	// tratamento de erro
	if err != nil {
		log.Println("repository1", err)
		return []domain.Task{}, err
	}

	// criando um slice como lista das tasks que vão ser selecionadas
	var taskList []domain.Task
	// criando o modelo das tasks que serão guardadas no slice
	var taskModel domain.Task

	// loop checa se tem outra linha
	for rows.Next() {
		// Scan() pega as informações da linha e copia para as variáveis mandadas
		// como parâmetro
		err := rows.Scan(
			// a cada loop essas informações são sobrepostas com as das novas linhas
			&taskModel.Id,
			&taskModel.Title,
			&taskModel.Description,
			&taskModel.Done,
			&taskModel.CreatedAt,
		)

		// tratamento de erro
		if err != nil {
			log.Println("repository2", err)
			return []domain.Task{}, err
		}

		// adiciona o modelo no slice das tasks
		taskList = append(taskList, taskModel)
	}

	// garante que as rows foram fechadas e a conexão esteja pronta para ser usada novamente
	rows.Close()

	return taskList, nil
}
