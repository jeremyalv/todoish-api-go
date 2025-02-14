package request

import (
	"time"
)

type CreateTodoRequest struct {
	OwnerId     string     `json:"ownerId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
}

type GetMyTodoRequest struct {
	UserId string `json:"userId"`
}

type GetTodoRequest struct {
	TodoId string `json:"todoId"`
}

type UpdateTodoRequest struct {
	TodoId      string     `json:"todoId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"isCompleted"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
}

type DeleteTodoRequest struct {
	TodoId string `json:"todoId"`
}

// Type to insert Todo to DB
type Todo struct {
	Id          string     `json:"id"`
	OwnerId     string     `json:"ownerId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"isCompleted"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
}
