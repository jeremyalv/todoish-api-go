package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeremyalv/go-todo-api/config"
)

func main() {
	cfg := config.LoadConfig()
	fmt.Println("cfg loaded. ", cfg.Username)
}
