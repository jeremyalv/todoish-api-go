package server

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jeremyalv/go-todo-api/api/v1/repository"
	"github.com/jeremyalv/go-todo-api/constants"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/mux"
	"github.com/jeremyalv/go
	"github.com/jeremyalv/go-todo-api/api/v1/handlers"
	"github.com/jeremyalv/go-todo-api/api/v1/services"
	"github.com/jeremyalv/go-todo-api/config"
	"github.com/gorilla/mux"
)

type Server struct {
	address string
	service services.IServices
	handler handlers.ITodoHandler
	cfg     *config.Config
	db      *sql.DB
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

	// Initialize Server.service dependencies (e.g. repositories)
	todoStore := repository.New(db)

	svr := &Server{
		address: addr,
		cfg:     cfg,
		db:      db,
	}

	svr.service = services.New().
		WithConfig(*cfg).
		WithTodoRepo(todoStore)

	svr.handler = handlers.New().
		WithService(svr.service)

	return svr
}

func (s *Server) registerRoutes(router *mux.Router) {

	
	v1Router.HandleFunc(constants.TodoEndpoint, s.handler.CreateTodo).Methods(http.MethodPost)
	v1Router.HandleFunc(constants.TodoEndpoint, s.handler.GetTodo).Methods(http.MethodGet)
	v1Router.HandleFunc(constants.TodoEndpoint, s.handler.UpdateTodo).Methods(http.MethodPut)
	v1Router.HandleFunc(constants.TodoEndpoint, s.handler.DeleteTodo).Methods(http.MethodDelete)
}

func (s *Server) Start() {
	router := mux.NewRouter()

	
		Addr:    s.address,
		Addr: s.address,
		Handler: router,

	
	go func() {
		if err := httpServer.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Panicf("%v: error listening to address %s", err, s.address)
		}

	

	
	// Create channel to listen to OS signals
	sigChan := make(chan os.Signal, 1)

	
	sig := <-sigChan

	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.GracefulServerTimeoutInSeconds)*time.Second)
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(s.cfg.GracefulServerTimeoutInSeconds) * time.Second)

	

	
	// Gracefully shut down the server
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Printf("Error shutting down server")

	
	// Resource cleanup
	if err := s.db.Close(); err != nil {
		log.Printf("Error in closing db connection")

	
	log.Printf("Server shut down gracefully")
e
}