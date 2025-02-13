package handlers

import (
	"context"
	"encoding/json"
	"io"
	"net/http"

	"github.com/jeremyalv/go-todo-api/constants"
	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
	"github.com/jeremyalv/go-todo-api/pkg/datetime"
)

func (h *todoHandler) CreateTodo(w http.ResponseWriter, req *http.Request) {
	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJSON)
	
	if req.Method != http.MethodPost {
		http.Error(w, constants.ErrInvalidMethod, http.StatusMethodNotAllowed)
	}

	payload := request.CreateTodoRequest{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
	}
	
	ctx := context.Background()
	
	err = h.Service.CreateTodo(ctx, payload)
	if err != nil {
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}

	response := response.CreateTodoResponse{
		Code: http.StatusCreated,
		Message: constants.MessageOk,
		ResponseTime: datetime.GetTimeNow(),
		Todo: &payload,
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}