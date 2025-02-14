package handlers

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/jeremyalv/go-todo-api/constants"
	"github.com/jeremyalv/go-todo-api/models/request"
	"github.com/jeremyalv/go-todo-api/models/response"
	"github.com/jeremyalv/go-todo-api/pkg/datetime"
)

func (h *todoHandler) CreateTodo(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		http.Error(w, constants.ErrInvalidMethod, http.StatusMethodNotAllowed)
		return
	}

	w.Header().Set(constants.HeaderContentType, constants.MIMEApplicationJSON)

	payload := request.CreateTodoRequest{}
	body, err := io.ReadAll(req.Body)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
		return
	}
	err = json.Unmarshal(body, &payload)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrBadRequest, http.StatusBadRequest)
		return
	}

	ctx := context.Background()

	err = h.Service.CreateTodo(ctx, payload)
	if err != nil {
		log.Printf("ERROR: %v", err)
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	res := response.CreateTodoResponse{
		Code:         http.StatusCreated,
		Message:      constants.MessageOk,
		ResponseTime: datetime.GetTimeNow(),
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, constants.ErrInternalServerError, http.StatusInternalServerError)
		return
	}
}
