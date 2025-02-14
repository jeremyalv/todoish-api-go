package main

import (
	"github.com/jeremyalv/go-todo-api/config"
	"github.com/jeremyalv/go-todo-api/server"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig()

	sv := server.New(cfg)
	if sv == nil {
		log.Panicf("Failed to start service")
		return
	}
	sv.Start()
}
