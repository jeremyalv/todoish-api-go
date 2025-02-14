package response

import (
	"time"
)

type CreateTodoResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}

type GetTodoResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
	Todo         *Todo  `json:"todo"`
}

type GetMyTodoResponse struct {
	Code         int     `json:"code"`
	Message      string  `json:"message"`
	ResponseTime string  `json:"responseTime"`
	Todos        []*Todo `json:"todos"`
}

type UpdateTodoResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}

type DeleteTodoResponse struct {
	Code         int    `json:"code"`
	Message      string `json:"message"`
	ResponseTime string `json:"responseTime"`
}

type Todo struct {
	Id          string     `json:"id"`
	OwnerId     string     `json:"ownerId"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	IsCompleted bool       `json:"isCompleted"`
	DueDate     *time.Time `json:"dueDate,omitempty"`
	Created     *time.Time `json:"created"`
}
