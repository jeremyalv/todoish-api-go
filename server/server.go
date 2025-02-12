package server

import (
	"database/sql"

	"github.com/jeremyalv/go-todo-api/api/v1/services"
	"github.com/jeremyalv/go-todo-api/config"
)

type Server struct {
	service services.IServices
	cfg *config.Config
	db *sql.DB
	address string
}