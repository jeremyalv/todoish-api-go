package handlers

import (
	"context"
	"encoding/json"
	"github.com/jeremyalv/go-todo-api/models/response"
	"github.com/jeremyalv/go-todo-api/pkg/datetime"
	"io"
	"log"
	"net/http"

	"github.com/jeremyalv/go-todo-api/constants"
	"github.com/jeremyalv/go-todo-api/models/request"
)

func (h *todoHandler) UpdateTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPut {
		http.Error(w, constants.ErrInvalidMethod, http.StatusMethodNotAllowed)
	}

	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJSON)

	payload := request.UpdateTodoRequest{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
	}

	ctx := context.Background()

	err = h.Service.UpdateTodo(ctx, payload)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}

	res := response.UpdateTodoResponse{
		Code:         http.StatusOK,
		Message:      constants.MessageOk,
		ResponseTime: datetime.GetTimeNow(),
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
	}
}
