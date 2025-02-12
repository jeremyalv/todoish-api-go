package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeremyalv/go-todo-api/config"
)

func main() {
	cfg := config.LoadConfig()
}
