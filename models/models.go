package models

import "time"

type Todo struct {
	Id string `json:"id"`
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	DueDate *time.Time `json:"dueDate,omitempty"`
}