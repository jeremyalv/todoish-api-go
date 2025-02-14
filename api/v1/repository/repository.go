package repository

import (
	"context"
	"database/sql"
	"github.com/jeremyalv/go-todo-api/models/response"

	"github.com/jeremyalv/go-todo-api/models/request"
)

type todoRepoImpl struct {
	DB *sql.DB
}

type TodoRepository interface {
	Save(ctx context.Context, req request.CreateTodoRequest) (int64, error)
	Get(ctx context.Context, req request.GetTodoRequest) (*response.Todo, error)
	GetUserTodos(ctx context.Context, req request.GetMyTodoRequest) ([]*request.Todo, error)
	Update(ctx context.Context, req request.UpdateTodoRequest) error
	Delete(ctx context.Context, req request.DeleteTodoRequest) error
}

func New(db *sql.DB) *todoRepoImpl {
	return &todoRepoImpl{
		DB: db,
	}
}

func (o *todoRepoImpl) Save(ctx context.Context, req request.CreateTodoRequest) (int64, error) {
	return o.save(ctx, req)
}

func (o *todoRepoImpl) Get(ctx context.Context, req request.GetTodoRequest) (*response.Todo, error) {
	return o.get(ctx, req)
}

func (o *todoRepoImpl) GetUserTodos(ctx context.Context, req request.GetMyTodoRequest) ([]*request.Todo, error) {
	return o.getByOwner(ctx, req)
}

func (o *todoRepoImpl) Update(ctx context.Context, req request.UpdateTodoRequest) error {
	return o.update(ctx, req)
}

func (o *todoRepoImpl) Delete(ctx context.Context, req request.DeleteTodoRequest) error {
	return o.delete(ctx, req)
}
