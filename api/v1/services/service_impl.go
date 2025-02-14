package services

import (
	"context"

	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
)

func (s *service) CreateTodo(ctx context.Context, req request.CreateTodoRequest) error {
	return nil
}

func (s *service) GetTodo(ctx context.Context, req request.GetTodoRequest) (*response.Todo, error) {
	todo, err := s.todoRepo.Get(ctx, req)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *service) UpdateTodo(ctx context.Context, req request.UpdateTodoRequest) error {
	return nil
}

func (s *service) DeleteTodo(ctx context.Context, req request.DeleteTodoRequest) error {
	return nil
}
