## Api Rest em Go

API REST para gerenciamento de tasks

## Como rodar

1. Clone o repositório
2. Copie o .env.example e preencha: DB_URL=postgres://user:senha@localhost:5432/dbname
3. Rode as migrations: 'psql -f migrations/001_create_tasks.sql'
4. `go run ./cmd/api/main.go`

## Endpoints

| Método | Rota            | Descrição          |
|--------|-----------------|--------------------|
| GET    | /tasks          | Lista todas        |
| GET    | /tasks/{id}     | Busca por ID       |
| POST   | /tasks          | Cria nova task     |
| PUT    | /tasks/{id}     | Marca como feita   |
| DELETE | /tasks/{id}     | Remove             |

## Exemplo de request

POST /tasks
```json
{ "title": "estudar Go", "description": "focar em testes" }
```
