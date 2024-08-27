package main

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/elyarsadig/todo-app/database"
	"github.com/elyarsadig/todo-app/logger"
	"github.com/elyarsadig/todo-app/router"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	logger, _ := initLog()
	ctx := context.Background()
	db, err := initDB(logger, "todos.db")
	if err != nil {
		logger.Fatal(err.Error())
	}
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      routes(),
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}
	initServer(ctx, logger, srv, db)
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

func initServer(ctx context.Context, logger logger.Logger, srv *http.Server, db *sql.DB) {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, syscall.SIGINT, syscall.SIGTERM)
	errorChan := make(chan error, 1)
	go func() {
		errorChan <- srv.ListenAndServe()
	}()
	logger.Info(fmt.Sprintf("Server is listening on %v", srv.Addr))

	select {
	case sig := <-interrupt:
		logger.Warn(fmt.Sprintf("Signal received: %v", sig))
	case err := <-errorChan:
		if err != http.ErrServerClosed {
			logger.Error(fmt.Sprintf("Server Error: %v", err))
		}
	}
	shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logger.Error(fmt.Sprintf("Server Shutdown Error: %v", err))
	}

	if err := db.Close(); err != nil {
		logger.Error(fmt.Sprintf("Database Close Error: %v", err))
	}
}

func routes() router.Router {
	r := router.NewV1()
	return r
}
