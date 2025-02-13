package server

import (
	"database/sql"
	"log"

	"github.com/jeremyalv/go-todo-api/api/v1/handlers"
	"github.com/jeremyalv/go-todo-api/api/v1/services"
	"github.com/jeremyalv/go-todo-api/config"
	"github.com/jeremyalv/go-todo-api/pkg/database"
)

type Server struct {
	address string
	service services.IServices
	handler handlers.ITodoHandler
	cfg *config.Config
	db *sql.DB
}

func New(cfg *config.Config) *Server {
	addr := ":9000"
	conn := database.NewDBConnection(*cfg)
	if conn == nil {
		log.Panic("Expecting DB connection object but received nil")
		return nil
	}
	db := conn.DBConnect()
	if db == nil {
		log.Panic("Expecting DB connection object but received nil")
		return nil
	}
	
	svr := &Server{
		address: addr,
		cfg: cfg,
		db: db,
	}

	svr.service = services.New().
					WithConfig(*cfg)

	svr.handler = handlers.New().WithService(svr.service)

	return svr
}