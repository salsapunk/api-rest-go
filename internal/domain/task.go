package domain

import "time"

type Response struct {
	Message string `json:"message"`
	Status  int    `json:"status"`
}

type Task struct {
	Id          int       `json:"id"`          // PK serial
	Title       string    `json:"title"`       // validar
	Description string    `json:"description"` // null
	Done        bool      `json:"done"`        // falso por padrão
	Created_At  time.Time `json:"created_at"`  // timestamp
	Created_By  string    `json:"created_by"`  // FK validar
}

const (
	LISTALL = "SELECT (title, description, done, created_at, created_by) FROM tasks;"
)
