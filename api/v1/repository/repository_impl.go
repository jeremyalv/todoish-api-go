package repository

import (
	"context"
	"database/sql"
	e "errors"
	"fmt"
	"github.com/jeremyalv/go-todo-api/models/response"

	"github.com/jeremyalv/go-todo-api/models/request"
)

func (o *todoRepoImpl) save(ctx context.Context, todoObj request.Todo) (int64, error) {
	var id int64

	query := `INSERT INTO todos (owner_id, title, description, is_completed, due_date) VALUES (?, ?, ?, ?, ?)`
	stmt, err := o.DB.Prepare(query)
	if err != nil {
		return id, fmt.Errorf("error while preparing query: %v", err)
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		todoObj.OwnerId, todoObj.Title, todoObj.Description, todoObj.IsCompleted, todoObj.DueDate,
	)
	if err != nil {
		return id, fmt.Errorf("error while executing statement: %v", err)
	}

	id, err = res.LastInsertId()
	if err != nil {
		return id, fmt.Errorf("error while reading LastInsertId: %v", err)
	}

	return id, nil
}

func (o *todoRepoImpl) get(ctx context.Context, req request.GetTodoRequest) (*response.Todo, error) {
	query := `SELECT id, owner_id, title, description, is_completed, due_date FROM todos WHERE id=?`
	rowQuery := o.DB.QueryRow(query, req.TodoId)
	res := new(response.Todo)
	err := rowQuery.Scan(&res.Id, &res.OwnerId, &res.Title, &res.Description, &res.IsCompleted, &res.DueDate)

	if err != nil {
		if e.Is(err, sql.ErrNoRows) {
			// In production, you would call span.RecordError and create a custom error
			// The custom error should explain the exact issue happening here
			return nil, fmt.Errorf("error due to no rows returned: %v", err)
		}

		// Here, you'd also return a specific custom error of a more general type (e.g. MyInternalServerError)
		return nil, fmt.Errorf("error while preparing query: %v", err)
	}
	return res, nil
}

func (o *todoRepoImpl) getByOwner(ctx context.Context, req request.GetMyTodoRequest) ([]*request.Todo, error) {
	// Create the result slice todos of unknown size
	todos := []*request.Todo{}

	query := `SELECT id, owner_id, title, description, is_completed, due_date FROM todos WHERE owner_id=?`
	rows, err := o.DB.Query(query, req.UserId)
	if err != nil {
		return nil, fmt.Errorf("error while querying from database: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		// Create new Todo object from the scanned values
		res := new(request.Todo)
		err := rows.Scan(&res.Id, &res.OwnerId, &res.Title, &res.Description, &res.IsCompleted, &res.DueDate)
		if err != nil {
			// Errors after this point should return the `todos` slice since some values may have been read from the DB before the app throws an error
			return todos, fmt.Errorf("error while scanning rows: %v", err)
		}

		todos = append(todos, res)
	}
	err = rows.Err()
	if err != nil {
		return todos, fmt.Errorf("encountered an error while iterating through rows: %v", err)
	}

	return todos, nil
}

func (o *todoRepoImpl) update(ctx context.Context, req request.UpdateTodoRequest) error {
	query := `UPDATE todos SET title=? description=? is_completed=? due_date=? WHERE id=?`
	stmt, err := o.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error while preparing query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.Title, req.Description, req.IsCompleted, req.DueDate, req.TodoId)
	if err != nil {
		return fmt.Errorf("error while scanning rows: %v", err)
	}

	return nil
}

func (o *todoRepoImpl) delete(ctx context.Context, req request.DeleteTodoRequest) error {
	query := `DELETE FROM todos WHERE id=?`
	stmt, err := o.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("error while preparing query: %v", err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(req.TodoId)
	if err != nil {
		return fmt.Errorf("error while preparing query: %v", err)
	}

	return nil
}
