package services

import (
	"context"

	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
)

//go:generate mockgen -source services.go -destination service_mock.go -package services
type IServices interface {
	CreateTodo(ctx context.Context, req request.CreateTodoRequest) error
	GetTodo(ctx context.Context, req request.GetTodoRequest) (*response.GetTodoResponse, error)
	UpdateTodo(ctx context.Context, req request.UpdateTodoRequest) (*response.UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, req request.DeleteTodoRequest) (*response.DeleteTodoResponse, error)
}