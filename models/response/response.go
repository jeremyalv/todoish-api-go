package response

import (
	"time"

	"github.com/jeremyalv/go-todo-api/models"
)

type GetTodoResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	DueDate *time.Time `json:"dueDate,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
}

type GetMyTodoResponse struct {
	Todos []models.Todo
}

type UpdateTodoResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	DueDate *time.Time `json:"dueDate,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
}