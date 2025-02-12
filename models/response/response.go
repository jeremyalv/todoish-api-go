package response

import "time"

type GetTodoResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	DueDate *time.Time `json:"dueDate,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
}

type UpdateTodoResponse struct {
	Title string `json:"title"`
	Description string `json:"description"`
	IsCompleted bool `json:"isCompleted"`
	DueDate *time.Time `json:"dueDate,omitempty"`
	CreatedAt *time.Time `json:"createdAt"`
}