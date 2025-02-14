package handlers

import (
	"context"
	"encoding/json"
	"github.com/jeremyalv/go-todo-api/constants"
	"github.com/jeremyalv/go-todo-api/models"
	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
	"github.com/jeremyalv/go-todo-api/pkg/validator"
	"net/http"
)

func (h *todoHandler) GetTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(w, constants.ErrInvalidMethod, http.StatusMethodNotAllowed)
	}

	// In production apps, this Context should be return value of some relevant process prior to this method's business logic
	// For example, it could be a context generated by your tracing library after starting a `Span`
	ctx := context.Background()

	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJSON)

	queryParams := req.URL.Query()
	todoId := queryParams.Get(constants.CtxTodoId)

	requestObj := request.GetTodoRequest{
		TodoId: todoId,
	}

	var todoValidate models.TodoValidator
	if todoId != "" {
		todoValidate.TodoId = todoId
		validateTodoErr := validator.ValidateRequest(todoValidate)
		if validateTodoErr != nil {
			http.Error(w, constants.ErrPreconditionFailed, http.StatusPreconditionFailed)
		}
	}

	resp, err := h.Service.GetTodo(ctx, requestObj)
	if err != nil {
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
	httpResponse := &response.GetTodoResponse{
		Title:       resp.Title,
		Description: resp.Description,
		IsCompleted: resp.IsCompleted,
		DueDate:     resp.DueDate,
		CreatedAt:   resp.CreatedAt,
	}

	err = json.NewEncoder(w).Encode(&httpResponse)
	if err != nil {
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}
}
