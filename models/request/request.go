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

type UpdateTodoRequest struct {
	TodoId string `json:"todoId"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	DueDate *time.Time `json:"dueDate,omitempty"`
}

type DeleteTodoRequest struct {
	TodoId string `json:"todoId"`
}