package services

import (
	"context"

	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
)

func (s *service) CreateTodo(ctx context.Context, req request.CreateTodoRequest) error {
	_, err := s.todoRepo.Save(ctx, req)
	return err
}

func (s *service) GetTodo(ctx context.Context, req request.GetTodoRequest) (*response.Todo, error) {
	res, err := s.todoRepo.Get(ctx, req)
	return res, err
}

func (s *service) UpdateTodo(ctx context.Context, req request.UpdateTodoRequest) error {
	err := s.todoRepo.Update(ctx, req)
	return err
}

func (s *service) DeleteTodo(ctx context.Context, req request.DeleteTodoRequest) error {
	err := s.todoRepo.Delete(ctx, req)
	return err
}
