package domain

import "time"

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Task struct {
	Id          int64     `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	CreatedAt   time.Time `json:"created_at"`
	//CreatedBy   string    `json:"created_by"`
}

const (
	CREATE_TASK = "INSERT INTO task(title, description, created_at) VALUES($1, $2, $3) RETURNING id;"
	LIST_TASKS  = "SELECT id, title, description, done, created_at FROM tasks;"
)
