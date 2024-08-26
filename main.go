package main

import (
	"database/sql"

	"github.com/elyarsadig/todo-app/database"
	"github.com/elyarsadig/todo-app/logger"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger, _ := initLog()
	db, err := initDB(logger, "todos.db")
	if err != nil {
		logger.Fatal(err.Error())
	}
	db.Close()
}

func initDB(logger logger.Logger, dbName string) (*sql.DB, error) {
	db, err := database.New(dbName)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("database connection established successfully!")
	err = database.CreateTables(db)
	if err != nil {
		logger.Fatal(err.Error())
	}
	logger.Info("tables created successfully!")
	return db, nil
}

func initLog() (logger.Logger, error) {
	logger := logger.New()
	return logger, nil
}
