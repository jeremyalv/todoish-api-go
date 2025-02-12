package services

import (
	"context"

	"github.com/jeremyalv/go-todo-api/config"
	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
)

//go:generate mockgen -source services.go -destination service_mock.go -package services
type IServices interface {
	CreateTodo(ctx context.Context, req request.CreateTodoRequest) error
	GetTodo(ctx context.Context, req request.GetTodoRequest) (*response.GetTodoResponse, error)
	UpdateTodo(ctx context.Context, req request.UpdateTodoRequest) (*response.UpdateTodoResponse, error)
	DeleteTodo(ctx context.Context, req request.DeleteTodoRequest) (error)
}

type service struct {
	// In a real world app, this type would contain many more fields, e.g. logger, tasks, and other interfaces
	cfg config.Config
}


//
func (s *service) WithConfig(cfg config.Config) *service {
	s.cfg = cfg
	return s
}