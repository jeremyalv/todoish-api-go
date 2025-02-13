package response

import (
	"time"

	"github.com/jeremyalv/go-todo-api/models"
	"github.com/jeremyalv/go-todo-api/models/request"
)

type CreateTodoResponse struct {
	Code int `json:"code"`
	Message string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Todo *request.CreateTodoRequest `json:"todo"`
}

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