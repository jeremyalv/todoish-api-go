package handlers

import (
	"context"
	"encoding/json"
	"github.com/jeremyalv/go-todo-api/constants"
	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
	"github.com/jeremyalv/go-todo-api/pkg/datetime"
	"io"
	"net/http"
)

func (h *todoHandler) DeleteTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodDelete {
		http.Error(w, constants.ErrInvalidMethod, http.StatusMethodNotAllowed)
	}

	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJSON)

	payload := request.DeleteTodoRequest{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
	}

	ctx := context.Background()

	err = h.Service.DeleteTodo(ctx, payload)
	if err != nil {
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}

	res := response.DeleteTodoResponse{
		Code:         http.StatusOK,
		Message:      constants.MessageOk,
		ResponseTime: datetime.GetTimeNow(),
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}
}
