package handlers

import (
	"net/http"

	"github.com/jeremyalv/go-todo-api/api/v1/services"
)

type todoHandler struct {
	Service services.IServices
}

type ITodoHandler interface {
	CreateTodo(w http.ResponseWriter, req *http.Request)
	GetTodo(w http.ResponseWriter, req *http.Request)
	UpdateTodo(w http.ResponseWriter, req *http.Request)
	DeleteTodo(w http.ResponseWriter, req *http.Request)
}

func (h *todoHandler) WithService(s services.IServices) *todoHandler {
	h.Service = s
	return h
}

func New() *todoHandler {
	return &todoHandler{}
}