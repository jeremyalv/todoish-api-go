package models

import "time"

// Entity Models
type Todo struct {
	Id          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"isCompleted"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
}

// Validator Models

type TodoValidator struct {
	TodoId string `json:"todoId" validate:"todoId"`
}
