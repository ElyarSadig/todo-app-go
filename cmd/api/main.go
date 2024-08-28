package main

import (
	"github.com/elyarsadig/todo-app/pkg/db/sqlite"
	"github.com/elyarsadig/todo-app/pkg/logger"
)

func main() {
	logger := logger.New()
	db, err := sqlite.NewSqliteDB("todo.db")
	if err != nil {
		logger.Fatal(err.Error())
	}
	defer db.Close()
}