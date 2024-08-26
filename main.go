package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger, _ := initLog()
	_, err := initDB(logger, "todos.db")
	if err != nil {
		logger.Fatal(err.Error())
	}

}

func initDB(logger Logger, dbName string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", dbName)
	if err != nil {
		return nil, fmt.Errorf("failed to open the database: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %w", err)
	}
	logger.Info("Database connection established successfully!")
	return db, nil
}

func initLog() (Logger, error) {
	logger := NewLogger()
	return logger, nil
}
