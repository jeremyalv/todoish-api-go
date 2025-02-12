package request

import (
	"time"
)

type CreateTodoRequest struct {
	Title string `json:"title"`
	Description string `json:"description"`
	DueDate *time.Time `json:"dueDate,omitempty"`
}

type GetTodoRequest struct {
	TodoId string `json:"todoId"`
}