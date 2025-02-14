package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jeremyalv/go-todo-api/config"
)

type Connection struct {
	MySQL config.Config
}

type MySQL struct {
	Host               string
	User               string
	Password           string
	DBName             string
	MaxPoolSize        int
	MaxIdleConnections int
}

type DBConnection interface {
	DBConnect()
}

func NewDBConnection(cfg config.Config) *Connection {
	return &Connection{
		MySQL: cfg,
	}
}

func (db *Connection) DBConnect() *sql.DB {
	dbConn, errConn := sql.Open("mysql", db.MySQL.URL)
	if errConn != nil {
		log.Panicf("Error while connecting to database: %v", errConn)
		return nil
	}
	dbConn.SetMaxOpenConns(db.MySQL.MaxOpenConnections)
	dbConn.SetMaxIdleConns(db.MySQL.MaxIdleConnections)
	return dbConn
}
