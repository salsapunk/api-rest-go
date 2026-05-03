package domain

import (
	"time"
)

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Task struct {
	Id          int       `json:"id"`
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	Done        bool      `json:"done"`
	Created_At  time.Time `json:"created_at"`
	Created_By  string    `json:"created_by"`
}

const (
	LISTALL  = "SELECT id, title, description, done, created_at, created_by FROM tasks;"
	LISTBYID = "SELECT id, title, description, done, created_at, created_by FROM tasks WHERE id = $1;"
	CREATE   = "INSERT INTO tasks(title, description, created_by) VALUES($1, $2, $3) RETURNING id;"
	UPDATE   = "UPDATE tasks SET done = true WHERE id = $1;"
	DELETE   = "DELETE FROM tasks WHERE id = $1;"
)
