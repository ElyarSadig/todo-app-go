package main

import (
	"os"

	"github.com/elyarsadig/todo-app/internal/server"
	"github.com/elyarsadig/todo-app/migrations"
	"github.com/elyarsadig/todo-app/pkg/db/sqlite"
	"github.com/elyarsadig/todo-app/pkg/logger"
)

func main() {
	logger := logger.New(os.Stdout)
	db, err := sqlite.NewSqliteDB("todo.db")
	if err != nil {
		logger.Fatal(err)
	}
	defer db.Close()
	err = migrations.RunMigrationsV1(db)
	if err != nil {
		logger.Fatal(err)
	}
	logger.Info("successfully migrated")
	s := server.New(db, logger)
	if err := s.Run(); err != nil {
		logger.Fatal(err)
	}
}
